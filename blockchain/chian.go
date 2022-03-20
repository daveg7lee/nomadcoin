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

func (b *blockchain) AddBlock() {
	newBlock := CreateBlock(b.NewestHash, b.Height+1, getDifficulty(b))
	b.NewestHash = newBlock.Hash
	b.Height = newBlock.Height
	b.CurrentDifficulty = newBlock.Difficulty
	persistBlockchain(b)
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func Blockchain() *blockchain {
	once.Do(initBlockchain)
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

func persistBlockchain(b *blockchain) {
	db.SaveBlockchain(utils.ToBytes(b))
}

func Blocks(b *blockchain) []Block {
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

func getDifficulty(b *blockchain) int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return calculateDifficulty(b)
	} else {
		return b.CurrentDifficulty
	}
}

func calculateDifficulty(b *blockchain) int {
	allBlocks := Blocks(b)
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

func UTxOutsByAddress(b *blockchain, address string) []*UTxOut {
	var unspentTxOuts []*UTxOut
	spentTxOuts := STxOutsByAddress(b, address)

	for _, block := range Blocks(b) {
		for _, tx := range block.Transactions {
			for index, output := range tx.TxOuts {
				if output.Owner == address {
					_, ok := spentTxOuts[tx.Id]
					if !ok {
						unspentTxOut := &UTxOut{TxId: tx.Id, Index: index, Amount: output.Amount}
						if !isOnMempool(unspentTxOut) {
							unspentTxOuts = append(unspentTxOuts, unspentTxOut)
						}
					}
				}
			}
		}
	}

	return unspentTxOuts
}

func STxOutsByAddress(b *blockchain, address string) map[string]bool {
	spentTxOuts := make(map[string]bool)

	for _, block := range Blocks(b) {
		for _, tx := range block.Transactions {
			for _, input := range tx.TxIns {
				if input.Owner == address {
					spentTxOuts[input.TxId] = true
				}
			}
		}
	}

	return spentTxOuts
}

func BalanceByAddress(b *blockchain, address string) int {
	var amount int
	txOuts := UTxOutsByAddress(b, address)

	for _, txOut := range txOuts {
		amount += txOut.Amount
	}

	return amount
}
