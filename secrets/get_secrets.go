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
	fileContent := common.GetJsonFileContent(config.GhFileNames()[1])
	json.Unmarshal(fileContent, &repo)

	if os.Getenv("gitleaks") == "true" {

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
				CloneRepo(value.CloneUrl, cloneDirectory)
				logger.LogP("Scanning repository : ", value.CloneUrl)
				RunGitleaks(cloneDirectory)
				RemoveDirectory(cloneDirectory)
			}
		}
	} else {
		for _, value := range repo {
			if !value.Fork {
				logger.LogP("Scanning repository : ", value.CloneUrl)
				RunTruffleHog(value.CloneUrl)
			}
		}
	}
}
