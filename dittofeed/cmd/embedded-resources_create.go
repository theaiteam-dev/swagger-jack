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
	embeddedResourcesCreateCmdBody string
	embeddedResourcesCreateCmdBodyFile string
	embeddedResourcesCreateCmd_workspaceId string
	embeddedResourcesCreateCmd_name string
	embeddedResourcesCreateCmd_resourceType string
)

var embeddedResourcesCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if embeddedResourcesCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedResourcesCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedResourcesCreateCmdBody = string(fileData)
		}
		if embeddedResourcesCreateCmdBody != "" {
			if !json.Valid([]byte(embeddedResourcesCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedResourcesCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api-l/embedded/resources/duplicate", pathParams, queryParams, bodyObj)
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
		bodyMap["workspaceId"] = embeddedResourcesCreateCmd_workspaceId
		bodyMap["name"] = embeddedResourcesCreateCmd_name
		bodyMap["resourceType"] = embeddedResourcesCreateCmd_resourceType
		resp, err := c.Do("POST", "/api-l/embedded/resources/duplicate", pathParams, queryParams, bodyMap)
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
	embeddedResourcesCmd.AddCommand(embeddedResourcesCreateCmd)
	embeddedResourcesCreateCmd.Flags().StringVar(&embeddedResourcesCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedResourcesCreateCmd.Flags().StringVar(&embeddedResourcesCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	embeddedResourcesCreateCmd.Flags().StringVar(&embeddedResourcesCreateCmd_workspaceId, "workspaceId", "", "")
	embeddedResourcesCreateCmd.Flags().StringVar(&embeddedResourcesCreateCmd_name, "name", "", "")
	embeddedResourcesCreateCmd.Flags().StringVar(&embeddedResourcesCreateCmd_resourceType, "resourceType", "", "")
	embeddedResourcesCreateCmd.MarkFlagRequired("workspaceId")
	embeddedResourcesCreateCmd.MarkFlagRequired("name")
	embeddedResourcesCreateCmd.MarkFlagRequired("resourceType")
}
