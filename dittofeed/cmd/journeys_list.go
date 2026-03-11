package cmd

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	journeysListCmd_workspaceId string
	journeysListCmd_getPartial bool
	journeysListCmd_resourceType string
)

var journeysListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = journeysListCmd_workspaceId
		queryParams["getPartial"] = strconv.FormatBool(journeysListCmd_getPartial)
		journeysListCmd_ids_vals, _ := cmd.Flags().GetStringArray("ids")
		queryParams["ids"] = strings.Join(journeysListCmd_ids_vals, ",")
		queryParams["resourceType"] = journeysListCmd_resourceType
		resp, err := c.Do("GET", "/api/journeys/", pathParams, queryParams, nil)
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
	journeysCmd.AddCommand(journeysListCmd)
	journeysListCmd.Flags().StringVar(&journeysListCmd_workspaceId, "workspaceId", "", "")
	journeysListCmd.Flags().BoolVar(&journeysListCmd_getPartial, "getPartial", false, "")
	journeysListCmd.Flags().StringArray("ids", nil, "")
	journeysListCmd.Flags().StringVar(&journeysListCmd_resourceType, "resourceType", "", "")
	journeysListCmd.MarkFlagRequired("workspaceId")
}
