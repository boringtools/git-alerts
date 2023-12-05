package cmd

import (
	"strconv"

	"github.com/boringtools/git-alerts/common"
	"github.com/boringtools/git-alerts/gh"
	"github.com/boringtools/git-alerts/logger"
	"github.com/boringtools/git-alerts/secrets"
	"github.com/spf13/cobra"
)

var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Scan with secrets detection",
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
		secrets.GetSecrets()

		logger.LogP("Scan ended : ", common.GetTime())
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)
}
