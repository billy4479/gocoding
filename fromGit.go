package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func fromGit(repo string) error {
	s := strings.Split(repo, "/")
	repoName := s[len(s)-1]
	gitPath, err := exec.LookPath("git")
	if err != nil {
		return err
	}

	clone := exec.Command(gitPath, "clone", repo, "--bare")
	clone.Stdout = os.Stdout
	clone.Stderr = os.Stderr
	checkout := exec.Command(gitPath, fmt.Sprintf("--git-dir=./%s.git", repoName), "--work-tree=.", "checkout")
	checkout.Stdout = os.Stdout
	checkout.Stderr = os.Stderr

	err = clone.Run()
	if err != nil {
		return err
	}

	err = checkout.Run()
	if err != nil {
		return err
	}

	return os.RemoveAll(fmt.Sprintf("./%s.git", repoName))
}
