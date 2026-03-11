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
	embeddedContentListCmd_workspaceId string
	embeddedContentListCmd_resourceType string
)

var embeddedContentListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = embeddedContentListCmd_workspaceId
		embeddedContentListCmd_ids_vals, _ := cmd.Flags().GetStringArray("ids")
		queryParams["ids"] = strings.Join(embeddedContentListCmd_ids_vals, ",")
		queryParams["resourceType"] = embeddedContentListCmd_resourceType
		resp, err := c.Do("GET", "/api-l/embedded/content/templates", pathParams, queryParams, nil)
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
	embeddedContentCmd.AddCommand(embeddedContentListCmd)
	embeddedContentListCmd.Flags().StringVar(&embeddedContentListCmd_workspaceId, "workspaceId", "", "")
	embeddedContentListCmd.Flags().StringArray("ids", nil, "")
	embeddedContentListCmd.Flags().StringVar(&embeddedContentListCmd_resourceType, "resourceType", "", "")
	embeddedContentListCmd.MarkFlagRequired("workspaceId")
}
