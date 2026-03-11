package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	computedPropertiesPeriodsListCmd_workspaceId string
	computedPropertiesPeriodsListCmd_step string
)

var computedPropertiesPeriodsListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = computedPropertiesPeriodsListCmd_workspaceId
		queryParams["step"] = computedPropertiesPeriodsListCmd_step
		resp, err := c.Do("GET", "/api/computed-properties/periods", pathParams, queryParams, nil)
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
	computedPropertiesPeriodsCmd.AddCommand(computedPropertiesPeriodsListCmd)
	computedPropertiesPeriodsListCmd.Flags().StringVar(&computedPropertiesPeriodsListCmd_workspaceId, "workspaceId", "", "")
	computedPropertiesPeriodsListCmd.Flags().StringVar(&computedPropertiesPeriodsListCmd_step, "step", "", "")
	computedPropertiesPeriodsListCmd.MarkFlagRequired("workspaceId")
	computedPropertiesPeriodsListCmd.MarkFlagRequired("step")
}
