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
	broadcastsCancelCreateCmdBody string
	broadcastsCancelCreateCmdBodyFile string
	broadcastsCancelCreateCmd_broadcastId string
	broadcastsCancelCreateCmd_workspaceId string
)

var broadcastsCancelCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if broadcastsCancelCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(broadcastsCancelCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			broadcastsCancelCreateCmdBody = string(fileData)
		}
		if broadcastsCancelCreateCmdBody != "" {
			if !json.Valid([]byte(broadcastsCancelCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(broadcastsCancelCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/broadcasts/cancel", pathParams, queryParams, bodyObj)
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
		bodyMap["broadcastId"] = broadcastsCancelCreateCmd_broadcastId
		bodyMap["workspaceId"] = broadcastsCancelCreateCmd_workspaceId
		resp, err := c.Do("POST", "/api/broadcasts/cancel", pathParams, queryParams, bodyMap)
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
	broadcastsCancelCmd.AddCommand(broadcastsCancelCreateCmd)
	broadcastsCancelCreateCmd.Flags().StringVar(&broadcastsCancelCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	broadcastsCancelCreateCmd.Flags().StringVar(&broadcastsCancelCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	broadcastsCancelCreateCmd.Flags().StringVar(&broadcastsCancelCreateCmd_broadcastId, "broadcastId", "", "")
	broadcastsCancelCreateCmd.Flags().StringVar(&broadcastsCancelCreateCmd_workspaceId, "workspaceId", "", "")
	broadcastsCancelCreateCmd.MarkFlagRequired("broadcastId")
	broadcastsCancelCreateCmd.MarkFlagRequired("workspaceId")
}
