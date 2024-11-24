package notification

import (
	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/github"
	"github.com/boringtools/git-alerts/pkg/ui"
)

var (
	slackMessage string
)

func Notify() {
	newRepos, _ := github.GetNewPublicRepositories()

	if len(newRepos) == 0 {
		ui.PrintMsg("No new public repositories detected")
	} else {
		ui.PrintTable(newRepos, "New Public Repositories Detected")

		if common.SlackNotification {
			slackMessage = slackMessage + ":eyes: *New Public Repositories Detected* \n"
			for _, value := range newRepos {
				repo := ":point_right: " + value + "\n"
				slackMessage = slackMessage + repo
			}

			errNotificaiton := SendSlackNotification(slackMessage)

			if errNotificaiton != nil {
				ui.PrintError("error in sending slack notificaiton : %s", errNotificaiton)
			} else {
				ui.PrintSuccess("slack notificaiton sent successfully!")
			}

		}
	}
}
