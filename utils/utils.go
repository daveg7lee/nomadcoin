package utils

import (
	"log"
	"time"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetYear() int {
	now := time.Now()
	return now.Year()
}
