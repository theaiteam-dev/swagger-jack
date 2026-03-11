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
	embeddedContentCreatePostCmdBody string
	embeddedContentCreatePostCmdBodyFile string
)

var embeddedContentCreatePostCmd = &cobra.Command{
	Use: "create-post",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if embeddedContentCreatePostCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedContentCreatePostCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedContentCreatePostCmdBody = string(fileData)
		}
		if embeddedContentCreatePostCmdBody != "" {
			if !json.Valid([]byte(embeddedContentCreatePostCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedContentCreatePostCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api-l/embedded/content/templates/test", pathParams, queryParams, bodyObj)
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
		resp, err := c.Do("POST", "/api-l/embedded/content/templates/test", pathParams, queryParams, nil)
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
	embeddedContentCmd.AddCommand(embeddedContentCreatePostCmd)
	embeddedContentCreatePostCmd.Flags().StringVar(&embeddedContentCreatePostCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedContentCreatePostCmd.Flags().StringVar(&embeddedContentCreatePostCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
}
