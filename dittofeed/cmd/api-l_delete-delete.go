package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	apiLDeleteDeleteCmd_workspaceId string
	apiLDeleteDeleteCmd_id string
	apiLDeleteDeleteCmd_type string
)

var apiLDeleteDeleteCmd = &cobra.Command{
	Use: "delete-delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = apiLDeleteDeleteCmd_workspaceId
		queryParams["id"] = apiLDeleteDeleteCmd_id
		queryParams["type"] = apiLDeleteDeleteCmd_type
		resp, err := c.Do("DELETE", "/api-l/embedded/content/templates/v2", pathParams, queryParams, nil)
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
	apiLCmd.AddCommand(apiLDeleteDeleteCmd)
	apiLDeleteDeleteCmd.Flags().StringVar(&apiLDeleteDeleteCmd_workspaceId, "workspaceId", "", "")
	apiLDeleteDeleteCmd.Flags().StringVar(&apiLDeleteDeleteCmd_id, "id", "", "")
	apiLDeleteDeleteCmd.Flags().StringVar(&apiLDeleteDeleteCmd_type, "type", "", "")
	apiLDeleteDeleteCmd.MarkFlagRequired("workspaceId")
	apiLDeleteDeleteCmd.MarkFlagRequired("id")
	apiLDeleteDeleteCmd.MarkFlagRequired("type")
}
