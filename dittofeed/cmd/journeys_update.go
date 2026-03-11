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
	journeysUpdateCmdBody string
	journeysUpdateCmdBodyFile string
	journeysUpdateCmd_workspaceId string
	journeysUpdateCmd_canRunMultiple bool
	journeysUpdateCmd_definitionEntryNode string
	journeysUpdateCmd_definitionExitNodeType string
	journeysUpdateCmd_definitionNodes []string
	journeysUpdateCmd_draft string
	journeysUpdateCmd_id string
	journeysUpdateCmd_name string
	journeysUpdateCmd_status string
	journeysUpdateCmd_updatedAt string
)

var journeysUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if err := validate.Enum("definition.exitNode.type", journeysUpdateCmd_definitionExitNodeType, []string{"ExitNode"}); err != nil { return err }
		if journeysUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(journeysUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			journeysUpdateCmdBody = string(fileData)
		}
		if journeysUpdateCmdBody != "" {
			if !json.Valid([]byte(journeysUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(journeysUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/journeys/", pathParams, queryParams, bodyObj)
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
		bodyMap["workspaceId"] = journeysUpdateCmd_workspaceId
		bodyMap["canRunMultiple"] = journeysUpdateCmd_canRunMultiple
		{
			_parts := strings.Split("definition.entryNode", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = journeysUpdateCmd_definitionEntryNode
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
			_cur[_parts[len(_parts)-1]] = journeysUpdateCmd_definitionExitNodeType
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
			_cur[_parts[len(_parts)-1]] = journeysUpdateCmd_definitionNodes
		}
		bodyMap["draft"] = journeysUpdateCmd_draft
		bodyMap["id"] = journeysUpdateCmd_id
		bodyMap["name"] = journeysUpdateCmd_name
		bodyMap["status"] = journeysUpdateCmd_status
		bodyMap["updatedAt"] = journeysUpdateCmd_updatedAt
		resp, err := c.Do("PUT", "/api/journeys/", pathParams, queryParams, bodyMap)
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
	journeysCmd.AddCommand(journeysUpdateCmd)
	journeysUpdateCmd.Flags().StringVar(&journeysUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	journeysUpdateCmd.Flags().StringVar(&journeysUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	journeysUpdateCmd.Flags().StringVar(&journeysUpdateCmd_workspaceId, "workspaceId", "", "")
	journeysUpdateCmd.Flags().BoolVar(&journeysUpdateCmd_canRunMultiple, "canRunMultiple", false, "")
	journeysUpdateCmd.Flags().StringVar(&journeysUpdateCmd_definitionEntryNode, "definition.entryNode", "", "")
	journeysUpdateCmd.Flags().StringVar(&journeysUpdateCmd_definitionExitNodeType, "definition.exitNode.type", "", "(ExitNode)")
	journeysUpdateCmd.RegisterFlagCompletionFunc("definition.exitNode.type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"ExitNode"}, cobra.ShellCompDirectiveNoFileComp
	})
	journeysUpdateCmd.Flags().StringArrayVar(&journeysUpdateCmd_definitionNodes, "definition.nodes", nil, "")
	journeysUpdateCmd.Flags().StringVar(&journeysUpdateCmd_draft, "draft", "", "")
	journeysUpdateCmd.Flags().StringVar(&journeysUpdateCmd_id, "id", "", "")
	journeysUpdateCmd.Flags().StringVar(&journeysUpdateCmd_name, "name", "", "")
	journeysUpdateCmd.Flags().StringVar(&journeysUpdateCmd_status, "status", "", "")
	journeysUpdateCmd.Flags().StringVar(&journeysUpdateCmd_updatedAt, "updatedAt", "", "")
	journeysUpdateCmd.MarkFlagRequired("workspaceId")
	journeysUpdateCmd.MarkFlagRequired("definition.entryNode")
	journeysUpdateCmd.MarkFlagRequired("definition.exitNode.type")
	journeysUpdateCmd.MarkFlagRequired("definition.nodes")
}
