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
	resourcesListCmd_workspaceId string
	resourcesListCmd_segments bool
	resourcesListCmd_userProperties bool
	resourcesListCmd_subscriptionGroups bool
	resourcesListCmd_broadcasts bool
	resourcesListCmd_journeys string
	resourcesListCmd_messageTemplates bool
)

var resourcesListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = resourcesListCmd_workspaceId
		queryParams["segments"] = strconv.FormatBool(resourcesListCmd_segments)
		queryParams["userProperties"] = strconv.FormatBool(resourcesListCmd_userProperties)
		queryParams["subscriptionGroups"] = strconv.FormatBool(resourcesListCmd_subscriptionGroups)
		queryParams["broadcasts"] = strconv.FormatBool(resourcesListCmd_broadcasts)
		queryParams["journeys"] = resourcesListCmd_journeys
		queryParams["messageTemplates"] = strconv.FormatBool(resourcesListCmd_messageTemplates)
		resp, err := c.Do("GET", "/api/resources/", pathParams, queryParams, nil)
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
	resourcesCmd.AddCommand(resourcesListCmd)
	resourcesListCmd.Flags().StringVar(&resourcesListCmd_workspaceId, "workspaceId", "", "")
	resourcesListCmd.Flags().BoolVar(&resourcesListCmd_segments, "segments", false, "")
	resourcesListCmd.Flags().BoolVar(&resourcesListCmd_userProperties, "userProperties", false, "")
	resourcesListCmd.Flags().BoolVar(&resourcesListCmd_subscriptionGroups, "subscriptionGroups", false, "")
	resourcesListCmd.Flags().BoolVar(&resourcesListCmd_broadcasts, "broadcasts", false, "")
	resourcesListCmd.Flags().StringVar(&resourcesListCmd_journeys, "journeys", "", "")
	resourcesListCmd.Flags().BoolVar(&resourcesListCmd_messageTemplates, "messageTemplates", false, "")
	resourcesListCmd.MarkFlagRequired("workspaceId")
}
