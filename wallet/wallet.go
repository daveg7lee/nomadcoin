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
	address    string
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
		w.privateKey = restoreKey()
	} else {
		key := createPrivateKey()
		persistKey(key)
		w.privateKey = key
	}
	w.address = createAddress(w.privateKey)
}

func restoreKey() *ecdsa.PrivateKey {
	keyAsbytes, err := os.ReadFile(walletName)
	utils.HandleErr(err)

	key, err := x509.ParseECPrivateKey(keyAsbytes)
	utils.HandleErr(err)

	return key
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

func createAddress(key *ecdsa.PrivateKey) string {

}
