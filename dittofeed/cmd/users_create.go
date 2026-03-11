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
	usersCreateCmdBody string
	usersCreateCmdBodyFile string
	usersCreateCmd_workspaceId string
	usersCreateCmd_direction string
	usersCreateCmd_includeSubscriptions bool
	usersCreateCmd_limit string
	usersCreateCmd_segmentFilter []string
	usersCreateCmd_subscriptionGroupFilter []string
	usersCreateCmd_userPropertyFilter []string
	usersCreateCmd_cursor string
	usersCreateCmd_exclusiveCursor bool
	usersCreateCmd_sortBy string
	usersCreateCmd_sortOrder string
	usersCreateCmd_userIds []string
)

var usersCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if usersCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(usersCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			usersCreateCmdBody = string(fileData)
		}
		if usersCreateCmdBody != "" {
			if !json.Valid([]byte(usersCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(usersCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/users/", pathParams, queryParams, bodyObj)
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
		bodyMap["workspaceId"] = usersCreateCmd_workspaceId
		bodyMap["direction"] = usersCreateCmd_direction
		bodyMap["includeSubscriptions"] = usersCreateCmd_includeSubscriptions
		bodyMap["limit"] = usersCreateCmd_limit
		bodyMap["segmentFilter"] = usersCreateCmd_segmentFilter
		bodyMap["subscriptionGroupFilter"] = usersCreateCmd_subscriptionGroupFilter
		bodyMap["userPropertyFilter"] = usersCreateCmd_userPropertyFilter
		bodyMap["cursor"] = usersCreateCmd_cursor
		bodyMap["exclusiveCursor"] = usersCreateCmd_exclusiveCursor
		bodyMap["sortBy"] = usersCreateCmd_sortBy
		bodyMap["sortOrder"] = usersCreateCmd_sortOrder
		bodyMap["userIds"] = usersCreateCmd_userIds
		resp, err := c.Do("POST", "/api/users/", pathParams, queryParams, bodyMap)
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
	usersCmd.AddCommand(usersCreateCmd)
	usersCreateCmd.Flags().StringVar(&usersCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	usersCreateCmd.Flags().StringVar(&usersCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	usersCreateCmd.Flags().StringVar(&usersCreateCmd_workspaceId, "workspaceId", "", "")
	usersCreateCmd.Flags().StringVar(&usersCreateCmd_direction, "direction", "", "")
	usersCreateCmd.Flags().BoolVar(&usersCreateCmd_includeSubscriptions, "includeSubscriptions", false, "")
	usersCreateCmd.Flags().StringVar(&usersCreateCmd_limit, "limit", "", "")
	usersCreateCmd.Flags().StringArrayVar(&usersCreateCmd_segmentFilter, "segmentFilter", nil, "")
	usersCreateCmd.Flags().StringArrayVar(&usersCreateCmd_subscriptionGroupFilter, "subscriptionGroupFilter", nil, "")
	usersCreateCmd.Flags().StringArrayVar(&usersCreateCmd_userPropertyFilter, "userPropertyFilter", nil, "")
	usersCreateCmd.Flags().StringVar(&usersCreateCmd_cursor, "cursor", "", "")
	usersCreateCmd.Flags().BoolVar(&usersCreateCmd_exclusiveCursor, "exclusiveCursor", false, "")
	usersCreateCmd.Flags().StringVar(&usersCreateCmd_sortBy, "sortBy", "", "")
	usersCreateCmd.Flags().StringVar(&usersCreateCmd_sortOrder, "sortOrder", "", "")
	usersCreateCmd.Flags().StringArrayVar(&usersCreateCmd_userIds, "userIds", nil, "")
	usersCreateCmd.MarkFlagRequired("workspaceId")
}
