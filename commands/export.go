package commands

import (
	"fmt"
	"os"
	"sort"

	"github.com/ciur/envok/profiles"
)

func ExportProfile(defaultConfigPath *string, name string) {
	var exportedKeys []string
	var keysToUnset []string

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
				exportedKeys = append(exportedKeys, k)
			}
		}
	}

	for _, profile := range items {
		for k := range profile.Vars {
			if !contains(exportedKeys, k) {
				keysToUnset = append(keysToUnset, k)
			}
		}
	}

	keysToUnset = uniqueStrings(keysToUnset)
	for _, item := range keysToUnset {
		fmt.Printf("unset %s\n", item)
	}
}

func uniqueStrings(input []string) []string {
	seen := make(map[string]struct{}) // use an empty struct because it uses no memory
	var result []string

	for _, str := range input {
		if _, exists := seen[str]; !exists {
			seen[str] = struct{}{}
			result = append(result, str)
		}
	}

	return result
}

func contains(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}
