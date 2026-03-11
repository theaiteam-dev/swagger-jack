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
	journeysStatsListCmd_workspaceId string
)

var journeysStatsListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = journeysStatsListCmd_workspaceId
		journeysStatsListCmd_journeyIds_vals, _ := cmd.Flags().GetStringArray("journeyIds")
		queryParams["journeyIds"] = strings.Join(journeysStatsListCmd_journeyIds_vals, ",")
		resp, err := c.Do("GET", "/api/journeys/stats", pathParams, queryParams, nil)
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
	journeysStatsCmd.AddCommand(journeysStatsListCmd)
	journeysStatsListCmd.Flags().StringVar(&journeysStatsListCmd_workspaceId, "workspaceId", "", "")
	journeysStatsListCmd.Flags().StringArray("journeyIds", nil, "")
	journeysStatsListCmd.MarkFlagRequired("workspaceId")
}
