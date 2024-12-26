package common

var (
	Command                string
	GitHubOrg              string
	ReportPath             string
	SlackNotification      bool
	TrufflehogScan         bool
	TrufflehogVerifiedScan bool
	GitleaksScan           bool
	UsersFilePath          string
)

var (
	AuthenticatedScan          bool
	NumberOfGitHubUsers        int
	NumberOfPublicRepositories int
	CloneDirectoryPath         string
	SlackMessage               string
	NewMonitorRepositories     []string
)
