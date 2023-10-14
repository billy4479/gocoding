package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	MODE_TYPESCRIPT = iota
	MODE_SVELTE
	MODE_SVELTEKIT
	MODE_GO
	MODE_CPP
	MODE_RUST_QT
)

type Config struct {
	mode        int
	disableGit  *bool
	workingDir  string
	projectName string
	err         error
}

func ParseFlags() Config {
	c := Config{disableGit: flag.Bool("no-git", false, "Does not create a new git repository")}
	printModes := flag.Bool("modes", false, "Display possibles modes")
	flag.Parse()

	if *printModes {
		help()
		c.err = fmt.Errorf("Printing help")
		return c
	}

	args := flag.Args()

	osDir, err := os.Getwd()
	if err != nil {
		c.err = err
		return c
	}

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "No enough arguments. See 'gocoding -help'")
		c.err = fmt.Errorf("No enough arguments. See 'gocoding -help'")
		return c
	}

	switch args[0] {
	case "ts":
		c.mode = MODE_TYPESCRIPT
	case "svelte":
		c.mode = MODE_SVELTE
	case "sveltekit":
		c.mode = MODE_SVELTEKIT
	case "go":
		c.mode = MODE_GO
	case "cpp":
		c.mode = MODE_CPP
	case "rust-qt":
		c.mode = MODE_RUST_QT
	default:
		fmt.Fprintf(os.Stderr, "Invalid mode: %s.\n", args[1])
		help()
		return c
	}

	if len(args) == 1 {
		c.workingDir = osDir
		c.projectName = filepath.Base(osDir)
		return c
	} else if len(args) == 2 {
		err := os.MkdirAll(args[1], 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid folder path: %s.\n", args[1])
			c.err = err
			return c
		}

		c.workingDir, err = filepath.Abs(args[1])
		if err != nil {
			c.err = err
			return c
		}
		c.projectName = filepath.Base(c.workingDir)

		err = os.Chdir(c.workingDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid folder path: %s.\n", args[1])
			c.err = err
			return c
		}
	} else {
		fmt.Fprintln(os.Stderr, "Too many arguments. See 'gocoding -help'")
		c.err = fmt.Errorf("Too many arguments. See 'gocoding -help'")
		return c
	}

	return c
}
