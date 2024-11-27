package github

import (
	"github.com/boringtools/git-alerts/internal/ui"
	"github.com/boringtools/git-alerts/pkg/common"
)

func Connecter() {
	common.ValidateFlags()
	common.CheckAuthenticatedScan()
	ui.PrintSuccess("Scan started : %s", common.GetCurrentTime())

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

	errSaveToJson := common.SaveToJson(users, common.GetReportFilePaths().GitHubOrgUsers)

	if errSaveToJson != nil {
		ui.PrintError("error saving data to JSON : %s", errSaveToJson)
	} else {
		ui.PrintSuccess("%s users fetched successfully %s", common.GitHubOrg, common.GetReportFilePaths().GitHubOrgUsers)
	}

	usersRepo, errGetGitHubUsers := GetGitHubUsersRepos()

	if errGetGitHubUsers != nil {
		ui.PrintError("%s", errGetGitHubUsers)
	}

	if common.Command == "monitor" {
		errSaveToJson := common.SaveToJson(usersRepo, common.GetReportFilePaths().GitHubOrgPublicReposNew)

		if errSaveToJson != nil {
			ui.PrintError("error saving data to JSON : %s", errSaveToJson)
		} else {
			ui.PrintSuccess("%s public repositories fetched successfully %s", common.GitHubOrg, common.GetReportFilePaths().GitHubOrgPublicRepos)
		}
	} else {
		errSaveToJson := common.SaveToJson(usersRepo, common.GetReportFilePaths().GitHubOrgPublicRepos)

		if errSaveToJson != nil {
			ui.PrintError("error saving data to JSON : %s", errSaveToJson)
		} else {
			ui.PrintSuccess("%s public repositories fetched successfully %s", common.GitHubOrg, common.GetReportFilePaths().GitHubOrgPublicRepos)
		}
	}

	common.PrintSummery()
}
