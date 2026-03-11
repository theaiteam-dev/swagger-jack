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
	segmentsListCmd_workspaceId string
	segmentsListCmd_resourceType string
)

var segmentsListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = segmentsListCmd_workspaceId
		segmentsListCmd_ids_vals, _ := cmd.Flags().GetStringArray("ids")
		queryParams["ids"] = strings.Join(segmentsListCmd_ids_vals, ",")
		queryParams["resourceType"] = segmentsListCmd_resourceType
		resp, err := c.Do("GET", "/api/segments/", pathParams, queryParams, nil)
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
	segmentsCmd.AddCommand(segmentsListCmd)
	segmentsListCmd.Flags().StringVar(&segmentsListCmd_workspaceId, "workspaceId", "", "")
	segmentsListCmd.Flags().StringArray("ids", nil, "")
	segmentsListCmd.Flags().StringVar(&segmentsListCmd_resourceType, "resourceType", "", "")
	segmentsListCmd.MarkFlagRequired("workspaceId")
}
