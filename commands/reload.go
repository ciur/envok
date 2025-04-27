package commands

func ReloadCurrentProfile(defaultConfigPath *string, name *string) {
	profileName := getCurrentProfile()

	if name != nil && *name != "" {
		profileName = *name
	}

	ExportProfile(defaultConfigPath, profileName)
}
