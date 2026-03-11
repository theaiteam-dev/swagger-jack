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
	adminContentCreatePostCmdBody string
	adminContentCreatePostCmdBodyFile string
)

var adminContentCreatePostCmd = &cobra.Command{
	Use: "create-post",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminContentCreatePostCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminContentCreatePostCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminContentCreatePostCmdBody = string(fileData)
		}
		if adminContentCreatePostCmdBody != "" {
			if !json.Valid([]byte(adminContentCreatePostCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminContentCreatePostCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/admin/content/templates/test", pathParams, queryParams, bodyObj)
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
		resp, err := c.Do("POST", "/api/admin/content/templates/test", pathParams, queryParams, nil)
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
	adminContentCmd.AddCommand(adminContentCreatePostCmd)
	adminContentCreatePostCmd.Flags().StringVar(&adminContentCreatePostCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminContentCreatePostCmd.Flags().StringVar(&adminContentCreatePostCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
}
