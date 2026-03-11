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
	userPropertyIndicesUpdateCmdBody string
	userPropertyIndicesUpdateCmdBodyFile string
	userPropertyIndicesUpdateCmd_userPropertyId string
	userPropertyIndicesUpdateCmd_workspaceId string
	userPropertyIndicesUpdateCmd_type string
)

var userPropertyIndicesUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if userPropertyIndicesUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(userPropertyIndicesUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			userPropertyIndicesUpdateCmdBody = string(fileData)
		}
		if userPropertyIndicesUpdateCmdBody != "" {
			if !json.Valid([]byte(userPropertyIndicesUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(userPropertyIndicesUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/user-property-indices/", pathParams, queryParams, bodyObj)
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
		bodyMap["userPropertyId"] = userPropertyIndicesUpdateCmd_userPropertyId
		bodyMap["workspaceId"] = userPropertyIndicesUpdateCmd_workspaceId
		bodyMap["type"] = userPropertyIndicesUpdateCmd_type
		resp, err := c.Do("PUT", "/api/user-property-indices/", pathParams, queryParams, bodyMap)
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
	userPropertyIndicesCmd.AddCommand(userPropertyIndicesUpdateCmd)
	userPropertyIndicesUpdateCmd.Flags().StringVar(&userPropertyIndicesUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	userPropertyIndicesUpdateCmd.Flags().StringVar(&userPropertyIndicesUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	userPropertyIndicesUpdateCmd.Flags().StringVar(&userPropertyIndicesUpdateCmd_userPropertyId, "userPropertyId", "", "")
	userPropertyIndicesUpdateCmd.Flags().StringVar(&userPropertyIndicesUpdateCmd_workspaceId, "workspaceId", "", "")
	userPropertyIndicesUpdateCmd.Flags().StringVar(&userPropertyIndicesUpdateCmd_type, "type", "", "")
	userPropertyIndicesUpdateCmd.MarkFlagRequired("userPropertyId")
	userPropertyIndicesUpdateCmd.MarkFlagRequired("workspaceId")
	userPropertyIndicesUpdateCmd.MarkFlagRequired("type")
}
