package generator

import (
	"fmt"
	"go/parser"
	"go/token"
	"strings"

	"github.com/queso/swagger-jack/internal/model"
)

// GenerateResourceCmd produces the source for cmd/<resource>.go, which
// declares a Cobra group command for the given resource. All operations on
// the resource are added as sub-commands of this group.
func GenerateResourceCmd(resource model.Resource) (string, error) {
	short := resource.Description
	if short == "" {
		short = resource.Name
	}

	varName := sanitizeIdentifier(resource.Name) + "Cmd"

	src := fmt.Sprintf(`package cmd

import "github.com/spf13/cobra"

var %s = &cobra.Command{
	Use: %q,
	Short: %q,
}

func init() {
	rootCmd.AddCommand(%s)
}
`, varName, resource.Name, short, varName)

	return validateGoSource(src)
}

// GenerateVerbCmd produces the source for cmd/<resource>_<verb>.go, which
// declares the individual Cobra command for a specific HTTP operation.
// Path arguments become positional args; query/body/header parameters become
// flags. cliName is used to derive the token env var and the import path for
// the internal client package.
func GenerateVerbCmd(resource model.Resource, cmd model.Command, cliName string) (string, error) {
	useField := buildUseField(cmd)
	argsExpr := buildArgsExpr(cmd.Args)
	varName := sanitizeIdentifier(resource.Name) + capitalise(sanitizeIdentifier(cmd.Name)) + "Cmd"
	resourceVarName := sanitizeIdentifier(resource.Name) + "Cmd"

	flagVars := buildFlagVarDeclarations(varName, cmd.Flags)
	flagInits := buildFlagInits(varName, cmd.Flags)
	requiredInits := buildRequiredFlagInits(varName, cmd.Flags)

	runEBody := buildRunEBody(cmd, varName, cliName)

	// Determine which imports are needed.
	imports := buildImports(cmd, cliName)

	src := fmt.Sprintf(`package cmd

%s

%s

var %s = &cobra.Command{
	Use: %q,
	Short: %q,
	Args: %s,
	RunE: func(cmd *cobra.Command, args []string) error {
%s	},
}

func init() {
	%s.AddCommand(%s)
%s%s}
`,
		imports,
		flagVars,
		varName,
		useField,
		cmd.Description,
		argsExpr,
		runEBody,
		resourceVarName,
		varName,
		flagInits,
		requiredInits,
	)

	// Clean up double blank lines that result from empty flagVars.
	for strings.Contains(src, "\n\n\n") {
		src = strings.ReplaceAll(src, "\n\n\n", "\n\n")
	}

	return validateGoSource(src)
}

// buildImports constructs the import block for the verb command file.
func buildImports(cmd model.Command, cliName string) string {
	hasBodyFlags := false
	hasIntOrBoolQueryFlag := false
	hasStringSliceQueryFlag := false
	for _, f := range cmd.Flags {
		if f.Source == model.FlagSourceBody {
			hasBodyFlags = true
		}
		if f.Source == model.FlagSourceQuery {
			switch f.Type {
			case model.FlagTypeInt, model.FlagTypeBool:
				hasIntOrBoolQueryFlag = true
			case model.FlagTypeStringSlice:
				hasStringSliceQueryFlag = true
			}
		}
	}

	var imports []string
	if hasBodyFlags {
		imports = append(imports, `"encoding/json"`)
	}
	imports = append(imports, `"fmt"`, `"os"`)
	if hasStringSliceQueryFlag {
		imports = append(imports, `"strings"`)
	}
	if hasIntOrBoolQueryFlag {
		imports = append(imports, `"strconv"`)
	}
	imports = append(imports, `"github.com/spf13/cobra"`)
	if cliName != "" {
		imports = append(imports, fmt.Sprintf(`%q`, cliName+"/internal/client"))
	}
	return "import (\n\t" + strings.Join(imports, "\n\t") + "\n)"
}

