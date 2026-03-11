package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	webhooksResendCreateCmdBody string
	webhooksResendCreateCmdBodyFile string
	webhooksResendCreateCmd_createdAt string
	webhooksResendCreateCmd_dataFrom string
	webhooksResendCreateCmd_dataSubject string
	webhooksResendCreateCmd_dataTagsBroadcastId string
	webhooksResendCreateCmd_dataTagsJourneyId string
	webhooksResendCreateCmd_dataTagsMessageId string
	webhooksResendCreateCmd_dataTagsNodeId string
	webhooksResendCreateCmd_dataTagsRunId string
	webhooksResendCreateCmd_dataTagsTemplateId string
	webhooksResendCreateCmd_dataTagsUserId string
	webhooksResendCreateCmd_dataTagsWorkspaceId string
	webhooksResendCreateCmd_dataTo []string
	webhooksResendCreateCmd_dataCreatedAt string
	webhooksResendCreateCmd_dataEmailId string
	webhooksResendCreateCmd_type string
)

var webhooksResendCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if webhooksResendCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(webhooksResendCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			webhooksResendCreateCmdBody = string(fileData)
		}
		if webhooksResendCreateCmdBody != "" {
			if !json.Valid([]byte(webhooksResendCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(webhooksResendCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/webhooks/resend", pathParams, queryParams, bodyObj)
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
		bodyMap["created_at"] = webhooksResendCreateCmd_createdAt
		{
			_parts := strings.Split("data.from", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataFrom
		}
		{
			_parts := strings.Split("data.subject", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataSubject
		}
		{
			_parts := strings.Split("data.tags.broadcastId", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataTagsBroadcastId
		}
		{
			_parts := strings.Split("data.tags.journeyId", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataTagsJourneyId
		}
		{
			_parts := strings.Split("data.tags.messageId", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataTagsMessageId
		}
		{
			_parts := strings.Split("data.tags.nodeId", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataTagsNodeId
		}
		{
			_parts := strings.Split("data.tags.runId", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataTagsRunId
		}
		{
			_parts := strings.Split("data.tags.templateId", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataTagsTemplateId
		}
		{
			_parts := strings.Split("data.tags.userId", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataTagsUserId
		}
		{
			_parts := strings.Split("data.tags.workspaceId", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataTagsWorkspaceId
		}
		{
			_parts := strings.Split("data.to", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataTo
		}
		{
			_parts := strings.Split("data.created_at", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataCreatedAt
		}
		{
			_parts := strings.Split("data.email_id", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = webhooksResendCreateCmd_dataEmailId
		}
		bodyMap["type"] = webhooksResendCreateCmd_type
		resp, err := c.Do("POST", "/api/webhooks/resend", pathParams, queryParams, bodyMap)
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
	webhooksResendCmd.AddCommand(webhooksResendCreateCmd)
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_createdAt, "created_at", "", "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_dataFrom, "data.from", "", "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_dataSubject, "data.subject", "", "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_dataTagsBroadcastId, "data.tags.broadcastId", "", "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_dataTagsJourneyId, "data.tags.journeyId", "", "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_dataTagsMessageId, "data.tags.messageId", "", "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_dataTagsNodeId, "data.tags.nodeId", "", "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_dataTagsRunId, "data.tags.runId", "", "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_dataTagsTemplateId, "data.tags.templateId", "", "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_dataTagsUserId, "data.tags.userId", "", "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_dataTagsWorkspaceId, "data.tags.workspaceId", "", "")
	webhooksResendCreateCmd.Flags().StringArrayVar(&webhooksResendCreateCmd_dataTo, "data.to", nil, "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_dataCreatedAt, "data.created_at", "", "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_dataEmailId, "data.email_id", "", "")
	webhooksResendCreateCmd.Flags().StringVar(&webhooksResendCreateCmd_type, "type", "", "")
	webhooksResendCreateCmd.MarkFlagRequired("created_at")
	webhooksResendCreateCmd.MarkFlagRequired("data.from")
	webhooksResendCreateCmd.MarkFlagRequired("data.subject")
	webhooksResendCreateCmd.MarkFlagRequired("data.to")
	webhooksResendCreateCmd.MarkFlagRequired("data.created_at")
	webhooksResendCreateCmd.MarkFlagRequired("data.email_id")
	webhooksResendCreateCmd.MarkFlagRequired("type")
}
