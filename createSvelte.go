package main

func CreateSvelte(c *Config) error {
	err := fromGit("https://github.com/billy4479/svelte-template")
	if err != nil {
		return err
	}

	//TODO: Edit package.json

	return nil
}
