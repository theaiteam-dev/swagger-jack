package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	embeddedUserPropertiesListCmd_workspaceId string
)

var embeddedUserPropertiesListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = embeddedUserPropertiesListCmd_workspaceId
		resp, err := c.Do("GET", "/api-l/embedded/user-properties/", pathParams, queryParams, nil)
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
	embeddedUserPropertiesCmd.AddCommand(embeddedUserPropertiesListCmd)
	embeddedUserPropertiesListCmd.Flags().StringVar(&embeddedUserPropertiesListCmd_workspaceId, "workspaceId", "", "")
	embeddedUserPropertiesListCmd.MarkFlagRequired("workspaceId")
}
