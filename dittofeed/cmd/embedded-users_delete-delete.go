package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	embeddedUsersDeleteDeleteCmd_userIds []string
	embeddedUsersDeleteDeleteCmd_workspaceId string
)

var embeddedUsersDeleteDeleteCmd = &cobra.Command{
	Use: "delete-delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		bodyMap := map[string]interface{}{}
		bodyMap["userIds"] = embeddedUsersDeleteDeleteCmd_userIds
		bodyMap["workspaceId"] = embeddedUsersDeleteDeleteCmd_workspaceId
		resp, err := c.Do("DELETE", "/api-l/embedded/users/", pathParams, queryParams, bodyMap)
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
	embeddedUsersCmd.AddCommand(embeddedUsersDeleteDeleteCmd)
	embeddedUsersDeleteDeleteCmd.Flags().StringArrayVar(&embeddedUsersDeleteDeleteCmd_userIds, "userIds", nil, "")
	embeddedUsersDeleteDeleteCmd.Flags().StringVar(&embeddedUsersDeleteDeleteCmd_workspaceId, "workspaceId", "", "")
	embeddedUsersDeleteDeleteCmd.MarkFlagRequired("userIds")
	embeddedUsersDeleteDeleteCmd.MarkFlagRequired("workspaceId")
}
