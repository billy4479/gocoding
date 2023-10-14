package main

import "fmt"

const helpMSG = `
Usage:
	gocoding [OPTIONS] <PROJECT_TYPE> [OUTPUT_DIR]

Project types:
	* cpp
	* go
	* svelte
	* sveltekit
	* ts
	* rust-qt

Options:
	See 'gocoding -help' to have a complete list of flags`

func help() {
	fmt.Println(helpMSG)
}
