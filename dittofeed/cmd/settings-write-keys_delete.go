package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	settingsWriteKeysDeleteCmd_workspaceId string
	settingsWriteKeysDeleteCmd_writeKeyName string
)

var settingsWriteKeysDeleteCmd = &cobra.Command{
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
		bodyMap["workspaceId"] = settingsWriteKeysDeleteCmd_workspaceId
		bodyMap["writeKeyName"] = settingsWriteKeysDeleteCmd_writeKeyName
		resp, err := c.Do("DELETE", "/api/settings/write-keys", pathParams, queryParams, bodyMap)
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
	settingsWriteKeysCmd.AddCommand(settingsWriteKeysDeleteCmd)
	settingsWriteKeysDeleteCmd.Flags().StringVar(&settingsWriteKeysDeleteCmd_workspaceId, "workspaceId", "", "")
	settingsWriteKeysDeleteCmd.Flags().StringVar(&settingsWriteKeysDeleteCmd_writeKeyName, "writeKeyName", "", "")
	settingsWriteKeysDeleteCmd.MarkFlagRequired("workspaceId")
	settingsWriteKeysDeleteCmd.MarkFlagRequired("writeKeyName")
}
