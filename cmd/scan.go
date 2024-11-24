package cmd

import (
	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/github"
	"github.com/boringtools/git-alerts/pkg/ui"
	"github.com/boringtools/git-alerts/pkg/utils"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan public repositories",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		common.Command = cmd.Use
		github.Connecter()

		ui.PrintSuccess("Scan Ended : %s", utils.GetCurrentTime())
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
