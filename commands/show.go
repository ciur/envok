package commands

import (
	"fmt"
	"os"
	"sort"

	"github.com/ciur/envok/profiles"
	"github.com/fatih/color"
)

func ShowCurrentProfile(defaultConfigPath *string, name *string) {

	currentProfileName := getCurrentProfile()

	if name != nil {
		currentProfileName = *name
	}

	configPath, err := getConfigPath(defaultConfigPath)
	if err != nil {
		fmt.Printf("Error getting config: %s", err)
		os.Exit(1)
	}

	items, err := profiles.Load(configPath)
	if err != nil {
		fmt.Printf("Error loading profiles: %s\n", err)
		os.Exit(1)
	}

	if currentProfileName == "" {
		fmt.Println("No current profile set")
		return
	}

	fmt.Printf("---%s---\n", currentProfileName)

	for _, profile := range items {
		if profile.Name == currentProfileName {

			keys := make([]string, 0, len(profile.Vars))
			for k := range profile.Vars {
				keys = append(keys, k)
			}

			sort.Strings(keys)

			// iterate profile.Vars with keys sorted
			for _, k := range keys {
				envVarValue := os.Getenv(k)
				if envVarValue != profile.Vars[k] {
					color.Red("%s=<value differ: expected=%q, actual=%q>\n", k, profile.Vars[k], envVarValue)
				} else {
					color.Green("%s=%s\n", k, profile.Vars[k])
				}
			}
		}
	}
}
