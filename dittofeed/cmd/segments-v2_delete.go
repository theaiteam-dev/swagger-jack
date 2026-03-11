package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	segmentsV2DeleteCmd_workspaceId string
	segmentsV2DeleteCmd_id string
)

var segmentsV2DeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = segmentsV2DeleteCmd_workspaceId
		queryParams["id"] = segmentsV2DeleteCmd_id
		resp, err := c.Do("DELETE", "/api/segments/v2", pathParams, queryParams, nil)
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
	segmentsV2Cmd.AddCommand(segmentsV2DeleteCmd)
	segmentsV2DeleteCmd.Flags().StringVar(&segmentsV2DeleteCmd_workspaceId, "workspaceId", "", "")
	segmentsV2DeleteCmd.Flags().StringVar(&segmentsV2DeleteCmd_id, "id", "", "")
	segmentsV2DeleteCmd.MarkFlagRequired("workspaceId")
	segmentsV2DeleteCmd.MarkFlagRequired("id")
}
