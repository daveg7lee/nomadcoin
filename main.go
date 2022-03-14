package main

import (
	"os"

	"github.com/fatih/color"
)

func main() {
	boldText := color.New(color.FgHiWhite, color.Bold)

	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "explorer":
		boldText.Println("Start HTML Explorer")
	case "rest":
		boldText.Println("Start REST API")
	default:
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
