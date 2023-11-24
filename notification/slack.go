package notification

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/boringtools/git-alerts/logger"
)

func SlackNotification(content string) {

	jsonData := []byte("{'text': '" + content + "'}")
	json.Marshal(jsonData)

	slackHook := os.Getenv("SLACK_HOOK")
	request, error := http.NewRequest("POST", slackHook, bytes.NewReader(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if error != nil {
		logger.LogERR("Error in making HTTP request to slack")
	}

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	sBody := string(body)

	if sBody == "ok" {
		logger.Log("Slack notification sent successfully")
	} else {
		logger.LogERR("Error in sending slack notification")
	}
}
