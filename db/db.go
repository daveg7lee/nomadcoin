package db

import (
	"github.com/boltdb/bolt"
	"github.com/daveg7lee/nomadcoin/utils"
)

var db *bolt.DB

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

func DB() *bolt.DB {
	if db == nil {
		initDB()
		initBuckets()
	}
	return db
}

func initDB() {
	dbPointer, err := bolt.Open(dbName, 0600, nil)
	db = dbPointer
	utils.HandleErr(err)
}

func initBuckets() {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
		utils.HandleErr(err)
		_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
		return err
	})
	utils.HandleErr(err)
}
