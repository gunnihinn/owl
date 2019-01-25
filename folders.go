package main

func initFolders() error {
	if err := os.MkdirAll(dirConfig(), 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(dirData(), 0755); err != nil {
		return err
	}

	return nil
}

func dirConfig() string {
	if dir, ok := os.LookupEnv("XDG_CONFIG_HOME"); ok {
		return dir
	}

	return path.Join(os.Getenv("HOME"), ".config")
}

func fileConfig() string {
	return path.Join(dirConfig(), "owl.cfg")
}

func dirData() string {
	if dir, ok := os.LookupEnv("XDG_DATA_HOME"); ok {
		return path.Join(dir, "owl")
	}

	return path.Join(os.Getenv("HOME"), ".local", "share", "owl")
}

func fileState() string {
	return path.Join(dirData(), "owl.state")
}

func fileExists(name string) bool {
	_, err := os.Stat(name)

	return err == nil
}