// buildRunEBody generates the RunE function body for the given command.
func buildRunEBody(cmd model.Command, varName, cliName string) string {
	var sb strings.Builder

	// Read base URL from root persistent flags.
	sb.WriteString("\t\tbaseURL, _ := cmd.Root().PersistentFlags().GetString(\"base-url\")\n")

	// Read token from env var.
	tokenEnvVar := cliNameToEnvPrefix(cliName) + "_TOKEN"
	fmt.Fprintf(&sb, "\t\ttoken := os.Getenv(%q)\n", tokenEnvVar)

	// Create client.
	if cliName != "" {
		sb.WriteString("\t\tc := client.NewClient(baseURL, token)\n")
	} else {
		sb.WriteString("\t\t_, _ = baseURL, token // client not available without cliName\n")
	}

	// Build pathParams map.
	sb.WriteString("\t\tpathParams := map[string]string{}\n")
	for i, arg := range cmd.Args {
		// The path param key is the original camelCase name from the path template.
		// We derive it from the arg name by reversing kebab→camelCase is complex;
		// instead we use the arg name directly as the key since that's what client.Do needs.
		fmt.Fprintf(&sb, "\t\tpathParams[%q] = args[%d]\n", argToPathKey(arg.Name), i)
	}

	// Build queryParams map — only if there are query flags.
	hasQueryFlags := false
	for _, f := range cmd.Flags {
		if f.Source == model.FlagSourceQuery {
			hasQueryFlags = true
			break
		}
	}

	hasBodyFlags := false
	for _, f := range cmd.Flags {
		if f.Source == model.FlagSourceBody {
			hasBodyFlags = true
			break
		}
	}

	if hasQueryFlags {
		sb.WriteString("\t\tqueryParams := map[string]string{}\n")
		for _, f := range cmd.Flags {
			if f.Source != model.FlagSourceQuery {
				continue
			}
			fVar := flagVarName(varName, f.Name)
			switch f.Type {
			case model.FlagTypeInt:
				fmt.Fprintf(&sb, "\t\tqueryParams[%q] = strconv.Itoa(%s)\n", f.Name, fVar)
			case model.FlagTypeBool:
				fmt.Fprintf(&sb, "\t\tqueryParams[%q] = strconv.FormatBool(%s)\n", f.Name, fVar)
			case model.FlagTypeStringSlice:
				// Read inline to avoid []string assignment into map[string]string.
				fmt.Fprintf(&sb, "\t\t%s_vals, _ := cmd.Flags().GetStringArray(%q)\n", fVar, f.Name)
				fmt.Fprintf(&sb, "\t\tqueryParams[%q] = strings.Join(%s_vals, \",\")\n", f.Name, fVar)
			default:
				fmt.Fprintf(&sb, "\t\tqueryParams[%q] = %s\n", f.Name, fVar)
			}
		}
	} else {
		sb.WriteString("\t\tqueryParams := map[string]string{}\n")
	}

	// Build body map — only if there are body flags.
	if hasBodyFlags {
		sb.WriteString("\t\tbody := map[string]interface{}{}\n")
		for _, f := range cmd.Flags {
			if f.Source != model.FlagSourceBody {
				continue
			}
			fVar := flagVarName(varName, f.Name)
			fmt.Fprintf(&sb, "\t\tbody[%q] = %s\n", f.Name, fVar)
		}
	}

	// Call client.Do.
	bodyArg := "nil"
	if hasBodyFlags {
		bodyArg = "body"
	}

	if cliName != "" {
		fmt.Fprintf(&sb, "\t\tresp, err := c.Do(%q, %q, pathParams, queryParams, %s)\n",
			cmd.HTTPMethod, cmd.Path, bodyArg)
		sb.WriteString("\t\tif err != nil {\n")
		sb.WriteString("\t\t\treturn err\n")
		sb.WriteString("\t\t}\n")
	} else {
		// No client available; emit placeholder that still compiles.
		sb.WriteString("\t\t_, _ = pathParams, queryParams\n")
		if hasBodyFlags {
			sb.WriteString("\t\t_ = body\n")
		}
		sb.WriteString("\t\treturn fmt.Errorf(\"no client configured\")\n")
	}

	// Read --json flag and output.
	if cliName != "" {
		if hasBodyFlags {
			sb.WriteString("\t\tjsonMode, _ := cmd.Root().PersistentFlags().GetBool(\"json\")\n")
			sb.WriteString("\t\tif jsonMode {\n")
			sb.WriteString("\t\t\tfmt.Printf(\"%s\\n\", string(resp))\n")
			sb.WriteString("\t\t} else {\n")
			sb.WriteString("\t\t\tvar out interface{}\n")
			sb.WriteString("\t\t\tif err := json.Unmarshal(resp, &out); err != nil {\n")
			sb.WriteString("\t\t\t\tfmt.Printf(\"%s\\n\", string(resp))\n")
			sb.WriteString("\t\t\t} else {\n")
			sb.WriteString("\t\t\t\tpretty, _ := json.MarshalIndent(out, \"\", \"  \")\n")
			sb.WriteString("\t\t\t\tfmt.Printf(\"%s\\n\", string(pretty))\n")
			sb.WriteString("\t\t\t}\n")
			sb.WriteString("\t\t}\n")
		} else {
			sb.WriteString("\t\tfmt.Printf(\"%s\\n\", string(resp))\n")
		}
		sb.WriteString("\t\treturn nil\n")
	}

	return sb.String()
}

