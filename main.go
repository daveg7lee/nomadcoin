package main

import (
	"github.com/daveg7lee/nomadcoin/explorer"
	"github.com/daveg7lee/nomadcoin/rest"
)

func main() {
	go rest.Start(5000)
	explorer.Start(4000)
}
