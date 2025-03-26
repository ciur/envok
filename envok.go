package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ciur/envok/commands"
)

const VERSION = "0.2"

var profileName = flag.String("p", "", "Profile name")
var version = flag.Bool("v", false, "show version and exit")

func main() {
	flag.Parse()

	flag.Usage = func() {
		w := flag.CommandLine.Output()

		fmt.Fprintf(
			os.Stderr,
			"Usage: %s [-p profile-name] export|list|show|reload\n",
			os.Args[0],
		)

		flag.PrintDefaults()

		fmt.Fprintf(w, "For more details check: https://github.com/ciur/envok\n")
	}

	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}

	switch args[0] {
	case "list":
		commands.ListProfiles()
	case "export":
		commands.ExportProfile(*profileName)
	case "show":
		commands.ShowCurrentProfile()
	case "reload":
		commands.ReloadCurrentProfile()
	default:
		flag.Usage()
	}
}
