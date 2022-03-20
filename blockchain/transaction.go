package blockchain

import (
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
	TxId  string `json:"txId"`
	Index int    `json:"index"`
	Owner string `json:"owner"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type UTxOut struct {
	TxId   string `json:"txId`
	Index  int    `json:"index"`
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

func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("dave")
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil
	return txs
}

func (t *Tx) calculateId() {
	t.Id = utils.Hash(t)
}

func makeTx(from, to string, amount int) (*Tx, error) {

}

func checkHaveEnoughMoney(from string, amount int) bool {
	return Blockchain().BalanceByAddress(from) < amount
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{TxId: "", Index: -1, Owner: "COINBASE"},
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
