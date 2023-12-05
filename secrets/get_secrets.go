package secrets

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/boringtools/git-alerts/common"
	"github.com/boringtools/git-alerts/config"
	"github.com/boringtools/git-alerts/logger"
)

type Repos struct {
	Name     string `json:"full_name"`
	CloneUrl string `json:"clone_url"`
	Fork     bool   `json:"fork"`
}

var (
	repo []Repos
)

func GetSecrets() {
	logger.Log("Running secrets scan")
	fileContent := common.GetJsonFileContent(config.GhFileNames()[1])
	json.Unmarshal(fileContent, &repo)

	directoryName := "cloned_repo"
	directoryPath := os.Getenv("rfp") + directoryName

	_, errDirExists := os.Stat(directoryPath)

	if errDirExists != nil {
		CreateDirectory(directoryPath)
	} else {
		RemoveDirectory(directoryPath)
		CreateDirectory(directoryPath)
	}

	for key, value := range repo {
		if !value.Fork {
			cloneDirectory := directoryPath + "/" + strconv.Itoa(key)
			tfTarget := "file://" + cloneDirectory

			CloneRepo(value.CloneUrl, cloneDirectory)
			RunTruffleHog(tfTarget)
			RemoveDirectory(cloneDirectory)
		}
	}
}
