package common

import (
	"github.com/boringtools/git-alerts/pkg/models"
)

func GetGitHubAPIEndPoints(username string) *models.GitHubAPIEndPoints {
	return &models.GitHubAPIEndPoints{
		GetUsers: GitHubAPIBaseURL + "/orgs/" + GitHubOrg + "/members",
		GetUsersRepo: GitHubAPIBaseURL + "/users/" + username + "/repos",
	}
}

func GetReportFileNames() *models.ReportFileNames {
	return &models.ReportFileNames{
		GitHubOrgUsers:          GitHubOrg + "_users.json",
		GitHubOrgPublicRepos:    GitHubOrg + "_public_repos.json",
		GitHubOrgPublicReposNew: GitHubOrg + "_public_repos_new.json",
	}
}

func GetReportFilePaths() *models.ReportFileNames {
	return &models.ReportFileNames{
		GitHubOrgUsers:          ReportPath + GetReportFileNames().GitHubOrgUsers,
		GitHubOrgPublicRepos:    ReportPath + GetReportFileNames().GitHubOrgPublicRepos,
		GitHubOrgPublicReposNew: ReportPath + GetReportFileNames().GitHubOrgPublicReposNew,
	}
}
