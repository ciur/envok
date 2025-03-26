package commands

func ReloadCurrentProfile() {
	name := getCurrentProfile()
	ExportProfile(name)
}
