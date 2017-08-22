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

func SimplePutAndGet() {
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

// 批量写入到db
func SimpleBatch() {
	db := newDb()

	batch := new(leveldb.Batch)
	batch.Put([]byte("k11"), []byte("v11"))
	batch.Put([]byte("k21"), []byte("v21"))
	batch.Put([]byte("k31"), []byte("v31"))

	//batch.Delete([]byte("baz"))

	if err := db.Write(batch, nil); err != nil {
		log.Println("leveldb batch\t-", "write", err.Error())
	}

	log.Println("leveldb batch\t-", "len", batch.Len())

	v1, err:= db.Get([]byte("k1"), nil)
	if err != nil {
		log.Println("leveldb batch\t-", "write batch then get", err.Error())
	}
	log.Println("leveldb batch\t-", "write batch then get", string(v1))

	v11, err:= db.Get([]byte("k11"), nil)
	if err != nil {
		log.Println("leveldb batch\t-", "write batch then get", err.Error())
	}
	log.Println("leveldb batch\t-", "write batch then get", string(v11))
}

func 
func UseOptions() {

}
