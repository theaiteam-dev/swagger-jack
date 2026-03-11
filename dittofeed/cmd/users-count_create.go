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
	usersCountCreateCmdBody string
	usersCountCreateCmdBodyFile string
	usersCountCreateCmd_cursor string
	usersCountCreateCmd_direction string
	usersCountCreateCmd_includeSubscriptions bool
	usersCountCreateCmd_limit string
	usersCountCreateCmd_sortOrder string
	usersCountCreateCmd_subscriptionGroupFilter []string
	usersCountCreateCmd_userIds []string
	usersCountCreateCmd_userPropertyFilter []string
	usersCountCreateCmd_exclusiveCursor bool
	usersCountCreateCmd_segmentFilter []string
	usersCountCreateCmd_sortBy string
	usersCountCreateCmd_workspaceId string
)

var usersCountCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if usersCountCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(usersCountCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			usersCountCreateCmdBody = string(fileData)
		}
		if usersCountCreateCmdBody != "" {
			if !json.Valid([]byte(usersCountCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(usersCountCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/users/count", pathParams, queryParams, bodyObj)
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
		bodyMap["cursor"] = usersCountCreateCmd_cursor
		bodyMap["direction"] = usersCountCreateCmd_direction
		bodyMap["includeSubscriptions"] = usersCountCreateCmd_includeSubscriptions
		bodyMap["limit"] = usersCountCreateCmd_limit
		bodyMap["sortOrder"] = usersCountCreateCmd_sortOrder
		bodyMap["subscriptionGroupFilter"] = usersCountCreateCmd_subscriptionGroupFilter
		bodyMap["userIds"] = usersCountCreateCmd_userIds
		bodyMap["userPropertyFilter"] = usersCountCreateCmd_userPropertyFilter
		bodyMap["exclusiveCursor"] = usersCountCreateCmd_exclusiveCursor
		bodyMap["segmentFilter"] = usersCountCreateCmd_segmentFilter
		bodyMap["sortBy"] = usersCountCreateCmd_sortBy
		bodyMap["workspaceId"] = usersCountCreateCmd_workspaceId
		resp, err := c.Do("POST", "/api/users/count", pathParams, queryParams, bodyMap)
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
	usersCountCmd.AddCommand(usersCountCreateCmd)
	usersCountCreateCmd.Flags().StringVar(&usersCountCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	usersCountCreateCmd.Flags().StringVar(&usersCountCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	usersCountCreateCmd.Flags().StringVar(&usersCountCreateCmd_cursor, "cursor", "", "")
	usersCountCreateCmd.Flags().StringVar(&usersCountCreateCmd_direction, "direction", "", "")
	usersCountCreateCmd.Flags().BoolVar(&usersCountCreateCmd_includeSubscriptions, "includeSubscriptions", false, "")
	usersCountCreateCmd.Flags().StringVar(&usersCountCreateCmd_limit, "limit", "", "")
	usersCountCreateCmd.Flags().StringVar(&usersCountCreateCmd_sortOrder, "sortOrder", "", "")
	usersCountCreateCmd.Flags().StringArrayVar(&usersCountCreateCmd_subscriptionGroupFilter, "subscriptionGroupFilter", nil, "")
	usersCountCreateCmd.Flags().StringArrayVar(&usersCountCreateCmd_userIds, "userIds", nil, "")
	usersCountCreateCmd.Flags().StringArrayVar(&usersCountCreateCmd_userPropertyFilter, "userPropertyFilter", nil, "")
	usersCountCreateCmd.Flags().BoolVar(&usersCountCreateCmd_exclusiveCursor, "exclusiveCursor", false, "")
	usersCountCreateCmd.Flags().StringArrayVar(&usersCountCreateCmd_segmentFilter, "segmentFilter", nil, "")
	usersCountCreateCmd.Flags().StringVar(&usersCountCreateCmd_sortBy, "sortBy", "", "")
	usersCountCreateCmd.Flags().StringVar(&usersCountCreateCmd_workspaceId, "workspaceId", "", "")
	usersCountCreateCmd.MarkFlagRequired("workspaceId")
}
