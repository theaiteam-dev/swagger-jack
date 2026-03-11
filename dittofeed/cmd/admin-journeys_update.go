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
	adminJourneysUpdateCmdBody string
	adminJourneysUpdateCmdBodyFile string
	adminJourneysUpdateCmd_status string
	adminJourneysUpdateCmd_updatedAt string
	adminJourneysUpdateCmd_workspaceId string
	adminJourneysUpdateCmd_canRunMultiple bool
	adminJourneysUpdateCmd_definitionEntryNode string
	adminJourneysUpdateCmd_definitionExitNodeType string
	adminJourneysUpdateCmd_definitionNodes []string
	adminJourneysUpdateCmd_draft string
	adminJourneysUpdateCmd_id string
	adminJourneysUpdateCmd_name string
)

var adminJourneysUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if err := validate.Enum("definition.exitNode.type", adminJourneysUpdateCmd_definitionExitNodeType, []string{"ExitNode"}); err != nil { return err }
		if adminJourneysUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminJourneysUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminJourneysUpdateCmdBody = string(fileData)
		}
		if adminJourneysUpdateCmdBody != "" {
			if !json.Valid([]byte(adminJourneysUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminJourneysUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/admin/journeys/", pathParams, queryParams, bodyObj)
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
		bodyMap["status"] = adminJourneysUpdateCmd_status
		bodyMap["updatedAt"] = adminJourneysUpdateCmd_updatedAt
		bodyMap["workspaceId"] = adminJourneysUpdateCmd_workspaceId
		bodyMap["canRunMultiple"] = adminJourneysUpdateCmd_canRunMultiple
		{
			_parts := strings.Split("definition.entryNode", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = adminJourneysUpdateCmd_definitionEntryNode
		}
		{
			_parts := strings.Split("definition.exitNode.type", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = adminJourneysUpdateCmd_definitionExitNodeType
		}
		{
			_parts := strings.Split("definition.nodes", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = adminJourneysUpdateCmd_definitionNodes
		}
		bodyMap["draft"] = adminJourneysUpdateCmd_draft
		bodyMap["id"] = adminJourneysUpdateCmd_id
		bodyMap["name"] = adminJourneysUpdateCmd_name
		resp, err := c.Do("PUT", "/api/admin/journeys/", pathParams, queryParams, bodyMap)
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
	adminJourneysCmd.AddCommand(adminJourneysUpdateCmd)
	adminJourneysUpdateCmd.Flags().StringVar(&adminJourneysUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminJourneysUpdateCmd.Flags().StringVar(&adminJourneysUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	adminJourneysUpdateCmd.Flags().StringVar(&adminJourneysUpdateCmd_status, "status", "", "")
	adminJourneysUpdateCmd.Flags().StringVar(&adminJourneysUpdateCmd_updatedAt, "updatedAt", "", "")
	adminJourneysUpdateCmd.Flags().StringVar(&adminJourneysUpdateCmd_workspaceId, "workspaceId", "", "")
	adminJourneysUpdateCmd.Flags().BoolVar(&adminJourneysUpdateCmd_canRunMultiple, "canRunMultiple", false, "")
	adminJourneysUpdateCmd.Flags().StringVar(&adminJourneysUpdateCmd_definitionEntryNode, "definition.entryNode", "", "")
	adminJourneysUpdateCmd.Flags().StringVar(&adminJourneysUpdateCmd_definitionExitNodeType, "definition.exitNode.type", "", "(ExitNode)")
	adminJourneysUpdateCmd.RegisterFlagCompletionFunc("definition.exitNode.type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"ExitNode"}, cobra.ShellCompDirectiveNoFileComp
	})
	adminJourneysUpdateCmd.Flags().StringArrayVar(&adminJourneysUpdateCmd_definitionNodes, "definition.nodes", nil, "")
	adminJourneysUpdateCmd.Flags().StringVar(&adminJourneysUpdateCmd_draft, "draft", "", "")
	adminJourneysUpdateCmd.Flags().StringVar(&adminJourneysUpdateCmd_id, "id", "", "")
	adminJourneysUpdateCmd.Flags().StringVar(&adminJourneysUpdateCmd_name, "name", "", "")
	adminJourneysUpdateCmd.MarkFlagRequired("workspaceId")
	adminJourneysUpdateCmd.MarkFlagRequired("definition.entryNode")
	adminJourneysUpdateCmd.MarkFlagRequired("definition.exitNode.type")
	adminJourneysUpdateCmd.MarkFlagRequired("definition.nodes")
}
