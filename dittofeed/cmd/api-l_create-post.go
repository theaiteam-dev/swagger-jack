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
	apiLCreatePostCmdBody string
	apiLCreatePostCmdBodyFile string
	apiLCreatePostCmd_broadcastId string
	apiLCreatePostCmd_workspaceId string
)

var apiLCreatePostCmd = &cobra.Command{
	Use: "create-post",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if apiLCreatePostCmdBodyFile != "" {
			fileData, err := os.ReadFile(apiLCreatePostCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			apiLCreatePostCmdBody = string(fileData)
		}
		if apiLCreatePostCmdBody != "" {
			if !json.Valid([]byte(apiLCreatePostCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(apiLCreatePostCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api-l/embedded/broadcasts/start", pathParams, queryParams, bodyObj)
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
		bodyMap["broadcastId"] = apiLCreatePostCmd_broadcastId
		bodyMap["workspaceId"] = apiLCreatePostCmd_workspaceId
		resp, err := c.Do("POST", "/api-l/embedded/broadcasts/start", pathParams, queryParams, bodyMap)
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
	apiLCmd.AddCommand(apiLCreatePostCmd)
	apiLCreatePostCmd.Flags().StringVar(&apiLCreatePostCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	apiLCreatePostCmd.Flags().StringVar(&apiLCreatePostCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	apiLCreatePostCmd.Flags().StringVar(&apiLCreatePostCmd_broadcastId, "broadcastId", "", "")
	apiLCreatePostCmd.Flags().StringVar(&apiLCreatePostCmd_workspaceId, "workspaceId", "", "")
	apiLCreatePostCmd.MarkFlagRequired("broadcastId")
	apiLCreatePostCmd.MarkFlagRequired("workspaceId")
}
