package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	CONFIG_1 = ".envok.yaml"
	CONFIG_2 = ".envok.yml"
)

func getCurrentProfile() string {
	const currentProfileNameFilePath = ".config/envok/prof_name.txt"

	home, err := os.UserHomeDir()
	filePath := fmt.Sprintf("%s/%s", home, currentProfileNameFilePath)
	data, err := os.ReadFile(filePath)
	var currentProfileName string

	if err != nil {
		fmt.Printf("%s\n", err)
		currentProfileName = ""
	} else {
		currentProfileName = strings.TrimSpace(string(data))
	}

	return currentProfileName
}

func getConfigPath(config *string) (string, error) {

	if config != nil {
		if _, err := os.Stat(*config); err == nil {
			return *config, nil // File found
		}
	}

	fileName, err := searchFileUpwards(CONFIG_1)
	if err != nil {
		fileName, err = searchFileUpwards(CONFIG_2)
		if err != nil {
			return "", fmt.Errorf("failed to get config file %s (or %s)\n", CONFIG_1, CONFIG_2)
		}
	}

	return fileName, nil

}

func searchFileUpwards(filename string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	for {
		filePath := filepath.Join(currentDir, filename)

		if _, err := os.Stat(filePath); err == nil {
			return filePath, nil // File found
		}
		if currentDir == homeDir || currentDir == "/" {
			break
		}

		// Move up one directory
		currentDir = filepath.Dir(currentDir)
	}

	return "", fmt.Errorf("file %q not found", filename)
}
