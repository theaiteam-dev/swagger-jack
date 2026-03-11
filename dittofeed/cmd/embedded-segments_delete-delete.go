package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	embeddedSegmentsDeleteDeleteCmd_workspaceId string
	embeddedSegmentsDeleteDeleteCmd_id string
)

var embeddedSegmentsDeleteDeleteCmd = &cobra.Command{
	Use: "delete-delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = embeddedSegmentsDeleteDeleteCmd_workspaceId
		queryParams["id"] = embeddedSegmentsDeleteDeleteCmd_id
		resp, err := c.Do("DELETE", "/api-l/embedded/segments/v2", pathParams, queryParams, nil)
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
	embeddedSegmentsCmd.AddCommand(embeddedSegmentsDeleteDeleteCmd)
	embeddedSegmentsDeleteDeleteCmd.Flags().StringVar(&embeddedSegmentsDeleteDeleteCmd_workspaceId, "workspaceId", "", "")
	embeddedSegmentsDeleteDeleteCmd.Flags().StringVar(&embeddedSegmentsDeleteDeleteCmd_id, "id", "", "")
	embeddedSegmentsDeleteDeleteCmd.MarkFlagRequired("workspaceId")
	embeddedSegmentsDeleteDeleteCmd.MarkFlagRequired("id")
}
