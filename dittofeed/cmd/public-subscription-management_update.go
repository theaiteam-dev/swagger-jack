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
	publicSubscriptionManagementUpdateCmdBody string
	publicSubscriptionManagementUpdateCmdBodyFile string
)

var publicSubscriptionManagementUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if publicSubscriptionManagementUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(publicSubscriptionManagementUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			publicSubscriptionManagementUpdateCmdBody = string(fileData)
		}
		if publicSubscriptionManagementUpdateCmdBody != "" {
			if !json.Valid([]byte(publicSubscriptionManagementUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(publicSubscriptionManagementUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/public/subscription-management/user-subscriptions", pathParams, queryParams, bodyObj)
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
		resp, err := c.Do("PUT", "/api/public/subscription-management/user-subscriptions", pathParams, queryParams, nil)
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
	publicSubscriptionManagementCmd.AddCommand(publicSubscriptionManagementUpdateCmd)
	publicSubscriptionManagementUpdateCmd.Flags().StringVar(&publicSubscriptionManagementUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	publicSubscriptionManagementUpdateCmd.Flags().StringVar(&publicSubscriptionManagementUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
}
