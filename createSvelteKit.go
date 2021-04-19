package main

func CreateSvelteKit(c *Config) error {
	err := fromGit("https://github.com/billy4479/sveltekit-template")
	if err != nil {
		return err
	}

	//TODO: Edit package.json

	return nil
}
