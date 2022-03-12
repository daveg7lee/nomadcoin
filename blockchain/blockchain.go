package blockchain

import "github.com/daveg7lee/nomadcoin/block"

type blockchain struct {
	blocks []block.Block
}

var b *blockchain

func GetBlockchain() *blockchain {
	initBlockchain()
	return b
}

func initBlockchain() {
	if b == nil {
		b = &blockchain{}
	}
}
