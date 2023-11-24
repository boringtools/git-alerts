package cmd

import (
	"os"
	"strconv"

	"github.com/boringtools/git-alerts/common"
	"github.com/boringtools/git-alerts/gh"

	"github.com/boringtools/git-alerts/config"
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

		_, errScanFile := os.Stat(config.GhFilePaths()[1])

		if errScanFile != nil {
			logger.LogERR("Previous scan files not found")
			logger.LogERR("If you are running it for the first time")
			logger.LogERR("Please consider running the SCAN command first")
			os.Exit(1)
		}
		gh.Connecter()
		reporter.Notify()
		logger.LogP("Scan ended : ", common.GetTime())
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.PersistentFlags().BoolVarP(&slack, "slack-alert", "s", false, "Slack notification")
}
