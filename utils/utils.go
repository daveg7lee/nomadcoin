package utils

import (
	"crypto/sha256"
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
