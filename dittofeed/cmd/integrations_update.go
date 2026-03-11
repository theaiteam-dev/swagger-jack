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
	integrationsUpdateCmdBody string
	integrationsUpdateCmdBodyFile string
	integrationsUpdateCmd_definitionSubscribedSegments []string
	integrationsUpdateCmd_definitionSubscribedUserProperties []string
	integrationsUpdateCmd_definitionType string
	integrationsUpdateCmd_enabled bool
	integrationsUpdateCmd_name string
	integrationsUpdateCmd_workspaceId string
)

var integrationsUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if err := validate.Enum("definition.type", integrationsUpdateCmd_definitionType, []string{"Sync"}); err != nil { return err }
		if integrationsUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(integrationsUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			integrationsUpdateCmdBody = string(fileData)
		}
		if integrationsUpdateCmdBody != "" {
			if !json.Valid([]byte(integrationsUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(integrationsUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/integrations/", pathParams, queryParams, bodyObj)
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
		{
			_parts := strings.Split("definition.subscribedSegments", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = integrationsUpdateCmd_definitionSubscribedSegments
		}
		{
			_parts := strings.Split("definition.subscribedUserProperties", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = integrationsUpdateCmd_definitionSubscribedUserProperties
		}
		{
			_parts := strings.Split("definition.type", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = integrationsUpdateCmd_definitionType
		}
		bodyMap["enabled"] = integrationsUpdateCmd_enabled
		bodyMap["name"] = integrationsUpdateCmd_name
		bodyMap["workspaceId"] = integrationsUpdateCmd_workspaceId
		resp, err := c.Do("PUT", "/api/integrations/", pathParams, queryParams, bodyMap)
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
	integrationsCmd.AddCommand(integrationsUpdateCmd)
	integrationsUpdateCmd.Flags().StringVar(&integrationsUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	integrationsUpdateCmd.Flags().StringVar(&integrationsUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	integrationsUpdateCmd.Flags().StringArrayVar(&integrationsUpdateCmd_definitionSubscribedSegments, "definition.subscribedSegments", nil, "")
	integrationsUpdateCmd.Flags().StringArrayVar(&integrationsUpdateCmd_definitionSubscribedUserProperties, "definition.subscribedUserProperties", nil, "")
	integrationsUpdateCmd.Flags().StringVar(&integrationsUpdateCmd_definitionType, "definition.type", "", "(Sync)")
	integrationsUpdateCmd.RegisterFlagCompletionFunc("definition.type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"Sync"}, cobra.ShellCompDirectiveNoFileComp
	})
	integrationsUpdateCmd.Flags().BoolVar(&integrationsUpdateCmd_enabled, "enabled", false, "")
	integrationsUpdateCmd.Flags().StringVar(&integrationsUpdateCmd_name, "name", "", "")
	integrationsUpdateCmd.Flags().StringVar(&integrationsUpdateCmd_workspaceId, "workspaceId", "", "")
	integrationsUpdateCmd.MarkFlagRequired("definition.subscribedSegments")
	integrationsUpdateCmd.MarkFlagRequired("definition.subscribedUserProperties")
	integrationsUpdateCmd.MarkFlagRequired("definition.type")
	integrationsUpdateCmd.MarkFlagRequired("name")
	integrationsUpdateCmd.MarkFlagRequired("workspaceId")
}
