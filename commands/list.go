package commands

import (
	"fmt"
	"os"

	"github.com/ciur/envok/profiles"
)

func ListProfiles() {
	items, err := profiles.Load(".envok.yaml")
	if err != nil {
		fmt.Printf("Error loading profiles: %s\n", err)
		os.Exit(1)
	}
	for _, profile := range items {
		fmt.Printf("%s\n", profile.Name)
	}
}
