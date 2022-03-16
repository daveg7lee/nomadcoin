package block

import (
	"errors"
	"fmt"

	"github.com/daveg7lee/nomadcoin/db"
	"github.com/daveg7lee/nomadcoin/utils"
)

type Block struct {
	Height     int    `json:"height"`
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"previous hash,omitempty"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

var ErrorNotFound = errors.New("block not found")

func (b *Block) calculateHash() {
	b.Hash = utils.Hash([]byte(b.Data + b.PrevHash + fmt.Sprint(b.Height)))
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
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

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrorNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}
