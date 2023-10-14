package main

func CreateRustQt(c *Config) error {
	err := fromGit("https://github.com/billy4479/rust-cpp-qt6-template")
	if err != nil {
		return err
	}

	//TODO: Change project name in CMakeLists.txt and Cargo.toml

	return nil
}
