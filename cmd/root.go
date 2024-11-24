package cmd

import (
	"os"

	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-alerts",
	Short: "A Public Git repository & misconfiguration detection tool",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&common.GitHubOrg, "org", "o", "", "GitHub organization name")
	rootCmd.MarkPersistentFlagRequired("org")
	rootCmd.PersistentFlags().StringVarP(&common.ReportPath, "report-path", "r", "/tmp/", "Report file path")
}
