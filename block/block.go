package block

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	data     string
	hash     string
	prevHash string
}

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

func (b *Block) setPrevHash(prevHash string) {
	b.prevHash = prevHash
}

func (b *Block) GetHash() string {
	return b.hash
}

func (b *Block) GetData() string {
	return b.data
}

func (b *Block) GetPrevHash() string {
	return b.prevHash
}

func CreateBlock(data, lastHash string) *Block {
	newBlock := Block{data: data, hash: "", prevHash: ""}
	newBlock.setPrevHash(lastHash)
	newBlock.calculateHash()
	return &newBlock
}
