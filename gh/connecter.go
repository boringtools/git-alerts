package gh

import (
	"github.com/boringtools/git-alerts/common"
	"github.com/boringtools/git-alerts/config"
)

func Connecter() {
	common.StartChecks()
	common.Start()

	CheckPatLimit()
	users := GetUsers()
	common.SaveToJson(users, config.GhFileNames()[0])

	usersRepo := GetUsersRepos()
	common.SaveToJson(usersRepo, config.GhFileNames()[1])

	GetData()
}
