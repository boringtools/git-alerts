package gh

import (
	"encoding/json"
	"os"

	"github.com/boringtools/git-alerts/common"

	"github.com/boringtools/git-alerts/config"
	"github.com/jedib0t/go-pretty/v6/table"
)

var Summery map[string]interface{}

var (
	usersIn     interface{}
	usersRepoIn interface{}
)

func GetData() {
	usersFile := common.GetJsonFileContent(config.GhFileNames()[0])
	json.Unmarshal(usersFile, &usersIn)
	usersLength, _ := usersIn.([]interface{})

	usersRepoFile := common.GetJsonFileContent(config.GhFileNames()[1])
	json.Unmarshal(usersRepoFile, &usersRepoIn)
	usersReposLength, _ := usersRepoIn.([]interface{})

	tbl := table.NewWriter()
	tbl.SetOutputMirror(os.Stdout)
	tbl.AppendHeader(table.Row{"Scan Summery", "Data"})
	tbl.AppendRows([]table.Row{
		{"Total Users", len(usersLength)},
		{"Total Users Repositories", len(usersReposLength)},
	})

	tbl.AppendSeparator()
	tbl.Render()

}

func init() {
	Summery = map[string]interface{}{}
}
