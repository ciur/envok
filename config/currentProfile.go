package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
)

// GetConfigFilePath returns the path where the configuration file should be stored
func GetConfigFilePath() (string, error) {
	// Get the current user info
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	// Define the config directory and file path
	configDir := filepath.Join(usr.HomeDir, ".config", "envok")
	configFilePath := filepath.Join(configDir, "current_profile")

	// Create the directory if it doesn't exist
	err = os.MkdirAll(configDir, 0700) // Make sure only the user has access
	if err != nil {
		return "", err
	}

	// Return the config file path
	return configFilePath, nil
}

// SaveConfigFile saves a simple file to the user's home directory
func SaveConfigFile(content string) error {
	// Get the path to the config file
	configFilePath, err := GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("failed to get config file path: %w", err)
	}

	// Create and open the file for writing
	file, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to config file: %w", err)
	}

	// Set the file permissions to be accessible only by the current user
	err = os.Chmod(configFilePath, 0600)
	if err != nil {
		return fmt.Errorf("failed to set file permissions: %w", err)
	}

	// Ensure that the file owner is the current user
	err = os.Lchown(configFilePath, syscall.Getuid(), syscall.Getgid())
	if err != nil {
		return fmt.Errorf("failed to set file owner: %w", err)
	}

	return nil
}
