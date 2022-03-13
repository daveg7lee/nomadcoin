package block

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"previous hash,omitempty"`
}

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func CreateBlock(data, lastHash string) *Block {
	newBlock := Block{Data: data, Hash: "", PrevHash: lastHash}
	newBlock.calculateHash()
	return &newBlock
}
