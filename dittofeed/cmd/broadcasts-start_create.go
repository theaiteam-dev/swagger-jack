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
	broadcastsStartCreateCmdBody string
	broadcastsStartCreateCmdBodyFile string
	broadcastsStartCreateCmd_workspaceId string
	broadcastsStartCreateCmd_broadcastId string
)

var broadcastsStartCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if broadcastsStartCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(broadcastsStartCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			broadcastsStartCreateCmdBody = string(fileData)
		}
		if broadcastsStartCreateCmdBody != "" {
			if !json.Valid([]byte(broadcastsStartCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(broadcastsStartCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/broadcasts/start", pathParams, queryParams, bodyObj)
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
		bodyMap["workspaceId"] = broadcastsStartCreateCmd_workspaceId
		bodyMap["broadcastId"] = broadcastsStartCreateCmd_broadcastId
		resp, err := c.Do("POST", "/api/broadcasts/start", pathParams, queryParams, bodyMap)
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
	broadcastsStartCmd.AddCommand(broadcastsStartCreateCmd)
	broadcastsStartCreateCmd.Flags().StringVar(&broadcastsStartCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	broadcastsStartCreateCmd.Flags().StringVar(&broadcastsStartCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	broadcastsStartCreateCmd.Flags().StringVar(&broadcastsStartCreateCmd_workspaceId, "workspaceId", "", "")
	broadcastsStartCreateCmd.Flags().StringVar(&broadcastsStartCreateCmd_broadcastId, "broadcastId", "", "")
	broadcastsStartCreateCmd.MarkFlagRequired("workspaceId")
	broadcastsStartCreateCmd.MarkFlagRequired("broadcastId")
}
