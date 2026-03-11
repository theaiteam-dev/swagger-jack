package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	embeddedBroadcastsListGetCmd_workspaceId string
)

var embeddedBroadcastsListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = embeddedBroadcastsListGetCmd_workspaceId
		resp, err := c.Do("GET", "/api-l/embedded/broadcasts/gmail-authorization", pathParams, queryParams, nil)
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
	embeddedBroadcastsCmd.AddCommand(embeddedBroadcastsListGetCmd)
	embeddedBroadcastsListGetCmd.Flags().StringVar(&embeddedBroadcastsListGetCmd_workspaceId, "workspaceId", "", "")
	embeddedBroadcastsListGetCmd.MarkFlagRequired("workspaceId")
}
