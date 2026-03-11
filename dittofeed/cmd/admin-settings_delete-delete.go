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
	adminSettingsDeleteDeleteCmd_workspaceId string
	adminSettingsDeleteDeleteCmd_type string
)

var adminSettingsDeleteDeleteCmd = &cobra.Command{
	Use: "delete-delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminSettingsDeleteDeleteCmd_workspaceId
		queryParams["type"] = adminSettingsDeleteDeleteCmd_type
		if err := validate.Enum("type", adminSettingsDeleteDeleteCmd_type, []string{"SegmentIO"}); err != nil { return err }
		resp, err := c.Do("DELETE", "/api/admin/settings/data-sources", pathParams, queryParams, nil)
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
	adminSettingsCmd.AddCommand(adminSettingsDeleteDeleteCmd)
	adminSettingsDeleteDeleteCmd.Flags().StringVar(&adminSettingsDeleteDeleteCmd_workspaceId, "workspaceId", "", "")
	adminSettingsDeleteDeleteCmd.Flags().StringVar(&adminSettingsDeleteDeleteCmd_type, "type", "", "(SegmentIO)")
	adminSettingsDeleteDeleteCmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"SegmentIO"}, cobra.ShellCompDirectiveNoFileComp
	})
	adminSettingsDeleteDeleteCmd.MarkFlagRequired("workspaceId")
	adminSettingsDeleteDeleteCmd.MarkFlagRequired("type")
}
