package secrets

import (
	"encoding/json"
	"path/filepath"
	"strconv"

	"github.com/boringtools/git-alerts/internal/ui"
	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/models"
	"github.com/go-git/go-git/v5/plumbing/transport"
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
				errClone := CloneRepo(value.CloneURL, cloneDirectory)

				if errClone != nil {
					if errClone == transport.ErrEmptyRemoteRepository {
						ui.PrintWarning("Skipping empty repository : %s", value.CloneURL)
					} else {
						ui.PrintError("Error while cloning repository : %s", errClone)
					}
				} else {
					ui.PrintMsg("Scanning repository : %s", value.CloneURL)
					RunGitleaks(cloneDirectory, true)
					common.RemoveDirectory(cloneDirectory)
				}

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
