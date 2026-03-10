package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	internalApiLListGetCmd_workspaceId string
	internalApiLListGetCmd_externalId string
)

var internalApiLListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = internalApiLListGetCmd_workspaceId
		queryParams["externalId"] = internalApiLListGetCmd_externalId
		resp, err := c.Do("GET", "/internal-api-l/workspaces/", pathParams, queryParams, nil)
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
	internalApiLCmd.AddCommand(internalApiLListGetCmd)
	internalApiLListGetCmd.Flags().StringVar(&internalApiLListGetCmd_workspaceId, "workspaceId", "", "")
	internalApiLListGetCmd.Flags().StringVar(&internalApiLListGetCmd_externalId, "externalId", "", "")
	internalApiLListGetCmd.MarkFlagRequired("workspaceId")
	internalApiLListGetCmd.MarkFlagRequired("externalId")
}
