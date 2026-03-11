package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	journeysV2DeleteCmd_workspaceId string
	journeysV2DeleteCmd_id string
)

var journeysV2DeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = journeysV2DeleteCmd_workspaceId
		queryParams["id"] = journeysV2DeleteCmd_id
		resp, err := c.Do("DELETE", "/api/journeys/v2", pathParams, queryParams, nil)
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
	journeysV2Cmd.AddCommand(journeysV2DeleteCmd)
	journeysV2DeleteCmd.Flags().StringVar(&journeysV2DeleteCmd_workspaceId, "workspaceId", "", "")
	journeysV2DeleteCmd.Flags().StringVar(&journeysV2DeleteCmd_id, "id", "", "")
	journeysV2DeleteCmd.MarkFlagRequired("workspaceId")
	journeysV2DeleteCmd.MarkFlagRequired("id")
}
