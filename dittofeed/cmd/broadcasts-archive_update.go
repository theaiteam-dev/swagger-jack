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
	broadcastsArchiveUpdateCmdBody string
	broadcastsArchiveUpdateCmdBodyFile string
	broadcastsArchiveUpdateCmd_archived bool
	broadcastsArchiveUpdateCmd_broadcastId string
	broadcastsArchiveUpdateCmd_workspaceId string
)

var broadcastsArchiveUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if broadcastsArchiveUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(broadcastsArchiveUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			broadcastsArchiveUpdateCmdBody = string(fileData)
		}
		if broadcastsArchiveUpdateCmdBody != "" {
			if !json.Valid([]byte(broadcastsArchiveUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(broadcastsArchiveUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/broadcasts/archive", pathParams, queryParams, bodyObj)
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
		bodyMap["archived"] = broadcastsArchiveUpdateCmd_archived
		bodyMap["broadcastId"] = broadcastsArchiveUpdateCmd_broadcastId
		bodyMap["workspaceId"] = broadcastsArchiveUpdateCmd_workspaceId
		resp, err := c.Do("PUT", "/api/broadcasts/archive", pathParams, queryParams, bodyMap)
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
	broadcastsArchiveCmd.AddCommand(broadcastsArchiveUpdateCmd)
	broadcastsArchiveUpdateCmd.Flags().StringVar(&broadcastsArchiveUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	broadcastsArchiveUpdateCmd.Flags().StringVar(&broadcastsArchiveUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	broadcastsArchiveUpdateCmd.Flags().BoolVar(&broadcastsArchiveUpdateCmd_archived, "archived", false, "")
	broadcastsArchiveUpdateCmd.Flags().StringVar(&broadcastsArchiveUpdateCmd_broadcastId, "broadcastId", "", "")
	broadcastsArchiveUpdateCmd.Flags().StringVar(&broadcastsArchiveUpdateCmd_workspaceId, "workspaceId", "", "")
	broadcastsArchiveUpdateCmd.MarkFlagRequired("broadcastId")
	broadcastsArchiveUpdateCmd.MarkFlagRequired("workspaceId")
}
