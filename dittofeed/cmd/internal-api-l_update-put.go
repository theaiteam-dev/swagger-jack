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
	internalApiLUpdatePutCmdBody string
	internalApiLUpdatePutCmdBodyFile string
	internalApiLUpdatePutCmd_domain string
	internalApiLUpdatePutCmd_externalId string
	internalApiLUpdatePutCmd_name string
)

var internalApiLUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if internalApiLUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(internalApiLUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			internalApiLUpdatePutCmdBody = string(fileData)
		}
		if internalApiLUpdatePutCmdBody != "" {
			if !json.Valid([]byte(internalApiLUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(internalApiLUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/internal-api-l/workspaces/", pathParams, queryParams, bodyObj)
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
		bodyMap["domain"] = internalApiLUpdatePutCmd_domain
		bodyMap["externalId"] = internalApiLUpdatePutCmd_externalId
		bodyMap["name"] = internalApiLUpdatePutCmd_name
		resp, err := c.Do("PUT", "/internal-api-l/workspaces/", pathParams, queryParams, bodyMap)
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
	internalApiLCmd.AddCommand(internalApiLUpdatePutCmd)
	internalApiLUpdatePutCmd.Flags().StringVar(&internalApiLUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	internalApiLUpdatePutCmd.Flags().StringVar(&internalApiLUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	internalApiLUpdatePutCmd.Flags().StringVar(&internalApiLUpdatePutCmd_domain, "domain", "", "")
	internalApiLUpdatePutCmd.Flags().StringVar(&internalApiLUpdatePutCmd_externalId, "externalId", "", "")
	internalApiLUpdatePutCmd.Flags().StringVar(&internalApiLUpdatePutCmd_name, "name", "", "")
	internalApiLUpdatePutCmd.MarkFlagRequired("name")
}
