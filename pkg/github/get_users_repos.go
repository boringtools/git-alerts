package github

import (
	"encoding/json"

	"github.com/boringtools/git-alerts/internal/ui"
	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/models"
)

var (
	users    []models.GitHubUser
	repos    []models.GitHubRepository
	allRepos []models.GitHubRepository
)

func GetGitHubUsersRepos() ([]byte, error) {
	ui.PrintMsg("Fetching %s users public repositories...", common.GitHubOrg)

	parameters := map[string]string{
		"per_page": "100",
	}

	if common.UsersFilePath == "" {
		data, _ := common.GetJSONFileContent(common.GetReportFilePaths().GitHubOrgUsers)
		json.Unmarshal(data, &users)

		for _, value := range users {
			usersRepo, _, _ := GetGitHubResponse(value.ReposUrl, common.AuthenticatedScan, parameters)

			json.Unmarshal(usersRepo, &repos)
			allRepos = append(allRepos, repos...)
		}

	} else {

		users, errUsers := common.GetCSVFileContent(common.UsersFilePath)

		if errUsers != nil {
			return nil, errUsers
		}

		for _, username := range users {
			repos_url := common.GetGitHubAPIEndPoints(username).GetUsersRepo
			usersRepo, _, _ := GetGitHubResponse(repos_url, common.AuthenticatedScan, parameters)
			json.Unmarshal(usersRepo, &repos)
			allRepos = append(allRepos, repos...)
		}
	}

	jsonData, err := json.Marshal(allRepos)

	if err != nil {
		return nil, err
	}
	common.NumberOfPublicRepositories = len(allRepos)
	return jsonData, nil
}
