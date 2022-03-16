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
	checkpoint := db.Checkpoint()
	if checkpoint == nil {
		b.AddBlock("Genesis Block")
	} else {
		b.restore(checkpoint)
	}
}

func (b *blockchain) AddBlock(data string) {
	newBlock := block.CreateBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = newBlock.Hash
	b.Height = newBlock.Height
	b.persist()
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) Blocks() []*block.Block {
	var blocks []*block.Block
	hashCursor := b.NewestHash
	for {
		block, _ := block.FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}
