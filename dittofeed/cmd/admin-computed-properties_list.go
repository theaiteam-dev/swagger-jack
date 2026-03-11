package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminComputedPropertiesListCmd_workspaceId string
	adminComputedPropertiesListCmd_step string
)

var adminComputedPropertiesListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminComputedPropertiesListCmd_workspaceId
		queryParams["step"] = adminComputedPropertiesListCmd_step
		resp, err := c.Do("GET", "/api/admin/computed-properties/periods", pathParams, queryParams, nil)
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
	adminComputedPropertiesCmd.AddCommand(adminComputedPropertiesListCmd)
	adminComputedPropertiesListCmd.Flags().StringVar(&adminComputedPropertiesListCmd_workspaceId, "workspaceId", "", "")
	adminComputedPropertiesListCmd.Flags().StringVar(&adminComputedPropertiesListCmd_step, "step", "", "")
	adminComputedPropertiesListCmd.MarkFlagRequired("workspaceId")
	adminComputedPropertiesListCmd.MarkFlagRequired("step")
}
