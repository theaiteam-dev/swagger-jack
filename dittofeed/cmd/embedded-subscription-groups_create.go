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
	embeddedSubscriptionGroupsCreateCmdBody string
	embeddedSubscriptionGroupsCreateCmdBodyFile string
)

var embeddedSubscriptionGroupsCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if embeddedSubscriptionGroupsCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedSubscriptionGroupsCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedSubscriptionGroupsCreateCmdBody = string(fileData)
		}
		if embeddedSubscriptionGroupsCreateCmdBody != "" {
			if !json.Valid([]byte(embeddedSubscriptionGroupsCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedSubscriptionGroupsCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api-l/embedded/subscription-groups/upload-csv", pathParams, queryParams, bodyObj)
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
		resp, err := c.Do("POST", "/api-l/embedded/subscription-groups/upload-csv", pathParams, queryParams, nil)
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
	embeddedSubscriptionGroupsCmd.AddCommand(embeddedSubscriptionGroupsCreateCmd)
	embeddedSubscriptionGroupsCreateCmd.Flags().StringVar(&embeddedSubscriptionGroupsCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedSubscriptionGroupsCreateCmd.Flags().StringVar(&embeddedSubscriptionGroupsCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
}
