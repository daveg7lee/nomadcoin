package blockchain

import (
	"sync"

	"github.com/daveg7lee/nomadcoin/block"
)

type blockchain struct {
	blocks []*block.Block
}

var b *blockchain
var once sync.Once

func (b *blockchain) addBlock(newBlock *block.Block) {
	b.blocks = append(b.blocks, newBlock)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].GetHash()
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(initBlockchain)
	}
	return b
}

func initBlockchain() {
	b = &blockchain{}
	genesisBlock := block.CreateBlock("Genesis Block", getLastHash())
	b.addBlock(genesisBlock)
}
