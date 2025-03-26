package commands

import (
	"fmt"
	"os"

	"github.com/ciur/envok/profiles"
	"github.com/fatih/color"
)

func ShowCurrentProfile() {

	currentProfileName := getCurrentProfile()
	configPath, err := getConfigPath()
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
			for k, v := range profile.Vars {
				envVarValue := os.Getenv(k)
				if envVarValue != v {
					color.Red("%s=<value differ: expected=%q, actual=%q>\n", k, v, envVarValue)
				} else {
					color.Green("%s=%s\n", k, v)
				}
			}
		}
	}
}
