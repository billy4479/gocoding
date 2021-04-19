package main

type AppConfig struct {
	license          string
	email            string
	gitUsername      string
	gitCloudProvider string
}

func (c *AppConfig) Validate() bool {
	result := true

	if c.email == "" {
		result = false
	}
	if c.gitCloudProvider == "" {
		result = false
	}
	if c.gitUsername == "" {
		result = false
	}
	if c.license == "" {
		result = false
	}

	return result
}

// Edit me!
var appConfig = AppConfig{
	license:          "MIT",                     // TODO: Use https://github.com/nishanths/license
	email:            "giachi.ellero@gmail.com", // TODO: Read from git config or from flags
	gitUsername:      "billy4479",
	gitCloudProvider: "github.com",
}
