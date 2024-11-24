package config

import (
	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/models"
)

func GetGitHubAPIEndPoints() *models.GitHubAPIEndPoints {
	return &models.GitHubAPIEndPoints{
		GetUsers: common.GitHubAPIBaseURL + "/orgs/" + common.GitHubOrg + "/members",
	}
}

func GetReportFileNames() *models.ReportFileNames {
	return &models.ReportFileNames{
		GitHubOrgUsers:          common.GitHubOrg + "_github_users.json",
		GitHubOrgPublicRepos:    common.GitHubOrg + "_public_repositories.json",
		GitHubOrgPublicReposNew: common.GitHubOrg + "_public_repositories_new.json",
	}
}

func GetReportFilePaths() *models.ReportFileNames {
	return &models.ReportFileNames{
		GitHubOrgUsers:          common.ReportPath + GetReportFileNames().GitHubOrgUsers,
		GitHubOrgPublicRepos:    common.ReportPath + GetReportFileNames().GitHubOrgPublicRepos,
		GitHubOrgPublicReposNew: common.ReportPath + GetReportFileNames().GitHubOrgPublicReposNew,
	}
}
