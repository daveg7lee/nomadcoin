package blockchain

import (
	"sync"

	"github.com/daveg7lee/nomadcoin/block"
)

type blockchain struct {
	NewestHash string `json:"newest hash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func Blockchain() *blockchain {
	if b == nil {
		once.Do(initBlockchain)
	}
	return b
}

func initBlockchain() {
	b = &blockchain{NewestHash: "", Height: 0}
	newBlock := block.CreateBlock("Genesis", Blockchain().NewestHash, Blockchain().Height+1)
	b.AddBlock(newBlock)
}

func (b *blockchain) AddBlock(block *block.Block) {
	b.NewestHash = block.Hash
	b.Height = block.Height
}
