package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	permissionsCreateCmdBody string
	permissionsCreateCmdBodyFile string
	permissionsCreateCmd_email string
	permissionsCreateCmd_role string
	permissionsCreateCmd_workspaceId string
)

var permissionsCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if permissionsCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(permissionsCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			permissionsCreateCmdBody = string(fileData)
		}
		if permissionsCreateCmdBody != "" {
			if !json.Valid([]byte(permissionsCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(permissionsCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/permissions/", pathParams, queryParams, bodyObj)
			if err != nil {
				return err
			}
			jsonMode, _ := cmd.Root().PersistentFlags().GetBool("json")
			noColor, _ := cmd.Root().PersistentFlags().GetBool("no-color")
			if jsonMode {
				fmt.Printf("%s\n", string(resp))
			} else {
				if err := output.PrintTable(resp, noColor); err != nil {
					fmt.Println(string(resp))
				}
			}
			return nil
		}
		bodyMap := map[string]interface{}{}
		bodyMap["email"] = permissionsCreateCmd_email
		bodyMap["role"] = permissionsCreateCmd_role
		bodyMap["workspaceId"] = permissionsCreateCmd_workspaceId
		resp, err := c.Do("POST", "/api/permissions/", pathParams, queryParams, bodyMap)
		if err != nil {
			return err
		}
		jsonMode, _ := cmd.Root().PersistentFlags().GetBool("json")
		noColor, _ := cmd.Root().PersistentFlags().GetBool("no-color")
		if jsonMode {
			fmt.Printf("%s\n", string(resp))
		} else {
			if err := output.PrintTable(resp, noColor); err != nil {
				fmt.Println(string(resp))
			}
		}
		return nil
	},
}

func init() {
	permissionsCmd.AddCommand(permissionsCreateCmd)
	permissionsCreateCmd.Flags().StringVar(&permissionsCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	permissionsCreateCmd.Flags().StringVar(&permissionsCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	permissionsCreateCmd.Flags().StringVar(&permissionsCreateCmd_email, "email", "", "")
	permissionsCreateCmd.Flags().StringVar(&permissionsCreateCmd_role, "role", "", "")
	permissionsCreateCmd.Flags().StringVar(&permissionsCreateCmd_workspaceId, "workspaceId", "", "")
	permissionsCreateCmd.MarkFlagRequired("email")
	permissionsCreateCmd.MarkFlagRequired("role")
	permissionsCreateCmd.MarkFlagRequired("workspaceId")
}
