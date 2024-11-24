package utils

import (
	"os"
	"time"

	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/config"
	"github.com/jedib0t/go-pretty/v6/table"
)

func GetCurrentTime() (formattedTime string) {
	currentTime := time.Now()
	formattedTime = currentTime.Format("Mon Jan 2 15:04:05 MST 2006")
	return formattedTime
}

func CheckAuthenticatedScan() {
	_, isGitHubPATExists := os.LookupEnv("GITHUB_PAT")

	if isGitHubPATExists {
		common.AuthenticatedScan = true
	} else {
		common.AuthenticatedScan = false
	}
}

func IsPreviousScanFileExists() bool {
	_, errScanFile := os.Stat(config.GetReportFilePaths().GitHubOrgPublicRepos)

	if errScanFile != nil {
		return false
	}

	return true
}

func PrintSummery() {
	tbl := table.NewWriter()
	tbl.SetOutputMirror(os.Stdout)
	tbl.AppendHeader(table.Row{"Scan Summery", "Data"})
	tbl.AppendRows([]table.Row{
		{"Total Users", common.NumberOfGitHubUsers},
		{"Total Users Repositories", common.NumberOfPublicRepositories},
	})
	tbl.AppendSeparator()
	tbl.Render()
}
