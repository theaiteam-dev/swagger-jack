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
	broadcastsPauseCreateCmdBody string
	broadcastsPauseCreateCmdBodyFile string
	broadcastsPauseCreateCmd_broadcastId string
	broadcastsPauseCreateCmd_workspaceId string
)

var broadcastsPauseCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if broadcastsPauseCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(broadcastsPauseCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			broadcastsPauseCreateCmdBody = string(fileData)
		}
		if broadcastsPauseCreateCmdBody != "" {
			if !json.Valid([]byte(broadcastsPauseCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(broadcastsPauseCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/broadcasts/pause", pathParams, queryParams, bodyObj)
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
		bodyMap["broadcastId"] = broadcastsPauseCreateCmd_broadcastId
		bodyMap["workspaceId"] = broadcastsPauseCreateCmd_workspaceId
		resp, err := c.Do("POST", "/api/broadcasts/pause", pathParams, queryParams, bodyMap)
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
	broadcastsPauseCmd.AddCommand(broadcastsPauseCreateCmd)
	broadcastsPauseCreateCmd.Flags().StringVar(&broadcastsPauseCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	broadcastsPauseCreateCmd.Flags().StringVar(&broadcastsPauseCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	broadcastsPauseCreateCmd.Flags().StringVar(&broadcastsPauseCreateCmd_broadcastId, "broadcastId", "", "")
	broadcastsPauseCreateCmd.Flags().StringVar(&broadcastsPauseCreateCmd_workspaceId, "workspaceId", "", "")
	broadcastsPauseCreateCmd.MarkFlagRequired("broadcastId")
	broadcastsPauseCreateCmd.MarkFlagRequired("workspaceId")
}
