package github

import (
	"encoding/json"
	"strconv"

	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/config"
	"github.com/boringtools/git-alerts/pkg/models"
	"github.com/boringtools/git-alerts/pkg/ui"
)

var (
	user     []models.GitHubUser
	allUsers []models.GitHubUser
)

func GetGitHubUsers() ([]byte, error) {
	ui.PrintMsg("Fetching %s users", common.GitHubOrg)

	url := config.GetGitHubAPIEndPoints().GetUsers
	parameters := map[string]string{
		"per_page": "100",
	}

	ghResponse, pageLength, _ := GetGitHubResponse(url, common.AuthenticatedScan, parameters)

	if pageLength == 0 {

		json.Unmarshal(ghResponse, &user)
		allUsers = append(allUsers, user...)
		jsonData, err := json.Marshal(allUsers)

		if err != nil {
			return nil, err
		}
		common.NumberOfGitHubUsers = len(allUsers)
		return jsonData, nil
	} else {
		for i := 1; i <= pageLength; i++ {
			parameters["page"] = strconv.Itoa(i)
			ghResponse, _, _ := GetGitHubResponse(url, common.AuthenticatedScan, parameters)

			json.Unmarshal(ghResponse, &user)
			allUsers = append(allUsers, user...)
		}

		jsonData, err := json.Marshal(allUsers)

		if err != nil {
			return nil, err
		}
		common.NumberOfGitHubUsers = len(allUsers)
		return jsonData, nil
	}
}
