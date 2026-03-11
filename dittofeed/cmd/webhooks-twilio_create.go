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
	webhooksTwilioCreateCmdBody string
	webhooksTwilioCreateCmdBodyFile string
	webhooksTwilioCreateCmd_workspaceId string
	webhooksTwilioCreateCmd_userId string
	webhooksTwilioCreateCmd_subscriptionGroupId string
	webhooksTwilioCreateCmd_messageId string
	webhooksTwilioCreateCmd_journeyId string
	webhooksTwilioCreateCmd_templateId string
	webhooksTwilioCreateCmd_nodeId string
	webhooksTwilioCreateCmd_runId string
	webhooksTwilioCreateCmd_From string
	webhooksTwilioCreateCmd_MessageSid string
	webhooksTwilioCreateCmd_To string
	webhooksTwilioCreateCmd_Body string
	webhooksTwilioCreateCmd_MessagingServiceSid string
	webhooksTwilioCreateCmd_SmsSid string
	webhooksTwilioCreateCmd_SmsStatus string
	webhooksTwilioCreateCmd_AccountSid string
	webhooksTwilioCreateCmd_ApiVersion string
)

var webhooksTwilioCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = webhooksTwilioCreateCmd_workspaceId
		queryParams["userId"] = webhooksTwilioCreateCmd_userId
		queryParams["subscriptionGroupId"] = webhooksTwilioCreateCmd_subscriptionGroupId
		queryParams["messageId"] = webhooksTwilioCreateCmd_messageId
		queryParams["journeyId"] = webhooksTwilioCreateCmd_journeyId
		queryParams["templateId"] = webhooksTwilioCreateCmd_templateId
		queryParams["nodeId"] = webhooksTwilioCreateCmd_nodeId
		queryParams["runId"] = webhooksTwilioCreateCmd_runId
		if webhooksTwilioCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(webhooksTwilioCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			webhooksTwilioCreateCmdBody = string(fileData)
		}
		if webhooksTwilioCreateCmdBody != "" {
			if !json.Valid([]byte(webhooksTwilioCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(webhooksTwilioCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/webhooks/twilio", pathParams, queryParams, bodyObj)
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
		bodyMap["From"] = webhooksTwilioCreateCmd_From
		bodyMap["MessageSid"] = webhooksTwilioCreateCmd_MessageSid
		bodyMap["To"] = webhooksTwilioCreateCmd_To
		bodyMap["Body"] = webhooksTwilioCreateCmd_Body
		bodyMap["MessagingServiceSid"] = webhooksTwilioCreateCmd_MessagingServiceSid
		bodyMap["SmsSid"] = webhooksTwilioCreateCmd_SmsSid
		bodyMap["SmsStatus"] = webhooksTwilioCreateCmd_SmsStatus
		bodyMap["AccountSid"] = webhooksTwilioCreateCmd_AccountSid
		bodyMap["ApiVersion"] = webhooksTwilioCreateCmd_ApiVersion
		resp, err := c.Do("POST", "/api/webhooks/twilio", pathParams, queryParams, bodyMap)
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
	webhooksTwilioCmd.AddCommand(webhooksTwilioCreateCmd)
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_workspaceId, "workspaceId", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_userId, "userId", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_subscriptionGroupId, "subscriptionGroupId", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_messageId, "messageId", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_journeyId, "journeyId", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_templateId, "templateId", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_nodeId, "nodeId", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_runId, "runId", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_From, "From", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_MessageSid, "MessageSid", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_To, "To", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_Body, "Body", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_MessagingServiceSid, "MessagingServiceSid", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_SmsSid, "SmsSid", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_SmsStatus, "SmsStatus", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_AccountSid, "AccountSid", "", "")
	webhooksTwilioCreateCmd.Flags().StringVar(&webhooksTwilioCreateCmd_ApiVersion, "ApiVersion", "", "")
	webhooksTwilioCreateCmd.MarkFlagRequired("workspaceId")
	webhooksTwilioCreateCmd.MarkFlagRequired("userId")
	webhooksTwilioCreateCmd.MarkFlagRequired("From")
	webhooksTwilioCreateCmd.MarkFlagRequired("MessageSid")
	webhooksTwilioCreateCmd.MarkFlagRequired("To")
	webhooksTwilioCreateCmd.MarkFlagRequired("MessagingServiceSid")
	webhooksTwilioCreateCmd.MarkFlagRequired("SmsSid")
	webhooksTwilioCreateCmd.MarkFlagRequired("SmsStatus")
	webhooksTwilioCreateCmd.MarkFlagRequired("AccountSid")
	webhooksTwilioCreateCmd.MarkFlagRequired("ApiVersion")
}
