package main

import (
	"github.com/daveg7lee/nomadcoin/explorer"
	"github.com/daveg7lee/nomadcoin/rest"
)

func main() {
	go rest.Start(4000)
	explorer.Start(5000)
}
