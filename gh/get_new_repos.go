package gh

import (
	"encoding/json"
	"os"

	"github.com/boringtools/git-alerts/common"
	"github.com/boringtools/git-alerts/config"
)

type Repos struct {
	URL string `json:"html_url"`
}

var (
	newRepos    []Repos
	oldRepos    []Repos
	allNewRepos []string
	allOldRepos []string
)

func GetNewRepos() (diff []string) {
	newFile := common.GetJsonFileContent(config.GhFileNames()[1])
	oldFile := common.GetJsonFileContent(os.Getenv("org") + "_users_repos_scan")

	json.Unmarshal(newFile, &newRepos)
	json.Unmarshal(oldFile, &oldRepos)

	for _, value := range newRepos {
		allNewRepos = append(allNewRepos, value.URL)
	}

	for _, value := range oldRepos {
		allOldRepos = append(allOldRepos, value.URL)
	}

	diff = common.SliceDiff(allNewRepos, allOldRepos)

	if len(diff) == 0 {
		return nil
	}

	os.Rename(config.GhFilePaths()[0], config.GhFilePaths()[1])
	return diff
}
