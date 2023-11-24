package reporter

import (
	"os"

	"github.com/boringtools/git-alerts/gh"
	"github.com/boringtools/git-alerts/logger"
	"github.com/boringtools/git-alerts/notification"
	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	slackMessage string
)

func Notify() {
	newRepos := gh.GetNewRepos()

	if len(newRepos) == 0 {
		logger.Log("No new public repositories detected")
	} else {

		tbl := table.NewWriter()
		tbl.SetOutputMirror(os.Stdout)
		tbl.AppendHeader(table.Row{"New Public Repositories Detected"})

		for _, value := range newRepos {
			tbl.AppendRows([]table.Row{
				{value},
			})
		}
		tbl.AppendSeparator()
		tbl.Render()

		if os.Getenv("slack") == "true" {
			slackMessage = slackMessage + ":eyes: *New Public Repositories Detected* \n"
			for _, value := range newRepos {
				repo := ":point_right: " + value + "\n"
				slackMessage = slackMessage + repo
			}

			notification.SlackNotification(slackMessage)
		}
	}
}
