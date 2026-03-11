package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	segmentsDownloadListCmd_workspaceId string
)

var segmentsDownloadListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = segmentsDownloadListCmd_workspaceId
		resp, err := c.Do("GET", "/api/segments/download", pathParams, queryParams, nil)
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
	segmentsDownloadCmd.AddCommand(segmentsDownloadListCmd)
	segmentsDownloadListCmd.Flags().StringVar(&segmentsDownloadListCmd_workspaceId, "workspaceId", "", "")
	segmentsDownloadListCmd.MarkFlagRequired("workspaceId")
}
