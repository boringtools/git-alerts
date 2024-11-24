package github

import (
	"encoding/json"
	"os"

	"github.com/boringtools/git-alerts/pkg/config"
	"github.com/boringtools/git-alerts/pkg/models"
	"github.com/boringtools/git-alerts/pkg/utils"
)

var (
	newRepos    []models.GitHubRepository
	oldRepos    []models.GitHubRepository
	allNewRepos []string
	allOldRepos []string
)

func GetNewPublicRepositories() ([]string, error) {
	newFile, _ := utils.GetJSONFileContent(config.GetReportFilePaths().GitHubOrgPublicReposNew)
	oldFile, _ := utils.GetJSONFileContent(config.GetReportFilePaths().GitHubOrgPublicRepos)

	errNewFile := json.Unmarshal(newFile, &newRepos)

	if errNewFile != nil {
		return nil, errNewFile
	}

	errOldFile := json.Unmarshal(oldFile, &oldRepos)

	if errOldFile != nil {
		return nil, errOldFile
	}

	for _, value := range newRepos {
		allNewRepos = append(allNewRepos, value.HtmlURL)
	}

	for _, value := range oldRepos {
		allOldRepos = append(allOldRepos, value.HtmlURL)
	}

	lisOfNewPublicRepos := utils.SliceDiff(allNewRepos, allOldRepos)

	os.Rename(config.GetReportFilePaths().GitHubOrgPublicReposNew, config.GetReportFilePaths().GitHubOrgPublicRepos)
	return lisOfNewPublicRepos, nil
}
