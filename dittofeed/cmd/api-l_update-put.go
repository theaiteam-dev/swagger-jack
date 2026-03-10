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
	apiLUpdatePutCmdBody string
	apiLUpdatePutCmdBodyFile string
	apiLUpdatePutCmd_userUpdates []string
	apiLUpdatePutCmd_workspaceId string
)

var apiLUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if apiLUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(apiLUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			apiLUpdatePutCmdBody = string(fileData)
		}
		if apiLUpdatePutCmdBody != "" {
			if !json.Valid([]byte(apiLUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(apiLUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api-l/embedded/subscription-groups/assignments", pathParams, queryParams, bodyObj)
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
		bodyMap["userUpdates"] = apiLUpdatePutCmd_userUpdates
		bodyMap["workspaceId"] = apiLUpdatePutCmd_workspaceId
		resp, err := c.Do("PUT", "/api-l/embedded/subscription-groups/assignments", pathParams, queryParams, bodyMap)
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
	apiLCmd.AddCommand(apiLUpdatePutCmd)
	apiLUpdatePutCmd.Flags().StringVar(&apiLUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	apiLUpdatePutCmd.Flags().StringVar(&apiLUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	apiLUpdatePutCmd.Flags().StringArrayVar(&apiLUpdatePutCmd_userUpdates, "userUpdates", nil, "")
	apiLUpdatePutCmd.Flags().StringVar(&apiLUpdatePutCmd_workspaceId, "workspaceId", "", "")
	apiLUpdatePutCmd.MarkFlagRequired("userUpdates")
	apiLUpdatePutCmd.MarkFlagRequired("workspaceId")
}
