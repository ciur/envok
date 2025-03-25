package main

import (
	"flag"

	"github.com/ciur/envok/commands"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return
	}

	if args[0] == "list" {
		commands.ListProfiles()
	}

}
