package commands

func ReloadCurrentProfile(defaultConfigPath *string, name *string) {
	profileName := getCurrentProfile()

	if name != nil {
		profileName = *name
	}

	ExportProfile(defaultConfigPath, profileName)
}
