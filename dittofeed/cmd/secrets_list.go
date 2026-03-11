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
	secretsListCmd_workspaceId string
)

var secretsListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = secretsListCmd_workspaceId
		secretsListCmd_names_vals, _ := cmd.Flags().GetStringArray("names")
		queryParams["names"] = strings.Join(secretsListCmd_names_vals, ",")
		resp, err := c.Do("GET", "/api/secrets/", pathParams, queryParams, nil)
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
	secretsCmd.AddCommand(secretsListCmd)
	secretsListCmd.Flags().StringVar(&secretsListCmd_workspaceId, "workspaceId", "", "")
	secretsListCmd.Flags().StringArray("names", nil, "")
	secretsListCmd.MarkFlagRequired("workspaceId")
}
