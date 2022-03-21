package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"

	"github.com/daveg7lee/nomadcoin/utils"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

const (
	walletName = "nomadcoin.wallet"
)

var w *wallet

func Wallet() *wallet {
	if w == nil {
		initWallet()
	}
	return w
}

func hasWalletFile() bool {
	_, err := os.Stat(walletName)
	return !os.IsNotExist(err)
}

func initWallet() {
	w = &wallet{}
	if hasWalletFile() {

	} else {
		key := createPrivateKey()
		persistKey(key)
		w.privateKey = key
	}
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)

	err = os.WriteFile(walletName, bytes, 0644)
	utils.HandleErr(err)
}

func createPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privateKey
}
