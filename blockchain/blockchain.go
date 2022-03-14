package blockchain

import (
	"errors"
	"sync"

	"github.com/daveg7lee/nomadcoin/block"
)

type blockchain struct {
	blocks []*block.Block
}

var b *blockchain
var once sync.Once
var ErrNotFound = errors.New("Block not found")

func (b *blockchain) AddBlock(data string) {
	newBlock := block.CreateBlock(data, b.getLastHash(), len(b.blocks)+1)
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) getLastHash() string {
	totalBlocks := len(b.blocks)
	if totalBlocks == 0 {
		return ""
	}
	return b.blocks[totalBlocks-1].Hash
}

func (b *blockchain) GetAllBlocks() []*block.Block {
	return b.blocks
}

func (b *blockchain) GetBlock(height int) (*block.Block, error) {
	if height > len(b.blocks) {
		return nil, ErrNotFound
	}
	return b.blocks[height-1], nil
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(initBlockchain)
	}
	return b
}

func initBlockchain() {
	b = &blockchain{}
	b.AddBlock("Genesis Block")
}
