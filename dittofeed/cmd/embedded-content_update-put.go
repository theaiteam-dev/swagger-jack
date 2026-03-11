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
	embeddedContentUpdatePutCmdBody string
	embeddedContentUpdatePutCmdBodyFile string
	embeddedContentUpdatePutCmd_workspaceId string
	embeddedContentUpdatePutCmd_definition string
	embeddedContentUpdatePutCmd_draft string
	embeddedContentUpdatePutCmd_id string
	embeddedContentUpdatePutCmd_name string
	embeddedContentUpdatePutCmd_resourceType string
)

var embeddedContentUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if embeddedContentUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedContentUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedContentUpdatePutCmdBody = string(fileData)
		}
		if embeddedContentUpdatePutCmdBody != "" {
			if !json.Valid([]byte(embeddedContentUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedContentUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api-l/embedded/content/templates", pathParams, queryParams, bodyObj)
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
		bodyMap["workspaceId"] = embeddedContentUpdatePutCmd_workspaceId
		bodyMap["definition"] = embeddedContentUpdatePutCmd_definition
		bodyMap["draft"] = embeddedContentUpdatePutCmd_draft
		bodyMap["id"] = embeddedContentUpdatePutCmd_id
		bodyMap["name"] = embeddedContentUpdatePutCmd_name
		bodyMap["resourceType"] = embeddedContentUpdatePutCmd_resourceType
		resp, err := c.Do("PUT", "/api-l/embedded/content/templates", pathParams, queryParams, bodyMap)
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
	embeddedContentCmd.AddCommand(embeddedContentUpdatePutCmd)
	embeddedContentUpdatePutCmd.Flags().StringVar(&embeddedContentUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedContentUpdatePutCmd.Flags().StringVar(&embeddedContentUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	embeddedContentUpdatePutCmd.Flags().StringVar(&embeddedContentUpdatePutCmd_workspaceId, "workspaceId", "", "")
	embeddedContentUpdatePutCmd.Flags().StringVar(&embeddedContentUpdatePutCmd_definition, "definition", "", "")
	embeddedContentUpdatePutCmd.Flags().StringVar(&embeddedContentUpdatePutCmd_draft, "draft", "", "")
	embeddedContentUpdatePutCmd.Flags().StringVar(&embeddedContentUpdatePutCmd_id, "id", "", "")
	embeddedContentUpdatePutCmd.Flags().StringVar(&embeddedContentUpdatePutCmd_name, "name", "", "")
	embeddedContentUpdatePutCmd.Flags().StringVar(&embeddedContentUpdatePutCmd_resourceType, "resourceType", "", "")
	embeddedContentUpdatePutCmd.MarkFlagRequired("workspaceId")
	embeddedContentUpdatePutCmd.MarkFlagRequired("name")
}
