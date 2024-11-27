package cmd

import (
	"github.com/boringtools/git-alerts/internal/ui"
	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/github"
	"github.com/boringtools/git-alerts/pkg/secrets"
	"github.com/spf13/cobra"
)

var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Scan with secrets detection",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		common.Command = cmd.Use
		github.Connecter()
		secrets.RunSecretsScan()

		ui.PrintSuccess("Scan Ended : %s", common.GetCurrentTime())
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)
	detectCmd.PersistentFlags().BoolVarP(&common.TrufflehogScan, "trufflehog", "t", false, "Scan secrets using Trufflehog")
	detectCmd.PersistentFlags().BoolVarP(&common.TrufflehogVerifiedScan, "trufflehog-verified", "v", true, "Scan secrets using Trufflehog verified option")
	detectCmd.PersistentFlags().BoolVarP(&common.GitleaksScan, "gitleaks", "g", false, "Scan secrets using Gitleaks")

	detectCmd.MarkFlagsOneRequired("trufflehog", "trufflehog-verified", "gitleaks")
}
