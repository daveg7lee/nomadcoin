package blockchain

import (
	"sync"

	"github.com/daveg7lee/nomadcoin/block"
	"github.com/daveg7lee/nomadcoin/db"
	"github.com/daveg7lee/nomadcoin/utils"
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
	b.AddBlock("Genesis Block")
}

func (b *blockchain) AddBlock(data string) {
	newBlock := block.CreateBlock(data, b.NewestHash, b.Height)
	b.NewestHash = newBlock.Hash
	b.Height = newBlock.Height
	b.persist()
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}
