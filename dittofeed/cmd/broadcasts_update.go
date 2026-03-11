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
	broadcastsUpdateCmdBody string
	broadcastsUpdateCmdBodyFile string
	broadcastsUpdateCmd_id string
	broadcastsUpdateCmd_name string
	broadcastsUpdateCmd_workspaceId string
)

var broadcastsUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if broadcastsUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(broadcastsUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			broadcastsUpdateCmdBody = string(fileData)
		}
		if broadcastsUpdateCmdBody != "" {
			if !json.Valid([]byte(broadcastsUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(broadcastsUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/broadcasts/", pathParams, queryParams, bodyObj)
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
		bodyMap["id"] = broadcastsUpdateCmd_id
		bodyMap["name"] = broadcastsUpdateCmd_name
		bodyMap["workspaceId"] = broadcastsUpdateCmd_workspaceId
		resp, err := c.Do("PUT", "/api/broadcasts/", pathParams, queryParams, bodyMap)
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
	broadcastsCmd.AddCommand(broadcastsUpdateCmd)
	broadcastsUpdateCmd.Flags().StringVar(&broadcastsUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	broadcastsUpdateCmd.Flags().StringVar(&broadcastsUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	broadcastsUpdateCmd.Flags().StringVar(&broadcastsUpdateCmd_id, "id", "", "")
	broadcastsUpdateCmd.Flags().StringVar(&broadcastsUpdateCmd_name, "name", "", "")
	broadcastsUpdateCmd.Flags().StringVar(&broadcastsUpdateCmd_workspaceId, "workspaceId", "", "")
	broadcastsUpdateCmd.MarkFlagRequired("id")
	broadcastsUpdateCmd.MarkFlagRequired("workspaceId")
}
