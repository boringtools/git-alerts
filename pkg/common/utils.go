package common

import (
	"os"
	"os/exec"
	"time"

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
		AuthenticatedScan = true
	} else {
		AuthenticatedScan = false
	}
}

func IsPreviousScanFileExists() bool {
	_, errScanFile := os.Stat(GetReportFilePaths().GitHubOrgPublicRepos)
	return errScanFile == nil || !os.IsNotExist(errScanFile)
}

func CreateDirectory(dirPath string) error {
	cmd := exec.Command("mkdir", dirPath)
	_, err := cmd.Output()

	if err != nil {
		return err
	}
	return nil
}

func RemoveDirectory(dirPath string) error {
	cmd := exec.Command("rm", "-rf", dirPath)
	_, err := cmd.Output()

	if err != nil {
		return err
	}

	return nil
}

func PrintSummery() {
	tbl := table.NewWriter()
	tbl.SetOutputMirror(os.Stdout)
	tbl.AppendHeader(table.Row{"Scan Summery", "Data"})
	tbl.AppendRows([]table.Row{
		{"Total Users", NumberOfGitHubUsers},
		{"Total Users Repositories", NumberOfPublicRepositories},
	})
	tbl.AppendSeparator()
	tbl.Render()
}
