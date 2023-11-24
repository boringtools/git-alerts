package common

import (
	"os"
	"time"

	"github.com/boringtools/git-alerts/logger"
)

func GetTime() (date string) {
	time := time.Now()
	date = time.Format("Mon Jan 2 15:04:05 MST 2006")
	return date
}

func SetEnvs(data map[string]string) {
	for key, value := range data {
		os.Setenv(key, value)
	}
}

func StartChecks() {

	_, isGitHubPat := os.LookupEnv("GITHUB_PAT")
	_, isSlackHook := os.LookupEnv("SLACK_HOOK")

	if !isGitHubPat {
		logger.LogERR("GITHUB_PAT is not configured in ENV variable")
		os.Exit(1)
	}

	if os.Getenv("slack") == "true" {
		if !isSlackHook {
			logger.LogERR("SLACK_HOOK is not configured in ENV variable")
			os.Exit(1)
		}
	}

}

func Start() {
	currentTime := GetTime()
	logger.LogP("Scan started : ", currentTime)

}
