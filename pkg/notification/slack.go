package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/boringtools/git-alerts/pkg/models"
)

var (
	slackMessage string
)

func SendSlackNotification(content string) error {

	payload := models.SlackPayload{Text: content}
	byteData, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	request, errorRequest := http.NewRequest("POST", os.Getenv("SLACK_HOOK"), bytes.NewBuffer(byteData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if errorRequest != nil {
		return errorRequest
	}

	client := &http.Client{}
	response, errorResponse := client.Do(request)

	if errorResponse != nil {
		return errorResponse
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	stringBody := string(body)

	if stringBody == "ok" {
		return nil
	} else {
		return fmt.Errorf("error in sending slack notifcation")
	}
}
