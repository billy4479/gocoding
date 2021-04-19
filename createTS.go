package main

func CreateTS(c *Config) error {
	err := fromGit("https://github.com/billy4479/js-base")
	if err != nil {
		return err
	}

	//TODO: Edit package.json

	return nil
}
