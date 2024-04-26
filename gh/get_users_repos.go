package gh

import (
	"encoding/json"
	"os"

	"github.com/boringtools/git-alerts/common"
	"github.com/boringtools/git-alerts/config"
	"github.com/boringtools/git-alerts/logger"
)

type Repo struct {
	FullName    string `json:"full_name"`
	Private     bool   `json:"private"`
	HtmlURL     string `json:"html_url"`
	Description string `json:"description"`
	Fork        bool   `json:"fork"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
	Push        string `json:"pushed_at"`
	CloneURL    string `json:"clone_url"`
	Visiability string `json:"visibility"`
}

type RepoURL struct {
	URL string `json:"repos_url"`
}

var (
	rURL     []RepoURL
	repos    []Repo
	allRepos []Repo
)

func GetUsersRepos() (jsonData []byte) {
	logger.Log("Fetching " + os.Getenv("org") + " users public repositories")

	users := common.GetJsonFileContent(config.GhFileNames()[0])

	json.Unmarshal(users, &rURL)

	parameters := map[string]string{
		"per_page": "100",
	}
	for _, value := range rURL {

		usersRepo, _ := GetResponse(value.URL, common.Auth, parameters)

		json.Unmarshal(usersRepo, &repos)
		allRepos = append(allRepos, repos...)

	}

	jsonData, err := json.Marshal(allRepos)

	if err != nil {
		logger.LogERR("GetUsers - Error in marshalling json data")
	}
	return jsonData
}
