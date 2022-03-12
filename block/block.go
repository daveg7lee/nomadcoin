package block

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	data     string
	hash     string
	prevHash string
}

func (b *Block) HashBlock() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

func (b *Block) setData() {

}
