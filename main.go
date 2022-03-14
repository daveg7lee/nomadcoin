package main

import (
	"os"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
}

func usage() {
	color.Yellow("Welcome to Nomad Coin\n\n")
	color.Red("Please use the following commands:\n\n")
	color.Cyan("explorer:     Start the HTML Explorer")
	color.Cyan("rest:         Start the REST API (recommended)")
	os.Exit(0)
}
