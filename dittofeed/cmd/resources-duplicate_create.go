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
	resourcesDuplicateCreateCmdBody string
	resourcesDuplicateCreateCmdBodyFile string
	resourcesDuplicateCreateCmd_name string
	resourcesDuplicateCreateCmd_resourceType string
	resourcesDuplicateCreateCmd_workspaceId string
)

var resourcesDuplicateCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if resourcesDuplicateCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(resourcesDuplicateCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			resourcesDuplicateCreateCmdBody = string(fileData)
		}
		if resourcesDuplicateCreateCmdBody != "" {
			if !json.Valid([]byte(resourcesDuplicateCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(resourcesDuplicateCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/resources/duplicate", pathParams, queryParams, bodyObj)
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
		bodyMap["name"] = resourcesDuplicateCreateCmd_name
		bodyMap["resourceType"] = resourcesDuplicateCreateCmd_resourceType
		bodyMap["workspaceId"] = resourcesDuplicateCreateCmd_workspaceId
		resp, err := c.Do("POST", "/api/resources/duplicate", pathParams, queryParams, bodyMap)
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
	resourcesDuplicateCmd.AddCommand(resourcesDuplicateCreateCmd)
	resourcesDuplicateCreateCmd.Flags().StringVar(&resourcesDuplicateCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	resourcesDuplicateCreateCmd.Flags().StringVar(&resourcesDuplicateCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	resourcesDuplicateCreateCmd.Flags().StringVar(&resourcesDuplicateCreateCmd_name, "name", "", "")
	resourcesDuplicateCreateCmd.Flags().StringVar(&resourcesDuplicateCreateCmd_resourceType, "resourceType", "", "")
	resourcesDuplicateCreateCmd.Flags().StringVar(&resourcesDuplicateCreateCmd_workspaceId, "workspaceId", "", "")
	resourcesDuplicateCreateCmd.MarkFlagRequired("name")
	resourcesDuplicateCreateCmd.MarkFlagRequired("resourceType")
	resourcesDuplicateCreateCmd.MarkFlagRequired("workspaceId")
}
