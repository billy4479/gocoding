package main

import (
	"fmt"
	"os"
)

func main() {
	c := ParseFlags()

	if c.err != nil {
		os.Exit(1)
	}

	var err error
	switch c.mode {
	case MODE_GO:
		err = CreateGo(&c)
	case MODE_CPP:
		err = CreateCpp(&c)
	case MODE_SVELTE:
		err = CreateSvelte(&c)
	case MODE_SVELTEKIT:
		err = CreateSvelteKit(&c)
	case MODE_TYPESCRIPT:
		err = CreateTS(&c)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
