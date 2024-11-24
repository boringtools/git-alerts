package cmd

import (
	"os"

	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/github"
	"github.com/boringtools/git-alerts/pkg/notification"
	"github.com/boringtools/git-alerts/pkg/ui"
	"github.com/boringtools/git-alerts/pkg/utils"

	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor public repositories",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		common.Command = cmd.Use

		if common.SlackNotification {
			_, isSlackHook := os.LookupEnv("SLACK_HOOK")

			if !isSlackHook {
				ui.PrintError("SLACK_HOOK is not configured in ENV variable")
				os.Exit(1)
			}
		}

		if !utils.IsPreviousScanFileExists() {
			ui.PrintError("Previous scan files not found, Please consider running SCAN command first")
			os.Exit(1)
		}

		github.Connecter()
		notification.Notify()
		ui.PrintSuccess("Scan Ended : %s", utils.GetCurrentTime())
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.PersistentFlags().BoolVarP(&common.SlackNotification, "slack-alert", "s", false, "Slack notification")
}
