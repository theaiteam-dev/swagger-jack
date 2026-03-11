package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
	"dittofeed/internal/validate"
)

var (
	settingsDataSourcesDeleteCmd_workspaceId string
	settingsDataSourcesDeleteCmd_type string
)

var settingsDataSourcesDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = settingsDataSourcesDeleteCmd_workspaceId
		queryParams["type"] = settingsDataSourcesDeleteCmd_type
		if err := validate.Enum("type", settingsDataSourcesDeleteCmd_type, []string{"SegmentIO"}); err != nil { return err }
		resp, err := c.Do("DELETE", "/api/settings/data-sources", pathParams, queryParams, nil)
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
	settingsDataSourcesCmd.AddCommand(settingsDataSourcesDeleteCmd)
	settingsDataSourcesDeleteCmd.Flags().StringVar(&settingsDataSourcesDeleteCmd_workspaceId, "workspaceId", "", "")
	settingsDataSourcesDeleteCmd.Flags().StringVar(&settingsDataSourcesDeleteCmd_type, "type", "", "(SegmentIO)")
	settingsDataSourcesDeleteCmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"SegmentIO"}, cobra.ShellCompDirectiveNoFileComp
	})
	settingsDataSourcesDeleteCmd.MarkFlagRequired("workspaceId")
	settingsDataSourcesDeleteCmd.MarkFlagRequired("type")
}
