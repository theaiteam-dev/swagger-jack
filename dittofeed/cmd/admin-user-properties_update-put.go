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
	adminUserPropertiesUpdatePutCmdBody string
	adminUserPropertiesUpdatePutCmdBodyFile string
)

var adminUserPropertiesUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminUserPropertiesUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminUserPropertiesUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminUserPropertiesUpdatePutCmdBody = string(fileData)
		}
		if adminUserPropertiesUpdatePutCmdBody != "" {
			if !json.Valid([]byte(adminUserPropertiesUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminUserPropertiesUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/admin/user-properties/", pathParams, queryParams, bodyObj)
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
		resp, err := c.Do("PUT", "/api/admin/user-properties/", pathParams, queryParams, nil)
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
	adminUserPropertiesCmd.AddCommand(adminUserPropertiesUpdatePutCmd)
	adminUserPropertiesUpdatePutCmd.Flags().StringVar(&adminUserPropertiesUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminUserPropertiesUpdatePutCmd.Flags().StringVar(&adminUserPropertiesUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
}
