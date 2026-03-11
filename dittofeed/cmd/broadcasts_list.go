package cmd

import (
	"fmt"
	"os"
	"strings"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	broadcastsListCmd_workspaceId string
)

var broadcastsListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = broadcastsListCmd_workspaceId
		broadcastsListCmd_ids_vals, _ := cmd.Flags().GetStringArray("ids")
		queryParams["ids"] = strings.Join(broadcastsListCmd_ids_vals, ",")
		resp, err := c.Do("GET", "/api/broadcasts/", pathParams, queryParams, nil)
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
	broadcastsCmd.AddCommand(broadcastsListCmd)
	broadcastsListCmd.Flags().StringVar(&broadcastsListCmd_workspaceId, "workspaceId", "", "")
	broadcastsListCmd.Flags().StringArray("ids", nil, "")
	broadcastsListCmd.MarkFlagRequired("workspaceId")
}
