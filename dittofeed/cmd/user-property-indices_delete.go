package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	userPropertyIndicesDeleteCmd_userPropertyId string
	userPropertyIndicesDeleteCmd_workspaceId string
)

var userPropertyIndicesDeleteCmd = &cobra.Command{
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
		bodyMap["userPropertyId"] = userPropertyIndicesDeleteCmd_userPropertyId
		bodyMap["workspaceId"] = userPropertyIndicesDeleteCmd_workspaceId
		resp, err := c.Do("DELETE", "/api/user-property-indices/", pathParams, queryParams, bodyMap)
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
	userPropertyIndicesCmd.AddCommand(userPropertyIndicesDeleteCmd)
	userPropertyIndicesDeleteCmd.Flags().StringVar(&userPropertyIndicesDeleteCmd_userPropertyId, "userPropertyId", "", "")
	userPropertyIndicesDeleteCmd.Flags().StringVar(&userPropertyIndicesDeleteCmd_workspaceId, "workspaceId", "", "")
	userPropertyIndicesDeleteCmd.MarkFlagRequired("userPropertyId")
	userPropertyIndicesDeleteCmd.MarkFlagRequired("workspaceId")
}
