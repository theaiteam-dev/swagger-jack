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
	contentTemplatesCreatePostCmdBody string
	contentTemplatesCreatePostCmdBodyFile string
	contentTemplatesCreatePostCmd_subscriptionGroupId string
	contentTemplatesCreatePostCmd_tags string
	contentTemplatesCreatePostCmd_userProperties string
	contentTemplatesCreatePostCmd_workspaceId string
	contentTemplatesCreatePostCmd_channel string
	contentTemplatesCreatePostCmd_contents string
)

var contentTemplatesCreatePostCmd = &cobra.Command{
	Use: "create-post",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if contentTemplatesCreatePostCmdBodyFile != "" {
			fileData, err := os.ReadFile(contentTemplatesCreatePostCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			contentTemplatesCreatePostCmdBody = string(fileData)
		}
		if contentTemplatesCreatePostCmdBody != "" {
			if !json.Valid([]byte(contentTemplatesCreatePostCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(contentTemplatesCreatePostCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/content/templates/render", pathParams, queryParams, bodyObj)
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
		bodyMap["subscriptionGroupId"] = contentTemplatesCreatePostCmd_subscriptionGroupId
		bodyMap["tags"] = contentTemplatesCreatePostCmd_tags
		bodyMap["userProperties"] = contentTemplatesCreatePostCmd_userProperties
		bodyMap["workspaceId"] = contentTemplatesCreatePostCmd_workspaceId
		bodyMap["channel"] = contentTemplatesCreatePostCmd_channel
		bodyMap["contents"] = contentTemplatesCreatePostCmd_contents
		resp, err := c.Do("POST", "/api/content/templates/render", pathParams, queryParams, bodyMap)
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
	contentTemplatesCmd.AddCommand(contentTemplatesCreatePostCmd)
	contentTemplatesCreatePostCmd.Flags().StringVar(&contentTemplatesCreatePostCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	contentTemplatesCreatePostCmd.Flags().StringVar(&contentTemplatesCreatePostCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	contentTemplatesCreatePostCmd.Flags().StringVar(&contentTemplatesCreatePostCmd_subscriptionGroupId, "subscriptionGroupId", "", "")
	contentTemplatesCreatePostCmd.Flags().StringVar(&contentTemplatesCreatePostCmd_tags, "tags", "", "")
	contentTemplatesCreatePostCmd.Flags().StringVar(&contentTemplatesCreatePostCmd_userProperties, "userProperties", "", "")
	contentTemplatesCreatePostCmd.Flags().StringVar(&contentTemplatesCreatePostCmd_workspaceId, "workspaceId", "", "")
	contentTemplatesCreatePostCmd.Flags().StringVar(&contentTemplatesCreatePostCmd_channel, "channel", "", "")
	contentTemplatesCreatePostCmd.Flags().StringVar(&contentTemplatesCreatePostCmd_contents, "contents", "", "")
	contentTemplatesCreatePostCmd.MarkFlagRequired("userProperties")
	contentTemplatesCreatePostCmd.MarkFlagRequired("workspaceId")
	contentTemplatesCreatePostCmd.MarkFlagRequired("channel")
	contentTemplatesCreatePostCmd.MarkFlagRequired("contents")
}
