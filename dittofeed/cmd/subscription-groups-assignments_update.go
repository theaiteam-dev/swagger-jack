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
	subscriptionGroupsAssignmentsUpdateCmdBody string
	subscriptionGroupsAssignmentsUpdateCmdBodyFile string
	subscriptionGroupsAssignmentsUpdateCmd_userUpdates []string
	subscriptionGroupsAssignmentsUpdateCmd_workspaceId string
)

var subscriptionGroupsAssignmentsUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if subscriptionGroupsAssignmentsUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(subscriptionGroupsAssignmentsUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			subscriptionGroupsAssignmentsUpdateCmdBody = string(fileData)
		}
		if subscriptionGroupsAssignmentsUpdateCmdBody != "" {
			if !json.Valid([]byte(subscriptionGroupsAssignmentsUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(subscriptionGroupsAssignmentsUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/subscription-groups/assignments", pathParams, queryParams, bodyObj)
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
		bodyMap["userUpdates"] = subscriptionGroupsAssignmentsUpdateCmd_userUpdates
		bodyMap["workspaceId"] = subscriptionGroupsAssignmentsUpdateCmd_workspaceId
		resp, err := c.Do("PUT", "/api/subscription-groups/assignments", pathParams, queryParams, bodyMap)
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
	subscriptionGroupsAssignmentsCmd.AddCommand(subscriptionGroupsAssignmentsUpdateCmd)
	subscriptionGroupsAssignmentsUpdateCmd.Flags().StringVar(&subscriptionGroupsAssignmentsUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	subscriptionGroupsAssignmentsUpdateCmd.Flags().StringVar(&subscriptionGroupsAssignmentsUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	subscriptionGroupsAssignmentsUpdateCmd.Flags().StringArrayVar(&subscriptionGroupsAssignmentsUpdateCmd_userUpdates, "userUpdates", nil, "")
	subscriptionGroupsAssignmentsUpdateCmd.Flags().StringVar(&subscriptionGroupsAssignmentsUpdateCmd_workspaceId, "workspaceId", "", "")
	subscriptionGroupsAssignmentsUpdateCmd.MarkFlagRequired("userUpdates")
	subscriptionGroupsAssignmentsUpdateCmd.MarkFlagRequired("workspaceId")
}
