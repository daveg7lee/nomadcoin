package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/daveg7lee/nomadcoin/utils"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	Address    string
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
	w.Address = createAddress(w.privateKey)
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
	x := key.X.Bytes()
	y := key.Y.Bytes()
	z := append(x, y...)
	return fmt.Sprintf("%x", z)
}

func Sign(payload string, w *wallet) string {
	payloadAsBytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadAsBytes)
	utils.HandleErr(err)

	signature := append(r.Bytes(), s.Bytes()...)

	return fmt.Sprintf("%x", signature)
}

func Verify(signature, payload, publicKey string) bool {

}
