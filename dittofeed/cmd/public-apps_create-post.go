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
	publicAppsCreatePostCmdBody string
	publicAppsCreatePostCmdBodyFile string
)

var publicAppsCreatePostCmd = &cobra.Command{
	Use: "create-post",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if publicAppsCreatePostCmdBodyFile != "" {
			fileData, err := os.ReadFile(publicAppsCreatePostCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			publicAppsCreatePostCmdBody = string(fileData)
		}
		if publicAppsCreatePostCmdBody != "" {
			if !json.Valid([]byte(publicAppsCreatePostCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(publicAppsCreatePostCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/public/apps/alias", pathParams, queryParams, bodyObj)
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
		resp, err := c.Do("POST", "/api/public/apps/alias", pathParams, queryParams, nil)
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
	publicAppsCmd.AddCommand(publicAppsCreatePostCmd)
	publicAppsCreatePostCmd.Flags().StringVar(&publicAppsCreatePostCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	publicAppsCreatePostCmd.Flags().StringVar(&publicAppsCreatePostCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
}
