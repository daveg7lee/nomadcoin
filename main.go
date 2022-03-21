package main

import (
	"github.com/daveg7lee/nomadcoin/cli"
	"github.com/daveg7lee/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
