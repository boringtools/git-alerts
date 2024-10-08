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
			"org":                 org,
			"rfp":                 report,
			"command":             cmd.Use,
			"csv":                 strconv.FormatBool(csv),
			"trufflehog":          strconv.FormatBool(trufflehog),
			"trufflehog-verified": strconv.FormatBool(trufflehogVerified),
			"gitleaks":            strconv.FormatBool(gitleaks),
		}
		common.SetEnvs(envs)

		gh.Connecter()
		secrets.GetSecrets()

		logger.LogP("Scan ended : ", common.GetTime())
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)
	detectCmd.PersistentFlags().BoolVarP(&trufflehog, "trufflehog", "t", false, "Scan secrets using Trufflehog")
	detectCmd.PersistentFlags().BoolVarP(&trufflehogVerified, "trufflehog-verified", "v", true, "Scan Trufflehog verified secrets")
	detectCmd.PersistentFlags().BoolVarP(&gitleaks, "gitleaks", "g", false, "Scan secrets using Gitleaks")

	detectCmd.MarkFlagsOneRequired("trufflehog", "trufflehog-verified", "gitleaks")
}
