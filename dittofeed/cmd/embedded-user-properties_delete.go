package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	embeddedUserPropertiesDeleteCmd_id string
	embeddedUserPropertiesDeleteCmd_workspaceId string
)

var embeddedUserPropertiesDeleteCmd = &cobra.Command{
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
		bodyMap["id"] = embeddedUserPropertiesDeleteCmd_id
		bodyMap["workspaceId"] = embeddedUserPropertiesDeleteCmd_workspaceId
		resp, err := c.Do("DELETE", "/api-l/embedded/user-properties/", pathParams, queryParams, bodyMap)
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
	embeddedUserPropertiesCmd.AddCommand(embeddedUserPropertiesDeleteCmd)
	embeddedUserPropertiesDeleteCmd.Flags().StringVar(&embeddedUserPropertiesDeleteCmd_id, "id", "", "")
	embeddedUserPropertiesDeleteCmd.Flags().StringVar(&embeddedUserPropertiesDeleteCmd_workspaceId, "workspaceId", "", "")
	embeddedUserPropertiesDeleteCmd.MarkFlagRequired("id")
	embeddedUserPropertiesDeleteCmd.MarkFlagRequired("workspaceId")
}
