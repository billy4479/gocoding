package main

func CreateCpp(c *Config) error {
	err := fromGit("https://github.com/billy4479/cpp_template")
	if err != nil {
		return err
	}

	//TODO: Change project name in CMakeLists.txt

	return nil
}
