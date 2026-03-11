package cmd

import (
	"fmt"
	"os"
	"strconv"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	embeddedResourcesListCmd_workspaceId string
	embeddedResourcesListCmd_segments bool
	embeddedResourcesListCmd_userProperties bool
	embeddedResourcesListCmd_subscriptionGroups bool
	embeddedResourcesListCmd_broadcasts bool
	embeddedResourcesListCmd_journeys string
	embeddedResourcesListCmd_messageTemplates bool
)

var embeddedResourcesListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = embeddedResourcesListCmd_workspaceId
		queryParams["segments"] = strconv.FormatBool(embeddedResourcesListCmd_segments)
		queryParams["userProperties"] = strconv.FormatBool(embeddedResourcesListCmd_userProperties)
		queryParams["subscriptionGroups"] = strconv.FormatBool(embeddedResourcesListCmd_subscriptionGroups)
		queryParams["broadcasts"] = strconv.FormatBool(embeddedResourcesListCmd_broadcasts)
		queryParams["journeys"] = embeddedResourcesListCmd_journeys
		queryParams["messageTemplates"] = strconv.FormatBool(embeddedResourcesListCmd_messageTemplates)
		resp, err := c.Do("GET", "/api-l/embedded/resources/", pathParams, queryParams, nil)
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
	embeddedResourcesCmd.AddCommand(embeddedResourcesListCmd)
	embeddedResourcesListCmd.Flags().StringVar(&embeddedResourcesListCmd_workspaceId, "workspaceId", "", "")
	embeddedResourcesListCmd.Flags().BoolVar(&embeddedResourcesListCmd_segments, "segments", false, "")
	embeddedResourcesListCmd.Flags().BoolVar(&embeddedResourcesListCmd_userProperties, "userProperties", false, "")
	embeddedResourcesListCmd.Flags().BoolVar(&embeddedResourcesListCmd_subscriptionGroups, "subscriptionGroups", false, "")
	embeddedResourcesListCmd.Flags().BoolVar(&embeddedResourcesListCmd_broadcasts, "broadcasts", false, "")
	embeddedResourcesListCmd.Flags().StringVar(&embeddedResourcesListCmd_journeys, "journeys", "", "")
	embeddedResourcesListCmd.Flags().BoolVar(&embeddedResourcesListCmd_messageTemplates, "messageTemplates", false, "")
	embeddedResourcesListCmd.MarkFlagRequired("workspaceId")
}
