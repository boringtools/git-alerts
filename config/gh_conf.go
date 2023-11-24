package config

import "os"

func GhUrls() (urls []string) {
	api := "https://api.github.com"
	rateLimit := api + "/rate_limit"
	users := api + "/orgs/" + os.Getenv("org") + "/members"

	urls = append(urls, api, rateLimit, users)
	return urls
}

func GhFileNames() (fileNames []string) {
	users := os.Getenv("org") + "_users_" + os.Getenv("command")
	usersRepo := os.Getenv("org") + "_users_repos_" + os.Getenv("command")

	fileNames = append(fileNames, users, usersRepo)
	return fileNames
}

func GhFilePaths() (filePaths []string) {
	userRepoMonitor := os.Getenv("rfp") + GhFileNames()[1] + ".json"
	userRepoScan := os.Getenv("rfp") + os.Getenv("org") + "_users_repos_scan.json"

	filePaths = append(filePaths, userRepoMonitor, userRepoScan)
	return filePaths
}
