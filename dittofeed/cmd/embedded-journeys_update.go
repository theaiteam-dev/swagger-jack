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
	embeddedJourneysUpdateCmdBody string
	embeddedJourneysUpdateCmdBodyFile string
	embeddedJourneysUpdateCmd_draft string
	embeddedJourneysUpdateCmd_id string
	embeddedJourneysUpdateCmd_name string
	embeddedJourneysUpdateCmd_status string
	embeddedJourneysUpdateCmd_updatedAt string
	embeddedJourneysUpdateCmd_workspaceId string
	embeddedJourneysUpdateCmd_canRunMultiple bool
	embeddedJourneysUpdateCmd_definitionEntryNode string
	embeddedJourneysUpdateCmd_definitionExitNodeType string
	embeddedJourneysUpdateCmd_definitionNodes []string
)

var embeddedJourneysUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if err := validate.Enum("definition.exitNode.type", embeddedJourneysUpdateCmd_definitionExitNodeType, []string{"ExitNode"}); err != nil { return err }
		if embeddedJourneysUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedJourneysUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedJourneysUpdateCmdBody = string(fileData)
		}
		if embeddedJourneysUpdateCmdBody != "" {
			if !json.Valid([]byte(embeddedJourneysUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedJourneysUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api-l/embedded/journeys/", pathParams, queryParams, bodyObj)
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
		bodyMap["draft"] = embeddedJourneysUpdateCmd_draft
		bodyMap["id"] = embeddedJourneysUpdateCmd_id
		bodyMap["name"] = embeddedJourneysUpdateCmd_name
		bodyMap["status"] = embeddedJourneysUpdateCmd_status
		bodyMap["updatedAt"] = embeddedJourneysUpdateCmd_updatedAt
		bodyMap["workspaceId"] = embeddedJourneysUpdateCmd_workspaceId
		bodyMap["canRunMultiple"] = embeddedJourneysUpdateCmd_canRunMultiple
		{
			_parts := strings.Split("definition.entryNode", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = embeddedJourneysUpdateCmd_definitionEntryNode
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
			_cur[_parts[len(_parts)-1]] = embeddedJourneysUpdateCmd_definitionExitNodeType
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
			_cur[_parts[len(_parts)-1]] = embeddedJourneysUpdateCmd_definitionNodes
		}
		resp, err := c.Do("PUT", "/api-l/embedded/journeys/", pathParams, queryParams, bodyMap)
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
	embeddedJourneysCmd.AddCommand(embeddedJourneysUpdateCmd)
	embeddedJourneysUpdateCmd.Flags().StringVar(&embeddedJourneysUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedJourneysUpdateCmd.Flags().StringVar(&embeddedJourneysUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	embeddedJourneysUpdateCmd.Flags().StringVar(&embeddedJourneysUpdateCmd_draft, "draft", "", "")
	embeddedJourneysUpdateCmd.Flags().StringVar(&embeddedJourneysUpdateCmd_id, "id", "", "")
	embeddedJourneysUpdateCmd.Flags().StringVar(&embeddedJourneysUpdateCmd_name, "name", "", "")
	embeddedJourneysUpdateCmd.Flags().StringVar(&embeddedJourneysUpdateCmd_status, "status", "", "")
	embeddedJourneysUpdateCmd.Flags().StringVar(&embeddedJourneysUpdateCmd_updatedAt, "updatedAt", "", "")
	embeddedJourneysUpdateCmd.Flags().StringVar(&embeddedJourneysUpdateCmd_workspaceId, "workspaceId", "", "")
	embeddedJourneysUpdateCmd.Flags().BoolVar(&embeddedJourneysUpdateCmd_canRunMultiple, "canRunMultiple", false, "")
	embeddedJourneysUpdateCmd.Flags().StringVar(&embeddedJourneysUpdateCmd_definitionEntryNode, "definition.entryNode", "", "")
	embeddedJourneysUpdateCmd.Flags().StringVar(&embeddedJourneysUpdateCmd_definitionExitNodeType, "definition.exitNode.type", "", "(ExitNode)")
	embeddedJourneysUpdateCmd.RegisterFlagCompletionFunc("definition.exitNode.type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"ExitNode"}, cobra.ShellCompDirectiveNoFileComp
	})
	embeddedJourneysUpdateCmd.Flags().StringArrayVar(&embeddedJourneysUpdateCmd_definitionNodes, "definition.nodes", nil, "")
	embeddedJourneysUpdateCmd.MarkFlagRequired("workspaceId")
	embeddedJourneysUpdateCmd.MarkFlagRequired("definition.entryNode")
	embeddedJourneysUpdateCmd.MarkFlagRequired("definition.exitNode.type")
	embeddedJourneysUpdateCmd.MarkFlagRequired("definition.nodes")
}
