package common

import (
	"os"
	"os/exec"
	"time"
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
