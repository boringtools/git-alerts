package gh

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/boringtools/git-alerts/common"
	"github.com/boringtools/git-alerts/config"
	"github.com/boringtools/git-alerts/logger"
)

type Users struct {
	Username          string `json:"login"`
	Url               string `json:"url"`
	ProfileUrl        string `json:"html_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	Admin             bool   `json:"site_admin"`
}

var (
	user     []Users
	allUsers []Users
)

func GetUsers() (jsonData []byte) {
	logger.Log("Fetching " + os.Getenv("org") + " users")

	url := config.GhUrls()[2]
	parameters := map[string]string{
		"per_page": "100",
	}
	ghResponse, pageLength := GetResponse(url, common.Auth, parameters)

	if pageLength == 0 {

		json.Unmarshal(ghResponse, &user)
		allUsers = append(allUsers, user...)
		jsonData, err := json.Marshal(allUsers)

		if err != nil {
			logger.LogERR("GetUsers - Error in marshalling json data")
		}
		return jsonData
	} else {
		for i := 1; i <= pageLength; i++ {
			parameters["page"] = strconv.Itoa(i)
			ghResponse, _ := GetResponse(url, common.Auth, parameters)

			json.Unmarshal(ghResponse, &user)
			allUsers = append(allUsers, user...)
		}

		jsonData, err := json.Marshal(allUsers)

		if err != nil {
			logger.LogERR("GetUsers - Error in marshalling json data")
		}
		return jsonData
	}
}
