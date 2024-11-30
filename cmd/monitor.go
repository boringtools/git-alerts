package cmd

import (
	"os"

	"github.com/boringtools/git-alerts/internal/ui"
	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/github"
	"github.com/boringtools/git-alerts/pkg/monitor"

	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor public repositories",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		common.Command = cmd.Use

		if !common.IsPreviousScanFileExists() {
			ui.PrintError("Previous scan files not found, Please consider running SCAN command first")
			os.Exit(1)
		}

		github.Connecter()
		count, errNewPublicRepos := github.GetNewPublicRepositories()

		if errNewPublicRepos != nil {
			ui.PrintError("error in finding new repositories: %s", errNewPublicRepos)
			os.Exit(1)
		}

		if len(count) != 0 {
			ui.PrintWarning("New public repositories detected")
			monitor.OutputNewRepository()
		} else {
			ui.PrintMsg("No new public repositories detected")
		}
		ui.PrintSuccess("Scan Ended : %s", common.GetCurrentTime())
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.PersistentFlags().BoolVarP(&common.SlackNotification, "slack-alert", "s", false, "Slack notification")
	monitorCmd.PersistentFlags().BoolVarP(&common.GitleaksScan, "gitleaks", "g", false, "Scan secrets using Gitleaks")
}
