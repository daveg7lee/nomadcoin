package block

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/daveg7lee/nomadcoin/db"
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

func (b *Block) toBytes() []byte {
	var blockBuffer bytes.Buffer
	encoder := gob.NewEncoder(&blockBuffer)
	utils.HandleErr(encoder.Encode(b))
	return blockBuffer.Bytes()
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, b.toBytes())
}

func CreateBlock(data, lastHash string, height int) *Block {
	newBlock := &Block{
		Data: data, Hash: "",
		PrevHash: lastHash,
		Height:   height,
	}
	newBlock.calculateHash()
	newBlock.persist()
	return newBlock
}
