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
	broadcastsV2UpdateCmdBody string
	broadcastsV2UpdateCmdBodyFile string
)

var broadcastsV2UpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if broadcastsV2UpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(broadcastsV2UpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			broadcastsV2UpdateCmdBody = string(fileData)
		}
		if broadcastsV2UpdateCmdBody != "" {
			if !json.Valid([]byte(broadcastsV2UpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(broadcastsV2UpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/broadcasts/v2", pathParams, queryParams, bodyObj)
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
		resp, err := c.Do("PUT", "/api/broadcasts/v2", pathParams, queryParams, nil)
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
	broadcastsV2Cmd.AddCommand(broadcastsV2UpdateCmd)
	broadcastsV2UpdateCmd.Flags().StringVar(&broadcastsV2UpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	broadcastsV2UpdateCmd.Flags().StringVar(&broadcastsV2UpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
}
