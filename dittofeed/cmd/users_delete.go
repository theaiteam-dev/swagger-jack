package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	usersDeleteCmd_userIds []string
	usersDeleteCmd_workspaceId string
)

var usersDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		bodyMap := map[string]interface{}{}
		bodyMap["userIds"] = usersDeleteCmd_userIds
		bodyMap["workspaceId"] = usersDeleteCmd_workspaceId
		resp, err := c.Do("DELETE", "/api/users/", pathParams, queryParams, bodyMap)
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
	usersCmd.AddCommand(usersDeleteCmd)
	usersDeleteCmd.Flags().StringArrayVar(&usersDeleteCmd_userIds, "userIds", nil, "")
	usersDeleteCmd.Flags().StringVar(&usersDeleteCmd_workspaceId, "workspaceId", "", "")
	usersDeleteCmd.MarkFlagRequired("userIds")
	usersDeleteCmd.MarkFlagRequired("workspaceId")
}
