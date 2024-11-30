package monitor

import (
	"path/filepath"
	"strconv"

	"github.com/boringtools/git-alerts/internal/ui"
	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/models"
	"github.com/boringtools/git-alerts/pkg/notification"
	"github.com/boringtools/git-alerts/pkg/secrets"
	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	newRepoTableHeader table.Row
	newRepoTableRow    []table.Row
	NewRepoData        []models.MonitorRepositoryData
)

func GetMonitorData() {
	for key, value := range common.NewMonitorRepositories {
		cloneDirectory := filepath.Join(common.CloneDirectoryPath, strconv.Itoa(key))
		secrets.CloneRepo(value, cloneDirectory)
		isSecretFound, _ := secrets.RunGitleaks(cloneDirectory, false)
		NewRepoData = append(NewRepoData, models.MonitorRepositoryData{Repository: value, Secrets: isSecretFound})
		common.RemoveDirectory(cloneDirectory)
	}
}

func GenerateSlackMessage() (string, error) {
	common.SlackMessage = common.SlackMessage + ":eyes: *New Public Repositories Detected* \n"

	for _, value := range NewRepoData {
		if common.GitleaksScan {
			msg := ":point_right: " + value.Repository + " *Secrets Detected* :eyes: *" + strconv.FormatBool(value.Secrets) + "*\n"
			common.SlackMessage = common.SlackMessage + msg
		} else {
			msg := ":point_right: " + value.Repository + "\n"
			common.SlackMessage = common.SlackMessage + msg
		}
	}

	return common.SlackMessage, nil
}

func OutputNewRepository() {
	newRepoTableHeader = table.Row{"Repository", "secret found"}

	GetMonitorData()
	for _, value := range NewRepoData {

		if common.GitleaksScan {
			tableRow := table.Row{value.Repository, value.Secrets}
			newRepoTableRow = append(newRepoTableRow, tableRow)
		} else {
			tableRow := table.Row{value.Repository, "Not Scanned"}
			newRepoTableRow = append(newRepoTableRow, tableRow)
		}
	}

	ui.PrintTable(newRepoTableHeader, newRepoTableRow)

	if common.SlackNotification {
		msg, _ := GenerateSlackMessage()
		errSendSlackNotification := notification.SendSlackNotification(msg)

		if errSendSlackNotification != nil {
			ui.PrintError("error in sending slack notificaiton : %s", errSendSlackNotification)
		} else {
			ui.PrintSuccess("Slack notification sent successfully")
		}
	}
}
