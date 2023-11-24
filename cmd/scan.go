package cmd

import (
	"strconv"

	"github.com/boringtools/git-alerts/common"
	"github.com/boringtools/git-alerts/gh"
	"github.com/boringtools/git-alerts/logger"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan public repositories",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		envs := map[string]string{
			"org":     org,
			"rfp":     report,
			"command": cmd.Use,
			"csv":     strconv.FormatBool(csv),
		}

		common.SetEnvs(envs)
		gh.Connecter()

		logger.LogP("Scan ended : ", common.GetTime())
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
