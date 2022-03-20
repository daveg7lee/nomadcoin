package wallet

import (
	"crypto/ecdsa"
	"os"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

func Wallet() *wallet {
	if w == nil {

	}
	return w
}

func hasWalletFile() bool {
	_, err := os.Stat("nomadcoin.wallet")
	return !os.IsNotExist(err)
}

func initWallet() {
	if hasWalletFile() {

	}
}
