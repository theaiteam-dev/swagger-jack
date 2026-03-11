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
	webhooksSegmentCreateCmdBody string
	webhooksSegmentCreateCmdBodyFile string
	webhooksSegmentCreateCmd_messageId string
	webhooksSegmentCreateCmd_timestamp string
)

var webhooksSegmentCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if webhooksSegmentCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(webhooksSegmentCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			webhooksSegmentCreateCmdBody = string(fileData)
		}
		if webhooksSegmentCreateCmdBody != "" {
			if !json.Valid([]byte(webhooksSegmentCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(webhooksSegmentCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/webhooks/segment", pathParams, queryParams, bodyObj)
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
		bodyMap["messageId"] = webhooksSegmentCreateCmd_messageId
		bodyMap["timestamp"] = webhooksSegmentCreateCmd_timestamp
		resp, err := c.Do("POST", "/api/webhooks/segment", pathParams, queryParams, bodyMap)
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
	webhooksSegmentCmd.AddCommand(webhooksSegmentCreateCmd)
	webhooksSegmentCreateCmd.Flags().StringVar(&webhooksSegmentCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	webhooksSegmentCreateCmd.Flags().StringVar(&webhooksSegmentCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	webhooksSegmentCreateCmd.Flags().StringVar(&webhooksSegmentCreateCmd_messageId, "messageId", "", "")
	webhooksSegmentCreateCmd.Flags().StringVar(&webhooksSegmentCreateCmd_timestamp, "timestamp", "", "")
	webhooksSegmentCreateCmd.MarkFlagRequired("messageId")
	webhooksSegmentCreateCmd.MarkFlagRequired("timestamp")
}
