package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func Hash(i interface{}) string {
	data := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func ToBytes(data interface{}) []byte {
	var dataBuffer bytes.Buffer
	encoder := gob.NewEncoder(&dataBuffer)
	HandleErr(encoder.Encode(data))
	return dataBuffer.Bytes()
}

func FromBytes(i interface{}, data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(decoder.Decode(i))
}

func Splitter(s string, sep string, i int) string {
	r := strings.Split(s, sep)
	if len(r)-1 < i {
		return ""
	}
	return r[i]
}
