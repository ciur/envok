package commands

import (
	"fmt"
	"os"
	"sort"

	"github.com/ciur/envok/profiles"
)

func ListProfiles(config *string) {

	currentProfileName := getCurrentProfile()
	configPath, err := getConfigPath(config)
	if err != nil {
		fmt.Printf("Error getting config: %s", err)
		os.Exit(1)
	}

	items, err := profiles.Load(configPath)

	if err != nil {
		fmt.Printf("Error loading profiles: %s\n", err)
		os.Exit(1)
	}

	sort.Sort(profiles.ByName(items))

	for _, profile := range items {
		if profile.Name == currentProfileName {
			fmt.Printf("%s*\n", profile.Name)
		} else {
			fmt.Printf("%s\n", profile.Name)
		}
	}
}
