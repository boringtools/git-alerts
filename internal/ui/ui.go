package ui

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func PrintBanner(s string) {
	fmt.Fprint(os.Stderr, s, "\n")
}

func PrintSuccess(s string, args ...any) {
	msg := fmt.Sprintf(s, args...)
	fmt.Fprint(os.Stderr, text.FgGreen.Sprint(msg), "\n")
}

func PrintMsg(s string, args ...any) {
	msg := fmt.Sprintf(s, args...)
	fmt.Fprint(os.Stderr, text.Bold.Sprint(msg), "\n")
}

func PrintWarning(s string, args ...any) {
	msg := fmt.Sprintf(s, args...)
	fmt.Fprint(os.Stdout, text.FgYellow.Sprint(msg), "\n")
}

func PrintError(s string, args ...any) {
	msg := fmt.Sprintf(s, args...)
	fmt.Fprint(os.Stderr, text.FgRed.Sprint(msg), "\n")
}

func PrintTable(data []string, header string) {
	tbl := table.NewWriter()
	tbl.SetOutputMirror(os.Stdout)
	tbl.AppendHeader(table.Row{header})

	for _, value := range data {
		tbl.AppendRows([]table.Row{
			{value},
		})
	}
	tbl.AppendSeparator()
	tbl.Render()
}
