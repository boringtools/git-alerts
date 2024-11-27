package common

var (
	Command                string
	GitHubOrg              string
	ReportPath             string
	SlackNotification      bool
	TrufflehogScan         bool
	TrufflehogVerifiedScan bool
	GitleaksScan           bool
)

var (
	AuthenticatedScan          bool
	NumberOfGitHubUsers        int
	NumberOfPublicRepositories int
	CloneDirectoryPath         string
)
