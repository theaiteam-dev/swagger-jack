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
	sessionsCreateCmdBody string
	sessionsCreateCmdBodyFile string
	sessionsCreateCmd_occupantId string
	sessionsCreateCmd_workspaceId string
)

var sessionsCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if sessionsCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(sessionsCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			sessionsCreateCmdBody = string(fileData)
		}
		if sessionsCreateCmdBody != "" {
			if !json.Valid([]byte(sessionsCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(sessionsCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api-l/sessions/", pathParams, queryParams, bodyObj)
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
		bodyMap["occupantId"] = sessionsCreateCmd_occupantId
		bodyMap["workspaceId"] = sessionsCreateCmd_workspaceId
		resp, err := c.Do("POST", "/api-l/sessions/", pathParams, queryParams, bodyMap)
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
	sessionsCmd.AddCommand(sessionsCreateCmd)
	sessionsCreateCmd.Flags().StringVar(&sessionsCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	sessionsCreateCmd.Flags().StringVar(&sessionsCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	sessionsCreateCmd.Flags().StringVar(&sessionsCreateCmd_occupantId, "occupantId", "", "")
	sessionsCreateCmd.Flags().StringVar(&sessionsCreateCmd_workspaceId, "workspaceId", "", "")
	sessionsCreateCmd.MarkFlagRequired("workspaceId")
}
