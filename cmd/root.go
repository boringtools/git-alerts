package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	org    string
	report string
	csv    bool
	slack  bool
	trufflehog bool
	trufflehogVerified bool
	gitleaks bool
)

var rootCmd = &cobra.Command{
	Use:   "git-alerts",
	Short: "A Public Git repository & misconfiguration detection tool",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&org, "org", "o", "", "GitHub organization name")
	rootCmd.MarkPersistentFlagRequired("org")
	rootCmd.PersistentFlags().BoolVarP(&csv, "csv", "c", false, "CSV report format")
	rootCmd.PersistentFlags().StringVarP(&report, "report-path", "r", "/tmp/", "Report file path")

}
