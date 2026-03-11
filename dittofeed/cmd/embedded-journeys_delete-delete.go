package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	embeddedJourneysDeleteDeleteCmd_workspaceId string
	embeddedJourneysDeleteDeleteCmd_id string
)

var embeddedJourneysDeleteDeleteCmd = &cobra.Command{
	Use: "delete-delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = embeddedJourneysDeleteDeleteCmd_workspaceId
		queryParams["id"] = embeddedJourneysDeleteDeleteCmd_id
		resp, err := c.Do("DELETE", "/api-l/embedded/journeys/v2", pathParams, queryParams, nil)
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
	embeddedJourneysCmd.AddCommand(embeddedJourneysDeleteDeleteCmd)
	embeddedJourneysDeleteDeleteCmd.Flags().StringVar(&embeddedJourneysDeleteDeleteCmd_workspaceId, "workspaceId", "", "")
	embeddedJourneysDeleteDeleteCmd.Flags().StringVar(&embeddedJourneysDeleteDeleteCmd_id, "id", "", "")
	embeddedJourneysDeleteDeleteCmd.MarkFlagRequired("workspaceId")
	embeddedJourneysDeleteDeleteCmd.MarkFlagRequired("id")
}
