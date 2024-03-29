package cmd

import (
	"strconv"

	"github.com/boringtools/git-alerts/common"
	"github.com/boringtools/git-alerts/gh"

	"github.com/boringtools/git-alerts/logger"
	"github.com/boringtools/git-alerts/reporter"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor public repositories",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		envs := map[string]string{
			"org":     org,
			"rfp":     report,
			"command": cmd.Use,
			"csv":     strconv.FormatBool(csv),
			"slack":   strconv.FormatBool(slack),
		}

		common.SetEnvs(envs)

		common.CheckScanFiles()
		gh.Connecter()
		reporter.Notify()
		logger.LogP("Scan ended : ", common.GetTime())
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.PersistentFlags().BoolVarP(&slack, "slack-alert", "s", false, "Slack notification")
}
