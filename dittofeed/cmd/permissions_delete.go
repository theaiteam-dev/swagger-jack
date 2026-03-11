package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	permissionsDeleteCmd_workspaceId string
	permissionsDeleteCmd_email string
)

var permissionsDeleteCmd = &cobra.Command{
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
		bodyMap["workspaceId"] = permissionsDeleteCmd_workspaceId
		bodyMap["email"] = permissionsDeleteCmd_email
		resp, err := c.Do("DELETE", "/api/permissions/", pathParams, queryParams, bodyMap)
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
	permissionsCmd.AddCommand(permissionsDeleteCmd)
	permissionsDeleteCmd.Flags().StringVar(&permissionsDeleteCmd_workspaceId, "workspaceId", "", "")
	permissionsDeleteCmd.Flags().StringVar(&permissionsDeleteCmd_email, "email", "", "")
	permissionsDeleteCmd.MarkFlagRequired("workspaceId")
	permissionsDeleteCmd.MarkFlagRequired("email")
}
