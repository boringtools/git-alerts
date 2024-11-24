package secrets

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/config"
	"github.com/boringtools/git-alerts/pkg/models"
	"github.com/boringtools/git-alerts/pkg/ui"
	"github.com/boringtools/git-alerts/pkg/utils"
)

var (
	repo []models.GitHubRepository
)

func RunSecretsScan() {
	fileContent, _ := utils.GetJSONFileContent(config.GetReportFilePaths().GitHubOrgPublicRepos)
	json.Unmarshal(fileContent, &repo)

	if common.GitleaksScan {

		directoryName := "cloned_repo"
		directoryPath := common.ReportPath + directoryName

		_, errDirExists := os.Stat(directoryPath)

		if os.IsNotExist(errDirExists) {
			CreateDirectory(directoryPath)
		} else {
			RemoveDirectory(directoryPath)
			CreateDirectory(directoryPath)
		}

		for key, value := range repo {
			if !value.Fork {
				cloneDirectory := directoryPath + "/" + strconv.Itoa(key)
				CloneRepo(value.CloneURL, cloneDirectory)
				ui.PrintMsg("Scanning repository : %s", value.CloneURL)
				RunGitleaks(cloneDirectory)
				RemoveDirectory(cloneDirectory)
			} else {
				ui.PrintWarning("Skipping forked repository : %s", value.CloneURL)
			}
		}
	} else {
		for _, value := range repo {
			if !value.Fork {
				ui.PrintMsg("Scanning repository : %s", value.CloneURL)
				RunTruffleHog(value.CloneURL)
			} else {
				ui.PrintWarning("Skipping forked repository : %s", value.CloneURL)
			}
		}
	}
}
