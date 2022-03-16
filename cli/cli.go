package cli

import (
	"flag"
	"os"
	"runtime"

	"github.com/daveg7lee/nomadcoin/explorer"
	"github.com/daveg7lee/nomadcoin/rest"
	"github.com/fatih/color"
)

func Start() {
	usage()
	executeServer(parseFlags())
}

func parseFlags() (string, int) {
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	flag.Parse()

	return *mode, *port
}

func executeServer(mode string, port int) {
	switch mode {
	case "rest":
		rest.Start(port)
	case "html":
		explorer.Start(port)
	case "both":
		go rest.Start(port)
		explorer.Start(port + 1)
	}
}

func usage() {
	if len(os.Args) < 2 {
		color.Yellow("Welcome to Nomad Coin\n\n")
		color.Red("Please use the following flags:\n\n")
		color.Cyan("-port:     Set the port of the server")
		color.Cyan("-mode:     Choose between 'html', 'rest', and 'both'")
		runtime.Goexit()
	}
}