// cliNameToEnvPrefix converts a cliName (e.g. "my-cli") to an env var prefix
// (e.g. "MY_CLI") by uppercasing and replacing hyphens/dots with underscores.
func cliNameToEnvPrefix(cliName string) string {
	return strings.ToUpper(strings.NewReplacer("-", "_", ".", "_").Replace(cliName))
}

// argToPathKey converts a kebab-case arg name back to the camelCase path param
// key used in the URL template (e.g. "user-id" → "userId").
func argToPathKey(name string) string {
	parts := strings.Split(name, "-")
	if len(parts) == 1 {
		return name
	}
	var sb strings.Builder
	sb.WriteString(parts[0])
	for _, p := range parts[1:] {
		if len(p) > 0 {
			sb.WriteString(strings.ToUpper(p[:1]) + p[1:])
		}
	}
	return sb.String()
}

// buildUseField constructs the Use string for a Cobra command: the command
// name followed by angle-bracket placeholders for each positional argument.
func buildUseField(cmd model.Command) string {
	parts := []string{cmd.Name}
	for _, arg := range cmd.Args {
		parts = append(parts, "<"+arg.Name+">")
	}
	return strings.Join(parts, " ")
}

// buildArgsExpr returns the cobra.ExactArgs(N) or cobra.NoArgs expression
// for the number of positional arguments the command expects.
func buildArgsExpr(args []model.Arg) string {
	if len(args) == 0 {
		return "cobra.NoArgs"
	}
	return fmt.Sprintf("cobra.ExactArgs(%d)", len(args))
}

// buildFlagVarDeclarations returns a var block declaring one variable per
// flag so that StringVar / IntVar / BoolVar / StringArrayVar can reference
// them. Returns an empty string when there are no flags.
// StringSlice query flags are excluded because they are read inline in RunE
// via cmd.Flags().GetStringArray().
func buildFlagVarDeclarations(cmdVarName string, flags []model.Flag) string {
	var lines []string
	for _, flag := range flags {
		if flag.Type == model.FlagTypeStringSlice && flag.Source == model.FlagSourceQuery {
			continue // read inline in RunE, no var needed
		}
		goType := flagGoType(flag.Type)
		varName := flagVarName(cmdVarName, flag.Name)
		lines = append(lines, fmt.Sprintf("\t%s %s", varName, goType))
	}
	if len(lines) == 0 {
		return ""
	}
	return "var (\n" + strings.Join(lines, "\n") + "\n)\n"
}

