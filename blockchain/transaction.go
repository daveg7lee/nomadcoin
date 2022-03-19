package blockchain

import (
	"errors"
	"time"

	"github.com/daveg7lee/nomadcoin/utils"
)

type mempool struct {
	Txs []*Tx
}

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

type TxIn struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

const (
	minerReward int = 10
)

var Mempool *mempool = &mempool{}

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("dave", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

func (t *Tx) calculateId() {
	t.Id = utils.Hash(t)
}

func makeTx(from, to string, amount int) (*Tx, error) {
	if checkHaveEnoughMoney(from, amount) {
		return nil, errors.New("not enough money")
	}
}

func checkHaveEnoughMoney(from string, amount int) bool {
	if Blockchain().BalanceByAddress(from) < amount {
		return true
	}
	return false
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{Owner: "COINBASE", Amount: minerReward},
	}
	txOuts := []*TxOut{
		{Owner: address, Amount: minerReward},
	}
	tx := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.calculateId()
	return &tx
}
