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
	dataKey      = "blockchain"
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

func SaveBlock(hash string, data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(dataKey), data)
		return err
	})
	utils.HandleErr(err)
}

func Checkpoint() []byte {
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(dataKey))
		return nil
	})
	return data
}

func Close() {
	DB().Close()
}

func Block(hash string) []byte {
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}
