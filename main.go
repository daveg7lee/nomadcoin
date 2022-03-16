package main

import (
	"github.com/daveg7lee/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.Blockchain()

	chain.AddBlock("firstBlock")
	chain.AddBlock("secondBlock")
	chain.AddBlock("thridBlock")
}
