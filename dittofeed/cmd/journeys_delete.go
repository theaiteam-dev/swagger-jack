package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	journeysDeleteCmd_workspaceId string
	journeysDeleteCmd_id string
)

var journeysDeleteCmd = &cobra.Command{
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
		bodyMap["workspaceId"] = journeysDeleteCmd_workspaceId
		bodyMap["id"] = journeysDeleteCmd_id
		resp, err := c.Do("DELETE", "/api/journeys/", pathParams, queryParams, bodyMap)
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
	journeysCmd.AddCommand(journeysDeleteCmd)
	journeysDeleteCmd.Flags().StringVar(&journeysDeleteCmd_workspaceId, "workspaceId", "", "")
	journeysDeleteCmd.Flags().StringVar(&journeysDeleteCmd_id, "id", "", "")
	journeysDeleteCmd.MarkFlagRequired("workspaceId")
	journeysDeleteCmd.MarkFlagRequired("id")
}
