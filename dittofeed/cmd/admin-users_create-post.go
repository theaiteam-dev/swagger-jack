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
	adminUsersCreatePostCmdBody string
	adminUsersCreatePostCmdBodyFile string
	adminUsersCreatePostCmd_direction string
	adminUsersCreatePostCmd_exclusiveCursor bool
	adminUsersCreatePostCmd_segmentFilter []string
	adminUsersCreatePostCmd_sortOrder string
	adminUsersCreatePostCmd_subscriptionGroupFilter []string
	adminUsersCreatePostCmd_userPropertyFilter []string
	adminUsersCreatePostCmd_cursor string
	adminUsersCreatePostCmd_includeSubscriptions bool
	adminUsersCreatePostCmd_limit string
	adminUsersCreatePostCmd_sortBy string
	adminUsersCreatePostCmd_userIds []string
	adminUsersCreatePostCmd_workspaceId string
)

var adminUsersCreatePostCmd = &cobra.Command{
	Use: "create-post",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminUsersCreatePostCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminUsersCreatePostCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminUsersCreatePostCmdBody = string(fileData)
		}
		if adminUsersCreatePostCmdBody != "" {
			if !json.Valid([]byte(adminUsersCreatePostCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminUsersCreatePostCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/admin/users/count", pathParams, queryParams, bodyObj)
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
		bodyMap["direction"] = adminUsersCreatePostCmd_direction
		bodyMap["exclusiveCursor"] = adminUsersCreatePostCmd_exclusiveCursor
		bodyMap["segmentFilter"] = adminUsersCreatePostCmd_segmentFilter
		bodyMap["sortOrder"] = adminUsersCreatePostCmd_sortOrder
		bodyMap["subscriptionGroupFilter"] = adminUsersCreatePostCmd_subscriptionGroupFilter
		bodyMap["userPropertyFilter"] = adminUsersCreatePostCmd_userPropertyFilter
		bodyMap["cursor"] = adminUsersCreatePostCmd_cursor
		bodyMap["includeSubscriptions"] = adminUsersCreatePostCmd_includeSubscriptions
		bodyMap["limit"] = adminUsersCreatePostCmd_limit
		bodyMap["sortBy"] = adminUsersCreatePostCmd_sortBy
		bodyMap["userIds"] = adminUsersCreatePostCmd_userIds
		bodyMap["workspaceId"] = adminUsersCreatePostCmd_workspaceId
		resp, err := c.Do("POST", "/api/admin/users/count", pathParams, queryParams, bodyMap)
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
	adminUsersCmd.AddCommand(adminUsersCreatePostCmd)
	adminUsersCreatePostCmd.Flags().StringVar(&adminUsersCreatePostCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminUsersCreatePostCmd.Flags().StringVar(&adminUsersCreatePostCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	adminUsersCreatePostCmd.Flags().StringVar(&adminUsersCreatePostCmd_direction, "direction", "", "")
	adminUsersCreatePostCmd.Flags().BoolVar(&adminUsersCreatePostCmd_exclusiveCursor, "exclusiveCursor", false, "")
	adminUsersCreatePostCmd.Flags().StringArrayVar(&adminUsersCreatePostCmd_segmentFilter, "segmentFilter", nil, "")
	adminUsersCreatePostCmd.Flags().StringVar(&adminUsersCreatePostCmd_sortOrder, "sortOrder", "", "")
	adminUsersCreatePostCmd.Flags().StringArrayVar(&adminUsersCreatePostCmd_subscriptionGroupFilter, "subscriptionGroupFilter", nil, "")
	adminUsersCreatePostCmd.Flags().StringArrayVar(&adminUsersCreatePostCmd_userPropertyFilter, "userPropertyFilter", nil, "")
	adminUsersCreatePostCmd.Flags().StringVar(&adminUsersCreatePostCmd_cursor, "cursor", "", "")
	adminUsersCreatePostCmd.Flags().BoolVar(&adminUsersCreatePostCmd_includeSubscriptions, "includeSubscriptions", false, "")
	adminUsersCreatePostCmd.Flags().StringVar(&adminUsersCreatePostCmd_limit, "limit", "", "")
	adminUsersCreatePostCmd.Flags().StringVar(&adminUsersCreatePostCmd_sortBy, "sortBy", "", "")
	adminUsersCreatePostCmd.Flags().StringArrayVar(&adminUsersCreatePostCmd_userIds, "userIds", nil, "")
	adminUsersCreatePostCmd.Flags().StringVar(&adminUsersCreatePostCmd_workspaceId, "workspaceId", "", "")
	adminUsersCreatePostCmd.MarkFlagRequired("workspaceId")
}
