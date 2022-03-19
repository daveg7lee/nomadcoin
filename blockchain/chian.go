package blockchain

import (
	"sync"

	"github.com/daveg7lee/nomadcoin/db"
	"github.com/daveg7lee/nomadcoin/utils"
)

type blockchain struct {
	NewestHash        string `json:"newest hash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"current difficulty"`
}

const (
	defaultDifficulty  int = 4
	difficultyInterval int = 5
	blockInterval      int = 3
	allowedRange       int = 1
)

var b *blockchain
var once sync.Once

func Blockchain() *blockchain {
	if b == nil {
		once.Do(initBlockchain)
	}
	return b
}

func initBlockchain() {
	b = &blockchain{Height: 0}
	checkpoint := db.Checkpoint()
	if checkpoint == nil {
		b.AddBlock()
	} else {
		b.restore(checkpoint)
	}
}

func (b *blockchain) AddBlock() {
	newBlock := CreateBlock(b.NewestHash, b.Height+1)
	b.NewestHash = newBlock.Hash
	b.Height = newBlock.Height
	b.CurrentDifficulty = newBlock.Difficulty
	b.persist()
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) Blocks() []Block {
	var blocks []Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, *block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return b.calculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

func (b *blockchain) calculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastCalculatedBlock := allBlocks[difficultyInterval-1]
	actualTime := (newestBlock.Timestamp / 60) - (lastCalculatedBlock.Timestamp / 60)
	expectedTime := difficultyInterval * blockInterval

	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime >= (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}
