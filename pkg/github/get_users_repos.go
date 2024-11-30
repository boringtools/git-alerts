package github

import (
	"encoding/json"

	"github.com/boringtools/git-alerts/internal/ui"
	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/models"
)

type RepoURL struct {
	URL string `json:"repos_url"`
}

var (
	rURL     []RepoURL
	repos    []models.GitHubRepository
	allRepos []models.GitHubRepository
)

func GetGitHubUsersRepos() ([]byte, error) {
	ui.PrintMsg("Fetching " + common.GitHubOrg + " users public repositories...")

	users, _ := common.GetJSONFileContent(common.GetReportFilePaths().GitHubOrgUsers)
	json.Unmarshal(users, &rURL)

	parameters := map[string]string{
		"per_page": "100",
	}
	for _, value := range rURL {
		usersRepo, _, _ := GetGitHubResponse(value.URL, common.AuthenticatedScan, parameters)

		json.Unmarshal(usersRepo, &repos)
		allRepos = append(allRepos, repos...)
	}

	jsonData, err := json.Marshal(allRepos)

	if err != nil {
		return nil, err
	}
	common.NumberOfPublicRepositories = len(allRepos)
	return jsonData, nil
}
