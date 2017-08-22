package leveldb

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

const DB_FILE_PATH = "demo.db"

func newDb() *leveldb.DB{
	db, err := leveldb.OpenFile(DB_FILE_PATH, nil)
	if err != nil {
		log.Fatal("leveldb\t-", "OpenFile", err.Error())
	}
	return db
}

func Demo1() {
	db := newDb()

	if err := db.Put([]byte("id"), []byte("36"), nil); err != nil {
		log.Println("leveldb\t-", "put key", err.Error())
	}

	value, err := db.Get([]byte("id"), nil)
	if err != nil {
		log.Println("leveldb\t-", "get key", err.Error())
	}
	log.Println("leveldb\t-", "key-velue", string(value))
}
