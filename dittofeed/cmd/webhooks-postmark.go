package cmd

import "github.com/spf13/cobra"

var webhooksPostmarkCmd = &cobra.Command{
	Use: "webhooks-postmark",
	Short: "webhooks-postmark",
}

func init() {
	rootCmd.AddCommand(webhooksPostmarkCmd)
}
