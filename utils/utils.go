package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func StrToInt(data string) int {
	intData, err := strconv.Atoi(data)
	HandleErr(err)
	return intData
}

func Hash(data []byte) string {
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

func ToBytes(data interface{}) []byte {
	var dataBuffer bytes.Buffer
	encoder := gob.NewEncoder(&dataBuffer)
	HandleErr(encoder.Encode(data))
	return dataBuffer.Bytes()
}
