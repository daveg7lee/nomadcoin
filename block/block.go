package block

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Height   int    `json:"height"`
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"previous hash,omitempty"`
}

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func CreateBlock(data, lastHash string, height int) *Block {
	newBlock := Block{Data: data, Hash: "", PrevHash: lastHash, Height: height}
	newBlock.calculateHash()
	return &newBlock
}
