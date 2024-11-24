package github

import (
	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/config"
	"github.com/boringtools/git-alerts/pkg/ui"
	"github.com/boringtools/git-alerts/pkg/utils"
)

func Connecter() {
	utils.CheckAuthenticatedScan()
	ui.PrintSuccess("Scan started : %s", utils.GetCurrentTime())

	gitHubAPILimit, errLimit := GetGitHubPATLimits()

	if errLimit != nil {
		ui.PrintError("error fetching GitHub API limits : %s", errLimit)
	}
	ui.PrintMsg("Remaining PAT limit : %v", gitHubAPILimit.Remaining)

	if gitHubAPILimit.Remaining <= 1 {
		ui.PrintError("GitHub PAT limit exceeded")
	}

	users, errGetGitHubUsers := GetGitHubUsers()
	if errGetGitHubUsers != nil {
		ui.PrintError("error fetching GitHub users : %s", errGetGitHubUsers)
	}

	errSaveToJson := utils.SaveToJson(users, config.GetReportFilePaths().GitHubOrgUsers)

	if errSaveToJson != nil {
		ui.PrintError("error saving data to JSON : %s", errSaveToJson)
	} else {
		ui.PrintSuccess("%s users fetched successfully %s", common.GitHubOrg, config.GetReportFilePaths().GitHubOrgUsers)
	}

	usersRepo, errGetGitHubUsers := GetGitHubUsersRepos()

	if errGetGitHubUsers != nil {
		ui.PrintError("%s", errGetGitHubUsers)
	}

	if common.Command == "monitor" {
		errSaveToJson := utils.SaveToJson(usersRepo, config.GetReportFilePaths().GitHubOrgPublicReposNew)

		if errSaveToJson != nil {
			ui.PrintError("error saving data to JSON : %s", errSaveToJson)
		} else {
			ui.PrintSuccess("%s public repositories fetched successfully %s", common.GitHubOrg, config.GetReportFilePaths().GitHubOrgPublicRepos)
		}
	} else {
		errSaveToJson := utils.SaveToJson(usersRepo, config.GetReportFilePaths().GitHubOrgPublicRepos)

		if errSaveToJson != nil {
			ui.PrintError("error saving data to JSON : %s", errSaveToJson)
		} else {
			ui.PrintSuccess("%s public repositories fetched successfully %s", common.GitHubOrg, config.GetReportFilePaths().GitHubOrgPublicRepos)
		}
	}

	utils.PrintSummery()
}
