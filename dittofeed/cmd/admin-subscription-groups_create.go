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
	adminSubscriptionGroupsCreateCmdBody string
	adminSubscriptionGroupsCreateCmdBodyFile string
)

var adminSubscriptionGroupsCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminSubscriptionGroupsCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminSubscriptionGroupsCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminSubscriptionGroupsCreateCmdBody = string(fileData)
		}
		if adminSubscriptionGroupsCreateCmdBody != "" {
			if !json.Valid([]byte(adminSubscriptionGroupsCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminSubscriptionGroupsCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/admin/subscription-groups/upload-csv", pathParams, queryParams, bodyObj)
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
		resp, err := c.Do("POST", "/api/admin/subscription-groups/upload-csv", pathParams, queryParams, nil)
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
	adminSubscriptionGroupsCmd.AddCommand(adminSubscriptionGroupsCreateCmd)
	adminSubscriptionGroupsCreateCmd.Flags().StringVar(&adminSubscriptionGroupsCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminSubscriptionGroupsCreateCmd.Flags().StringVar(&adminSubscriptionGroupsCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
}
