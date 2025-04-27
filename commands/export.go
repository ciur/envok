package commands

import (
	"fmt"
	"os"
	"sort"

	"github.com/ciur/envok/profiles"
)

func ExportProfile(defaultConfigPath *string, name string) {
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

	for _, profile := range items {
		if profile.Name == name {

			keys := make([]string, 0, len(profile.Vars))
			for k := range profile.Vars {
				keys = append(keys, k)
			}

			sort.Strings(keys)

			// iterate profile.Vars with keys sorted
			for _, k := range keys {
				fmt.Printf("export %s=%s\n", k, profile.Vars[k])
			}
		}
	}
}
