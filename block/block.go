package block

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/daveg7lee/nomadcoin/db"
	"github.com/daveg7lee/nomadcoin/utils"
)

const difficulty int = 2

type Block struct {
	Height     int    `json:"height"`
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"previous hash,omitempty"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

var ErrorNotFound = errors.New("block not found")

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		blockAsBytes := []byte(fmt.Sprint(b))
		hash := fmt.Sprintf("%x", sha256.Sum256(blockAsBytes))
		fmt.Printf("Block as String:%s\nHash:%s\nTarget:%s\nNonce:%d\n\n\n", blockAsBytes, hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
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
		PrevHash:   lastHash,
		Height:     height,
		Difficulty: difficulty,
		Nonce:      0,
	}
	newBlock.mine()
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
