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
	stripeBillingPortalCreateCmdBody string
	stripeBillingPortalCreateCmdBodyFile string
	stripeBillingPortalCreateCmd_returnUrl string
	stripeBillingPortalCreateCmd_workspaceId string
)

var stripeBillingPortalCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if stripeBillingPortalCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(stripeBillingPortalCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			stripeBillingPortalCreateCmdBody = string(fileData)
		}
		if stripeBillingPortalCreateCmdBody != "" {
			if !json.Valid([]byte(stripeBillingPortalCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(stripeBillingPortalCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api-l/stripe/billing-portal", pathParams, queryParams, bodyObj)
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
		bodyMap["returnUrl"] = stripeBillingPortalCreateCmd_returnUrl
		bodyMap["workspaceId"] = stripeBillingPortalCreateCmd_workspaceId
		resp, err := c.Do("POST", "/api-l/stripe/billing-portal", pathParams, queryParams, bodyMap)
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
	stripeBillingPortalCmd.AddCommand(stripeBillingPortalCreateCmd)
	stripeBillingPortalCreateCmd.Flags().StringVar(&stripeBillingPortalCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	stripeBillingPortalCreateCmd.Flags().StringVar(&stripeBillingPortalCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	stripeBillingPortalCreateCmd.Flags().StringVar(&stripeBillingPortalCreateCmd_returnUrl, "returnUrl", "", "")
	stripeBillingPortalCreateCmd.Flags().StringVar(&stripeBillingPortalCreateCmd_workspaceId, "workspaceId", "", "")
	stripeBillingPortalCreateCmd.MarkFlagRequired("returnUrl")
	stripeBillingPortalCreateCmd.MarkFlagRequired("workspaceId")
}
