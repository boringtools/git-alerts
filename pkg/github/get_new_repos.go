package github

import (
	"encoding/json"
	"os"

	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/models"
)

var (
	newRepos    []models.GitHubRepository
	oldRepos    []models.GitHubRepository
	allNewRepos []string
	allOldRepos []string
)

func GetNewPublicRepositories() ([]string, error) {
	newFile, _ := common.GetJSONFileContent(common.GetReportFilePaths().GitHubOrgPublicReposNew)
	oldFile, _ := common.GetJSONFileContent(common.GetReportFilePaths().GitHubOrgPublicRepos)

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

	lisOfNewPublicRepos := common.SliceDiff(allNewRepos, allOldRepos)

	os.Rename(common.GetReportFilePaths().GitHubOrgPublicReposNew, common.GetReportFilePaths().GitHubOrgPublicRepos)
	return lisOfNewPublicRepos, nil
}
