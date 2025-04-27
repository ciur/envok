package commands

import (
	"fmt"
	"os"

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
			for k, v := range profile.Vars {
				fmt.Printf("export %s=%s\n", k, v)
			}
		}
	}
}
