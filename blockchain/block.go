package blockchain

import (
	"errors"
	"strings"
	"time"

	"github.com/daveg7lee/nomadcoin/db"
	"github.com/daveg7lee/nomadcoin/utils"
	"github.com/daveg7lee/nomadcoin/wallet"
)

type Block struct {
	Height       int    `json:"height"`
	Transactions []*Tx  `json:"transactions"`
	Hash         string `json:"hash"`
	PrevHash     string `json:"previous hash,omitempty"`
	Difficulty   int    `json:"difficulty"`
	Nonce        int    `json:"nonce"`
	Timestamp    int    `json:"timestamp"`
}

var ErrorNotFound = errors.New("block not found")

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
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

func CreateBlock(lastHash string, height, difficulty int) *Block {
	newBlock := &Block{
		Transactions: []*Tx{makeCoinbaseTx(wallet.Wallet().Address)},
		Hash:         "",
		PrevHash:     lastHash,
		Height:       height,
		Difficulty:   difficulty,
		Nonce:        0,
	}
	newBlock.mine()
	newBlock.Transactions = Mempool.TxToConfirm()
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
