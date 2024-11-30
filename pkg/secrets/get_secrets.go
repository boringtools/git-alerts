package secrets

import (
	"encoding/json"
	"path/filepath"
	"strconv"

	"github.com/boringtools/git-alerts/internal/ui"
	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/models"
)

var (
	repo []models.GitHubRepository
)

func RunSecretsScan() {
	fileContent, _ := common.GetJSONFileContent(common.GetReportFilePaths().GitHubOrgPublicRepos)
	json.Unmarshal(fileContent, &repo)

	if common.GitleaksScan {

		for key, value := range repo {
			if !value.Fork {
				cloneDirectory := filepath.Join(common.CloneDirectoryPath, strconv.Itoa(key))
				CloneRepo(value.CloneURL, cloneDirectory)
				ui.PrintMsg("Scanning repository : %s", value.CloneURL)
				RunGitleaks(cloneDirectory, true)
				common.RemoveDirectory(cloneDirectory)
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
