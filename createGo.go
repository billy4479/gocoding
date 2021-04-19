package main

import (
	"fmt"
	"os/exec"
)

func CreateGo(c *Config) error {
	command := exec.Command("go", "mod", "init", fmt.Sprintf("%s/%s/%s", appConfig.gitCloudProvider, appConfig.gitUsername, c.projectName))
	return command.Run()
	//TODO: .gitignore
}
