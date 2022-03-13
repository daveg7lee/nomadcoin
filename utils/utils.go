package utils

import (
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
