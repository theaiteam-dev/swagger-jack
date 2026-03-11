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
	webhooksPostmarkCreateCmdBody string
	webhooksPostmarkCreateCmdBodyFile string
	webhooksPostmarkCreateCmd_Metadata string
	webhooksPostmarkCreateCmd_templateId string
	webhooksPostmarkCreateCmd_workspaceId string
	webhooksPostmarkCreateCmd_runId string
	webhooksPostmarkCreateCmd_DeliveredAt string
	webhooksPostmarkCreateCmd_MessageStream string
	webhooksPostmarkCreateCmd_RecordType string
	webhooksPostmarkCreateCmd_broadcastId string
	webhooksPostmarkCreateCmd_journeyId string
	webhooksPostmarkCreateCmd_userId string
	webhooksPostmarkCreateCmd_BouncedAt string
	webhooksPostmarkCreateCmd_Tag string
	webhooksPostmarkCreateCmd_nodeId string
	webhooksPostmarkCreateCmd_MessageID string
	webhooksPostmarkCreateCmd_ReceivedAt string
	webhooksPostmarkCreateCmd_messageId string
)

var webhooksPostmarkCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if webhooksPostmarkCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(webhooksPostmarkCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			webhooksPostmarkCreateCmdBody = string(fileData)
		}
		if webhooksPostmarkCreateCmdBody != "" {
			if !json.Valid([]byte(webhooksPostmarkCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(webhooksPostmarkCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/webhooks/postmark", pathParams, queryParams, bodyObj)
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
		bodyMap["Metadata"] = webhooksPostmarkCreateCmd_Metadata
		bodyMap["templateId"] = webhooksPostmarkCreateCmd_templateId
		bodyMap["workspaceId"] = webhooksPostmarkCreateCmd_workspaceId
		bodyMap["runId"] = webhooksPostmarkCreateCmd_runId
		bodyMap["DeliveredAt"] = webhooksPostmarkCreateCmd_DeliveredAt
		bodyMap["MessageStream"] = webhooksPostmarkCreateCmd_MessageStream
		bodyMap["RecordType"] = webhooksPostmarkCreateCmd_RecordType
		bodyMap["broadcastId"] = webhooksPostmarkCreateCmd_broadcastId
		bodyMap["journeyId"] = webhooksPostmarkCreateCmd_journeyId
		bodyMap["userId"] = webhooksPostmarkCreateCmd_userId
		bodyMap["BouncedAt"] = webhooksPostmarkCreateCmd_BouncedAt
		bodyMap["Tag"] = webhooksPostmarkCreateCmd_Tag
		bodyMap["nodeId"] = webhooksPostmarkCreateCmd_nodeId
		bodyMap["MessageID"] = webhooksPostmarkCreateCmd_MessageID
		bodyMap["ReceivedAt"] = webhooksPostmarkCreateCmd_ReceivedAt
		bodyMap["messageId"] = webhooksPostmarkCreateCmd_messageId
		resp, err := c.Do("POST", "/api/webhooks/postmark", pathParams, queryParams, bodyMap)
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
	webhooksPostmarkCmd.AddCommand(webhooksPostmarkCreateCmd)
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_Metadata, "Metadata", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_templateId, "templateId", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_workspaceId, "workspaceId", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_runId, "runId", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_DeliveredAt, "DeliveredAt", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_MessageStream, "MessageStream", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_RecordType, "RecordType", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_broadcastId, "broadcastId", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_journeyId, "journeyId", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_userId, "userId", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_BouncedAt, "BouncedAt", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_Tag, "Tag", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_nodeId, "nodeId", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_MessageID, "MessageID", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_ReceivedAt, "ReceivedAt", "", "")
	webhooksPostmarkCreateCmd.Flags().StringVar(&webhooksPostmarkCreateCmd_messageId, "messageId", "", "")
	webhooksPostmarkCreateCmd.MarkFlagRequired("Metadata")
	webhooksPostmarkCreateCmd.MarkFlagRequired("MessageStream")
	webhooksPostmarkCreateCmd.MarkFlagRequired("RecordType")
	webhooksPostmarkCreateCmd.MarkFlagRequired("Tag")
	webhooksPostmarkCreateCmd.MarkFlagRequired("MessageID")
}
