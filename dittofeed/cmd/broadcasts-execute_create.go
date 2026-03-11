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
	broadcastsExecuteCreateCmdBody string
	broadcastsExecuteCreateCmdBodyFile string
	broadcastsExecuteCreateCmd_broadcastName string
	broadcastsExecuteCreateCmd_messageTemplateDefinition string
	broadcastsExecuteCreateCmd_segmentDefinitionEntryNode string
	broadcastsExecuteCreateCmd_segmentDefinitionNodes []string
	broadcastsExecuteCreateCmd_subscriptionGroupId string
	broadcastsExecuteCreateCmd_workspaceId string
)

var broadcastsExecuteCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if broadcastsExecuteCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(broadcastsExecuteCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			broadcastsExecuteCreateCmdBody = string(fileData)
		}
		if broadcastsExecuteCreateCmdBody != "" {
			if !json.Valid([]byte(broadcastsExecuteCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(broadcastsExecuteCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/broadcasts/execute", pathParams, queryParams, bodyObj)
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
		bodyMap["broadcastName"] = broadcastsExecuteCreateCmd_broadcastName
		bodyMap["messageTemplateDefinition"] = broadcastsExecuteCreateCmd_messageTemplateDefinition
		{
			_parts := strings.Split("segmentDefinition.entryNode", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = broadcastsExecuteCreateCmd_segmentDefinitionEntryNode
		}
		{
			_parts := strings.Split("segmentDefinition.nodes", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = broadcastsExecuteCreateCmd_segmentDefinitionNodes
		}
		bodyMap["subscriptionGroupId"] = broadcastsExecuteCreateCmd_subscriptionGroupId
		bodyMap["workspaceId"] = broadcastsExecuteCreateCmd_workspaceId
		resp, err := c.Do("POST", "/api/broadcasts/execute", pathParams, queryParams, bodyMap)
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
	broadcastsExecuteCmd.AddCommand(broadcastsExecuteCreateCmd)
	broadcastsExecuteCreateCmd.Flags().StringVar(&broadcastsExecuteCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	broadcastsExecuteCreateCmd.Flags().StringVar(&broadcastsExecuteCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	broadcastsExecuteCreateCmd.Flags().StringVar(&broadcastsExecuteCreateCmd_broadcastName, "broadcastName", "", "")
	broadcastsExecuteCreateCmd.Flags().StringVar(&broadcastsExecuteCreateCmd_messageTemplateDefinition, "messageTemplateDefinition", "", "")
	broadcastsExecuteCreateCmd.Flags().StringVar(&broadcastsExecuteCreateCmd_segmentDefinitionEntryNode, "segmentDefinition.entryNode", "", "")
	broadcastsExecuteCreateCmd.Flags().StringArrayVar(&broadcastsExecuteCreateCmd_segmentDefinitionNodes, "segmentDefinition.nodes", nil, "")
	broadcastsExecuteCreateCmd.Flags().StringVar(&broadcastsExecuteCreateCmd_subscriptionGroupId, "subscriptionGroupId", "", "")
	broadcastsExecuteCreateCmd.Flags().StringVar(&broadcastsExecuteCreateCmd_workspaceId, "workspaceId", "", "")
	broadcastsExecuteCreateCmd.MarkFlagRequired("broadcastName")
	broadcastsExecuteCreateCmd.MarkFlagRequired("messageTemplateDefinition")
	broadcastsExecuteCreateCmd.MarkFlagRequired("segmentDefinition.entryNode")
	broadcastsExecuteCreateCmd.MarkFlagRequired("segmentDefinition.nodes")
	broadcastsExecuteCreateCmd.MarkFlagRequired("workspaceId")
}
