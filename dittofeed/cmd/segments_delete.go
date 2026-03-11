package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	segmentsDeleteCmd_id string
	segmentsDeleteCmd_workspaceId string
)

var segmentsDeleteCmd = &cobra.Command{
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
		bodyMap["id"] = segmentsDeleteCmd_id
		bodyMap["workspaceId"] = segmentsDeleteCmd_workspaceId
		resp, err := c.Do("DELETE", "/api/segments/", pathParams, queryParams, bodyMap)
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
	segmentsCmd.AddCommand(segmentsDeleteCmd)
	segmentsDeleteCmd.Flags().StringVar(&segmentsDeleteCmd_id, "id", "", "")
	segmentsDeleteCmd.Flags().StringVar(&segmentsDeleteCmd_workspaceId, "workspaceId", "", "")
	segmentsDeleteCmd.MarkFlagRequired("id")
	segmentsDeleteCmd.MarkFlagRequired("workspaceId")
}
