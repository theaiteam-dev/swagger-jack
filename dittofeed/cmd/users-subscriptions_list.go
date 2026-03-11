package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	usersSubscriptionsListCmd_workspaceId string
	usersSubscriptionsListCmd_userId string
)

var usersSubscriptionsListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = usersSubscriptionsListCmd_workspaceId
		queryParams["userId"] = usersSubscriptionsListCmd_userId
		resp, err := c.Do("GET", "/api/users/subscriptions", pathParams, queryParams, nil)
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
	usersSubscriptionsCmd.AddCommand(usersSubscriptionsListCmd)
	usersSubscriptionsListCmd.Flags().StringVar(&usersSubscriptionsListCmd_workspaceId, "workspaceId", "", "")
	usersSubscriptionsListCmd.Flags().StringVar(&usersSubscriptionsListCmd_userId, "userId", "", "")
	usersSubscriptionsListCmd.MarkFlagRequired("workspaceId")
	usersSubscriptionsListCmd.MarkFlagRequired("userId")
}
