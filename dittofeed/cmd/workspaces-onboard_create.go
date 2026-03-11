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
	workspacesOnboardCreateCmdBody string
	workspacesOnboardCreateCmdBodyFile string
	workspacesOnboardCreateCmd_name string
	workspacesOnboardCreateCmd_domain string
	workspacesOnboardCreateCmd_externalId string
)

var workspacesOnboardCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if workspacesOnboardCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(workspacesOnboardCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			workspacesOnboardCreateCmdBody = string(fileData)
		}
		if workspacesOnboardCreateCmdBody != "" {
			if !json.Valid([]byte(workspacesOnboardCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(workspacesOnboardCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api-l/workspaces/onboard", pathParams, queryParams, bodyObj)
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
		bodyMap["name"] = workspacesOnboardCreateCmd_name
		bodyMap["domain"] = workspacesOnboardCreateCmd_domain
		bodyMap["externalId"] = workspacesOnboardCreateCmd_externalId
		resp, err := c.Do("POST", "/api-l/workspaces/onboard", pathParams, queryParams, bodyMap)
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
	workspacesOnboardCmd.AddCommand(workspacesOnboardCreateCmd)
	workspacesOnboardCreateCmd.Flags().StringVar(&workspacesOnboardCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	workspacesOnboardCreateCmd.Flags().StringVar(&workspacesOnboardCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	workspacesOnboardCreateCmd.Flags().StringVar(&workspacesOnboardCreateCmd_name, "name", "", "")
	workspacesOnboardCreateCmd.Flags().StringVar(&workspacesOnboardCreateCmd_domain, "domain", "", "")
	workspacesOnboardCreateCmd.Flags().StringVar(&workspacesOnboardCreateCmd_externalId, "externalId", "", "")
	workspacesOnboardCreateCmd.MarkFlagRequired("name")
}
