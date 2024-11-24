package main

import (
	"github.com/boringtools/git-alerts/cmd"
	"github.com/boringtools/git-alerts/pkg/ui"
)

var banner string = `
 _____    _   _                _                 _         
/ ____|  (_) | |       /\     | |               | |        
| |  __   _  | |_     /  \    | |   ___   _ __  | |_   ___ 
| | |_ | | | | __|   / /\ \   | |  / _ \ | '__| | __| / __|
| |__| | | | | |_   / ____ \  | | |  __/ | |    | |_  \__ \
 \_____| |_|  \__| /_/    \_\ |_|  \___| |_|     \__| |___/														 
`

func main() {
	ui.PrintBanner(banner)
	cmd.Execute()
}
