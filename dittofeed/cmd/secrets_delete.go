package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	secretsDeleteCmd_id string
	secretsDeleteCmd_workspaceId string
)

var secretsDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["id"] = secretsDeleteCmd_id
		queryParams["workspaceId"] = secretsDeleteCmd_workspaceId
		resp, err := c.Do("DELETE", "/api/secrets/", pathParams, queryParams, nil)
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
	secretsCmd.AddCommand(secretsDeleteCmd)
	secretsDeleteCmd.Flags().StringVar(&secretsDeleteCmd_id, "id", "", "")
	secretsDeleteCmd.Flags().StringVar(&secretsDeleteCmd_workspaceId, "workspaceId", "", "")
	secretsDeleteCmd.MarkFlagRequired("id")
	secretsDeleteCmd.MarkFlagRequired("workspaceId")
}
