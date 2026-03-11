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
	embeddedUsersCreatePostCmdBody string
	embeddedUsersCreatePostCmdBodyFile string
	embeddedUsersCreatePostCmd_cursor string
	embeddedUsersCreatePostCmd_direction string
	embeddedUsersCreatePostCmd_limit string
	embeddedUsersCreatePostCmd_segmentFilter []string
	embeddedUsersCreatePostCmd_sortOrder string
	embeddedUsersCreatePostCmd_subscriptionGroupFilter []string
	embeddedUsersCreatePostCmd_exclusiveCursor bool
	embeddedUsersCreatePostCmd_includeSubscriptions bool
	embeddedUsersCreatePostCmd_sortBy string
	embeddedUsersCreatePostCmd_userIds []string
	embeddedUsersCreatePostCmd_userPropertyFilter []string
	embeddedUsersCreatePostCmd_workspaceId string
)

var embeddedUsersCreatePostCmd = &cobra.Command{
	Use: "create-post",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if embeddedUsersCreatePostCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedUsersCreatePostCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedUsersCreatePostCmdBody = string(fileData)
		}
		if embeddedUsersCreatePostCmdBody != "" {
			if !json.Valid([]byte(embeddedUsersCreatePostCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedUsersCreatePostCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api-l/embedded/users/", pathParams, queryParams, bodyObj)
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
		bodyMap["cursor"] = embeddedUsersCreatePostCmd_cursor
		bodyMap["direction"] = embeddedUsersCreatePostCmd_direction
		bodyMap["limit"] = embeddedUsersCreatePostCmd_limit
		bodyMap["segmentFilter"] = embeddedUsersCreatePostCmd_segmentFilter
		bodyMap["sortOrder"] = embeddedUsersCreatePostCmd_sortOrder
		bodyMap["subscriptionGroupFilter"] = embeddedUsersCreatePostCmd_subscriptionGroupFilter
		bodyMap["exclusiveCursor"] = embeddedUsersCreatePostCmd_exclusiveCursor
		bodyMap["includeSubscriptions"] = embeddedUsersCreatePostCmd_includeSubscriptions
		bodyMap["sortBy"] = embeddedUsersCreatePostCmd_sortBy
		bodyMap["userIds"] = embeddedUsersCreatePostCmd_userIds
		bodyMap["userPropertyFilter"] = embeddedUsersCreatePostCmd_userPropertyFilter
		bodyMap["workspaceId"] = embeddedUsersCreatePostCmd_workspaceId
		resp, err := c.Do("POST", "/api-l/embedded/users/", pathParams, queryParams, bodyMap)
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
	embeddedUsersCmd.AddCommand(embeddedUsersCreatePostCmd)
	embeddedUsersCreatePostCmd.Flags().StringVar(&embeddedUsersCreatePostCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedUsersCreatePostCmd.Flags().StringVar(&embeddedUsersCreatePostCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	embeddedUsersCreatePostCmd.Flags().StringVar(&embeddedUsersCreatePostCmd_cursor, "cursor", "", "")
	embeddedUsersCreatePostCmd.Flags().StringVar(&embeddedUsersCreatePostCmd_direction, "direction", "", "")
	embeddedUsersCreatePostCmd.Flags().StringVar(&embeddedUsersCreatePostCmd_limit, "limit", "", "")
	embeddedUsersCreatePostCmd.Flags().StringArrayVar(&embeddedUsersCreatePostCmd_segmentFilter, "segmentFilter", nil, "")
	embeddedUsersCreatePostCmd.Flags().StringVar(&embeddedUsersCreatePostCmd_sortOrder, "sortOrder", "", "")
	embeddedUsersCreatePostCmd.Flags().StringArrayVar(&embeddedUsersCreatePostCmd_subscriptionGroupFilter, "subscriptionGroupFilter", nil, "")
	embeddedUsersCreatePostCmd.Flags().BoolVar(&embeddedUsersCreatePostCmd_exclusiveCursor, "exclusiveCursor", false, "")
	embeddedUsersCreatePostCmd.Flags().BoolVar(&embeddedUsersCreatePostCmd_includeSubscriptions, "includeSubscriptions", false, "")
	embeddedUsersCreatePostCmd.Flags().StringVar(&embeddedUsersCreatePostCmd_sortBy, "sortBy", "", "")
	embeddedUsersCreatePostCmd.Flags().StringArrayVar(&embeddedUsersCreatePostCmd_userIds, "userIds", nil, "")
	embeddedUsersCreatePostCmd.Flags().StringArrayVar(&embeddedUsersCreatePostCmd_userPropertyFilter, "userPropertyFilter", nil, "")
	embeddedUsersCreatePostCmd.Flags().StringVar(&embeddedUsersCreatePostCmd_workspaceId, "workspaceId", "", "")
	embeddedUsersCreatePostCmd.MarkFlagRequired("workspaceId")
}
