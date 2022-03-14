package block

import (
	"fmt"

	"github.com/daveg7lee/nomadcoin/utils"
)

type Block struct {
	Height   int    `json:"height"`
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"previous hash,omitempty"`
}

func (b *Block) calculateHash() {
	b.Hash = utils.Hash([]byte(b.Data + b.PrevHash + fmt.Sprint(b.Height)))
}

func CreateBlock(data, lastHash string, height int) *Block {
	newBlock := &Block{
		Data: data, Hash: "",
		PrevHash: lastHash,
		Height:   height,
	}
	newBlock.calculateHash()
	return newBlock
}