// buildFlagInits returns the lines that register each flag with Cobra inside
// the init() function. Each line is indented by one tab.
func buildFlagInits(cmdVarName string, flags []model.Flag) string {
	var lines []string
	for _, flag := range flags {
		var line string
		if flag.Type == model.FlagTypeStringSlice && flag.Source == model.FlagSourceQuery {
			// Read inline in RunE — register without a pre-declared var.
			line = fmt.Sprintf(`%s.Flags().StringArray(%q, nil, %q)`,
				cmdVarName, flag.Name, flag.Description)
		} else {
			varName := flagVarName(cmdVarName, flag.Name)
			line = buildFlagRegistration(cmdVarName, varName, flag)
		}
		lines = append(lines, "\t"+line)
	}
	if len(lines) == 0 {
		return ""
	}
	return strings.Join(lines, "\n") + "\n"
}

// buildFlagRegistration produces the single Flags().XxxVar(...) call for one flag.
func buildFlagRegistration(cmdVarName, varName string, flag model.Flag) string {
	switch flag.Type {
	case model.FlagTypeInt:
		return fmt.Sprintf(`%s.Flags().IntVar(&%s, %q, 0, %q)`,
			cmdVarName, varName, flag.Name, flag.Description)
	case model.FlagTypeBool:
		return fmt.Sprintf(`%s.Flags().BoolVar(&%s, %q, false, %q)`,
			cmdVarName, varName, flag.Name, flag.Description)
	case model.FlagTypeStringSlice:
		return fmt.Sprintf(`%s.Flags().StringArrayVar(&%s, %q, nil, %q)`,
			cmdVarName, varName, flag.Name, flag.Description)
	default: // FlagTypeString
		return fmt.Sprintf(`%s.Flags().StringVar(&%s, %q, "", %q)`,
			cmdVarName, varName, flag.Name, flag.Description)
	}
}

// buildRequiredFlagInits returns the MarkFlagRequired lines for flags that
// are marked Required: true. Each line is indented by one tab.
func buildRequiredFlagInits(cmdVarName string, flags []model.Flag) string {
	var lines []string
	for _, flag := range flags {
		if flag.Required {
			lines = append(lines, fmt.Sprintf("\t%s.MarkFlagRequired(%q)", cmdVarName, flag.Name))
		}
	}
	if len(lines) == 0 {
		return ""
	}
	return strings.Join(lines, "\n") + "\n"
}

// flagGoType maps a model.FlagType to its Go type string for variable
// declarations.
func flagGoType(t model.FlagType) string {
	switch t {
	case model.FlagTypeInt:
		return "int"
	case model.FlagTypeBool:
		return "bool"
	case model.FlagTypeStringSlice:
		return "[]string"
	default:
		return "string"
	}
}

// flagVarName produces a valid Go identifier for a flag variable by
// combining the command var name with a sanitized version of the flag name.
func flagVarName(cmdVarName, flagName string) string {
	return cmdVarName + "_" + sanitizeIdentifier(flagName)
}

// capitalise returns the string with its first rune uppercased.
func capitalise(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// sanitizeIdentifier converts a kebab-case or dot-notation name into a valid
// Go identifier by replacing hyphens and dots with underscores and
// capitalising each segment after the first.
func sanitizeIdentifier(name string) string {
	parts := strings.FieldsFunc(name, func(r rune) bool {
		return r == '-' || r == '.' || r == '_'
	})
	if len(parts) == 0 {
		return "_"
	}
	var sb strings.Builder
	sb.WriteString(parts[0])
	for _, p := range parts[1:] {
		if len(p) > 0 {
			sb.WriteString(strings.ToUpper(p[:1]) + p[1:])
		}
	}
	return sb.String()
}

// validateGoSource parses the provided source string to confirm it is
// syntactically valid Go and returns it unchanged. An error is returned when
// the source cannot be parsed, with the raw source appended for debugging.
func validateGoSource(src string) (string, error) {
	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, "", src, parser.AllErrors); err != nil {
		return "", fmt.Errorf("generated Go source has syntax errors: %w\n---\n%s", err, src)
	}
	return src, nil
}
