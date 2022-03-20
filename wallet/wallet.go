package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/daveg7lee/nomadcoin/utils"
)

const (
	hashedMessage string = "253b27734439cc9ccda1d6efe98ab811d181a848a28b2291f791a96d738bc2a6"
	privateKey    string = "30770201010420d2dd1247794c617dc567a208773d9f28b8bb1410c70c9ed1880e4e275129329da00a06082a8648ce3d030107a14403420004563ad758a3f04b28afdc3c350336516ec58e265d611213a5e28ee8e222fac34d368bf6f9001e2eaed5a2260aa9b06c7a998a22707e260da38747bc65fbe4a441"
	signature     string = "34504d2f918ad6abced891abdf05b424c2c7ffd30722733458abdfaca101529bd493005d0eea5b256a49c54e1f36fb4d56cccee526dbae497345dd7dea76c831"
)

func Start() {
	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	_, err = x509.ParseECPrivateKey(privBytes)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)
}
