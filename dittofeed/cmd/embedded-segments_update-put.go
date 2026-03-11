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
	embeddedSegmentsUpdatePutCmdBody string
	embeddedSegmentsUpdatePutCmdBodyFile string
)

var embeddedSegmentsUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if embeddedSegmentsUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedSegmentsUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedSegmentsUpdatePutCmdBody = string(fileData)
		}
		if embeddedSegmentsUpdatePutCmdBody != "" {
			if !json.Valid([]byte(embeddedSegmentsUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedSegmentsUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api-l/embedded/segments/", pathParams, queryParams, bodyObj)
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
		resp, err := c.Do("PUT", "/api-l/embedded/segments/", pathParams, queryParams, nil)
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
	embeddedSegmentsCmd.AddCommand(embeddedSegmentsUpdatePutCmd)
	embeddedSegmentsUpdatePutCmd.Flags().StringVar(&embeddedSegmentsUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedSegmentsUpdatePutCmd.Flags().StringVar(&embeddedSegmentsUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
}
