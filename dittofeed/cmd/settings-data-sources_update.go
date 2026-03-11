package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
	"dittofeed/internal/validate"
)

var (
	settingsDataSourcesUpdateCmdBody string
	settingsDataSourcesUpdateCmdBodyFile string
	settingsDataSourcesUpdateCmd_variantType string
	settingsDataSourcesUpdateCmd_variantSharedSecret string
	settingsDataSourcesUpdateCmd_workspaceId string
)

var settingsDataSourcesUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if err := validate.Enum("variant.type", settingsDataSourcesUpdateCmd_variantType, []string{"SegmentIO"}); err != nil { return err }
		if settingsDataSourcesUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(settingsDataSourcesUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			settingsDataSourcesUpdateCmdBody = string(fileData)
		}
		if settingsDataSourcesUpdateCmdBody != "" {
			if !json.Valid([]byte(settingsDataSourcesUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(settingsDataSourcesUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/settings/data-sources", pathParams, queryParams, bodyObj)
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
		{
			_parts := strings.Split("variant.type", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = settingsDataSourcesUpdateCmd_variantType
		}
		{
			_parts := strings.Split("variant.sharedSecret", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = settingsDataSourcesUpdateCmd_variantSharedSecret
		}
		bodyMap["workspaceId"] = settingsDataSourcesUpdateCmd_workspaceId
		resp, err := c.Do("PUT", "/api/settings/data-sources", pathParams, queryParams, bodyMap)
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
	settingsDataSourcesCmd.AddCommand(settingsDataSourcesUpdateCmd)
	settingsDataSourcesUpdateCmd.Flags().StringVar(&settingsDataSourcesUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	settingsDataSourcesUpdateCmd.Flags().StringVar(&settingsDataSourcesUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	settingsDataSourcesUpdateCmd.Flags().StringVar(&settingsDataSourcesUpdateCmd_variantType, "variant.type", "", "(SegmentIO)")
	settingsDataSourcesUpdateCmd.RegisterFlagCompletionFunc("variant.type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"SegmentIO"}, cobra.ShellCompDirectiveNoFileComp
	})
	settingsDataSourcesUpdateCmd.Flags().StringVar(&settingsDataSourcesUpdateCmd_variantSharedSecret, "variant.sharedSecret", "", "")
	settingsDataSourcesUpdateCmd.Flags().StringVar(&settingsDataSourcesUpdateCmd_workspaceId, "workspaceId", "", "")
	settingsDataSourcesUpdateCmd.MarkFlagRequired("variant.type")
	settingsDataSourcesUpdateCmd.MarkFlagRequired("variant.sharedSecret")
	settingsDataSourcesUpdateCmd.MarkFlagRequired("workspaceId")
}
