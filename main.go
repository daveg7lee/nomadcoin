package main

import (
	"fmt"

	"github.com/daveg7lee/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Thrid Block")
	chain.AddBlock("Fourth Block")
	chain.AddBlock("Fifth Block")
	for _, block := range chain.GetAllBlocks() {
		fmt.Println(block.GetData())
		fmt.Println(block.GetHash())
		fmt.Println(block.GetPrevHash())
	}
}
