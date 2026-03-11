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
	contentTemplatesUpdatePutCmdBody string
	contentTemplatesUpdatePutCmdBodyFile string
	contentTemplatesUpdatePutCmd_emailContentsType string
	contentTemplatesUpdatePutCmd_journeyMetadataJourneyId string
	contentTemplatesUpdatePutCmd_journeyMetadataNodeId string
	contentTemplatesUpdatePutCmd_name string
	contentTemplatesUpdatePutCmd_type string
	contentTemplatesUpdatePutCmd_workspaceId string
)

var contentTemplatesUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if contentTemplatesUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(contentTemplatesUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			contentTemplatesUpdatePutCmdBody = string(fileData)
		}
		if contentTemplatesUpdatePutCmdBody != "" {
			if !json.Valid([]byte(contentTemplatesUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(contentTemplatesUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/content/templates/reset", pathParams, queryParams, bodyObj)
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
		bodyMap["emailContentsType"] = contentTemplatesUpdatePutCmd_emailContentsType
		{
			_parts := strings.Split("journeyMetadata.journeyId", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = contentTemplatesUpdatePutCmd_journeyMetadataJourneyId
		}
		{
			_parts := strings.Split("journeyMetadata.nodeId", ".")
			_cur := bodyMap
			for _, _p := range _parts[:len(_parts)-1] {
				if _, ok := _cur[_p]; !ok {
					_cur[_p] = map[string]interface{}{}
				}
				_cur = _cur[_p].(map[string]interface{})
			}
			_cur[_parts[len(_parts)-1]] = contentTemplatesUpdatePutCmd_journeyMetadataNodeId
		}
		bodyMap["name"] = contentTemplatesUpdatePutCmd_name
		bodyMap["type"] = contentTemplatesUpdatePutCmd_type
		bodyMap["workspaceId"] = contentTemplatesUpdatePutCmd_workspaceId
		resp, err := c.Do("PUT", "/api/content/templates/reset", pathParams, queryParams, bodyMap)
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
	contentTemplatesCmd.AddCommand(contentTemplatesUpdatePutCmd)
	contentTemplatesUpdatePutCmd.Flags().StringVar(&contentTemplatesUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	contentTemplatesUpdatePutCmd.Flags().StringVar(&contentTemplatesUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	contentTemplatesUpdatePutCmd.Flags().StringVar(&contentTemplatesUpdatePutCmd_emailContentsType, "emailContentsType", "", "")
	contentTemplatesUpdatePutCmd.Flags().StringVar(&contentTemplatesUpdatePutCmd_journeyMetadataJourneyId, "journeyMetadata.journeyId", "", "")
	contentTemplatesUpdatePutCmd.Flags().StringVar(&contentTemplatesUpdatePutCmd_journeyMetadataNodeId, "journeyMetadata.nodeId", "", "")
	contentTemplatesUpdatePutCmd.Flags().StringVar(&contentTemplatesUpdatePutCmd_name, "name", "", "")
	contentTemplatesUpdatePutCmd.Flags().StringVar(&contentTemplatesUpdatePutCmd_type, "type", "", "")
	contentTemplatesUpdatePutCmd.Flags().StringVar(&contentTemplatesUpdatePutCmd_workspaceId, "workspaceId", "", "")
	contentTemplatesUpdatePutCmd.MarkFlagRequired("journeyMetadata.journeyId")
	contentTemplatesUpdatePutCmd.MarkFlagRequired("journeyMetadata.nodeId")
	contentTemplatesUpdatePutCmd.MarkFlagRequired("name")
	contentTemplatesUpdatePutCmd.MarkFlagRequired("type")
	contentTemplatesUpdatePutCmd.MarkFlagRequired("workspaceId")
}
