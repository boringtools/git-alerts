package main

import (
	"fmt"

	"github.com/boringtools/git-alerts/cmd"
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
	fmt.Println(banner)
	cmd.Execute()
}
