package gh

import (
	"encoding/json"
	"os"
	"time"

	"github.com/boringtools/git-alerts/common"
	"github.com/boringtools/git-alerts/config"
	"github.com/boringtools/git-alerts/logger"
)

type Rate struct {
	Rate Limits `json:"rate"`
}
type Limits struct {
	Total     int `json:"limit"`
	Used      int `json:"used"`
	Remaining int `json:"remaining"`
	Reset     int `json:"reset"`
}

func CheckPatLimit() {
	var limit Rate
	url := config.GhUrls()[1]
	parameters := map[string]string{}

	ghResponse, _ := GetResponse(url, common.Auth, parameters)
	json.Unmarshal(ghResponse, &limit)

	remaining := limit.Rate.Remaining
	unixTime := limit.Rate.Reset
	time := time.Unix(int64(unixTime), 0)
	date := time.Format("Mon Jan 2 15:04:05 MST 2006")

	if remaining <= 10 {
		logger.LogERR("GitHub request limit exceeded")
		logger.LogERR("Please try again once the limit reset or try with GitHub PAT")
		logger.LogERRP("Remaining request limit : ", remaining)
		logger.LogERRP("Reset time : ", date)
		os.Exit(1)
	}

	logger.LogP("Remaining request limit : ", remaining)

}
