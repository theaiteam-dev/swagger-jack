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
	apiUpdatePutCmdBody string
	apiUpdatePutCmdBodyFile string
	apiUpdatePutCmd_id string
	apiUpdatePutCmd_name string
	apiUpdatePutCmd_type string
	apiUpdatePutCmd_workspaceId string
	apiUpdatePutCmd_channel string
)

var apiUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if apiUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(apiUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			apiUpdatePutCmdBody = string(fileData)
		}
		if apiUpdatePutCmdBody != "" {
			if !json.Valid([]byte(apiUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(apiUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/admin/subscription-groups/", pathParams, queryParams, bodyObj)
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
		bodyMap["id"] = apiUpdatePutCmd_id
		bodyMap["name"] = apiUpdatePutCmd_name
		bodyMap["type"] = apiUpdatePutCmd_type
		bodyMap["workspaceId"] = apiUpdatePutCmd_workspaceId
		bodyMap["channel"] = apiUpdatePutCmd_channel
		resp, err := c.Do("PUT", "/api/admin/subscription-groups/", pathParams, queryParams, bodyMap)
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
	apiCmd.AddCommand(apiUpdatePutCmd)
	apiUpdatePutCmd.Flags().StringVar(&apiUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	apiUpdatePutCmd.Flags().StringVar(&apiUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	apiUpdatePutCmd.Flags().StringVar(&apiUpdatePutCmd_id, "id", "", "")
	apiUpdatePutCmd.Flags().StringVar(&apiUpdatePutCmd_name, "name", "", "")
	apiUpdatePutCmd.Flags().StringVar(&apiUpdatePutCmd_type, "type", "", "")
	apiUpdatePutCmd.Flags().StringVar(&apiUpdatePutCmd_workspaceId, "workspaceId", "", "")
	apiUpdatePutCmd.Flags().StringVar(&apiUpdatePutCmd_channel, "channel", "", "")
	apiUpdatePutCmd.MarkFlagRequired("type")
	apiUpdatePutCmd.MarkFlagRequired("workspaceId")
	apiUpdatePutCmd.MarkFlagRequired("channel")
}
