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
	stripeCheckoutSessionCreateCmdBody string
	stripeCheckoutSessionCreateCmdBodyFile string
	stripeCheckoutSessionCreateCmd_cancelUrl string
	stripeCheckoutSessionCreateCmd_priceId string
	stripeCheckoutSessionCreateCmd_successUrl string
	stripeCheckoutSessionCreateCmd_workspaceId string
)

var stripeCheckoutSessionCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if stripeCheckoutSessionCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(stripeCheckoutSessionCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			stripeCheckoutSessionCreateCmdBody = string(fileData)
		}
		if stripeCheckoutSessionCreateCmdBody != "" {
			if !json.Valid([]byte(stripeCheckoutSessionCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(stripeCheckoutSessionCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api-l/stripe/checkout-session", pathParams, queryParams, bodyObj)
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
		bodyMap["cancelUrl"] = stripeCheckoutSessionCreateCmd_cancelUrl
		bodyMap["priceId"] = stripeCheckoutSessionCreateCmd_priceId
		bodyMap["successUrl"] = stripeCheckoutSessionCreateCmd_successUrl
		bodyMap["workspaceId"] = stripeCheckoutSessionCreateCmd_workspaceId
		resp, err := c.Do("POST", "/api-l/stripe/checkout-session", pathParams, queryParams, bodyMap)
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
	stripeCheckoutSessionCmd.AddCommand(stripeCheckoutSessionCreateCmd)
	stripeCheckoutSessionCreateCmd.Flags().StringVar(&stripeCheckoutSessionCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	stripeCheckoutSessionCreateCmd.Flags().StringVar(&stripeCheckoutSessionCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	stripeCheckoutSessionCreateCmd.Flags().StringVar(&stripeCheckoutSessionCreateCmd_cancelUrl, "cancelUrl", "", "")
	stripeCheckoutSessionCreateCmd.Flags().StringVar(&stripeCheckoutSessionCreateCmd_priceId, "priceId", "", "")
	stripeCheckoutSessionCreateCmd.Flags().StringVar(&stripeCheckoutSessionCreateCmd_successUrl, "successUrl", "", "")
	stripeCheckoutSessionCreateCmd.Flags().StringVar(&stripeCheckoutSessionCreateCmd_workspaceId, "workspaceId", "", "")
	stripeCheckoutSessionCreateCmd.MarkFlagRequired("cancelUrl")
	stripeCheckoutSessionCreateCmd.MarkFlagRequired("priceId")
	stripeCheckoutSessionCreateCmd.MarkFlagRequired("successUrl")
	stripeCheckoutSessionCreateCmd.MarkFlagRequired("workspaceId")
}
