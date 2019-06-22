package leveldb

import (
	"log"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	opt2 "github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

const DB_FILE_PATH = "demo.db"

func newDb() *leveldb.DB {
	db, err := leveldb.OpenFile(DB_FILE_PATH, nil)
	if err != nil {
		log.Fatal("leveldb\t-", "OpenFile", err.Error())
	}
	return db
}

func TestSimplePutAndGet(t *testing.T) {
	db := newDb()
	defer db.Close()

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
func TestSimpleBatch(t *testing.T) {
	db := newDb()
	defer db.Close()

	batch := new(leveldb.Batch)
	batch.Put([]byte("k11"), []byte("v11"))
	batch.Put([]byte("k21"), []byte("v21"))
	batch.Put([]byte("k31"), []byte("v31"))

	//batch.Delete([]byte("baz"))

	if err := db.Write(batch, nil); err != nil {
		log.Println("leveldb batch\t-", "write", err.Error())
	}

	log.Println("leveldb batch\t-", "len", batch.Len())

	v1, err := db.Get([]byte("k1"), nil)
	if err != nil {
		log.Println("leveldb batch\t-", "write batch then get", err.Error())
	}
	log.Println("leveldb batch\t-", "write batch then get", string(v1))

	v11, err := db.Get([]byte("k11"), nil)
	if err != nil {
		log.Println("leveldb batch\t-", "write batch then get", err.Error())
	}
	log.Println("leveldb batch\t-", "write batch then get", string(v11))
}

// 丢弃重载batch内容，丢弃之前batch的内容,看不出来有什么效果及用途
func TestSimpleBatchLoad(t *testing.T) {
	batch := new(leveldb.Batch)

	batch.Put([]byte("k12"), []byte("v12"))
	batch.Put([]byte("k13"), []byte("v13"))
	log.Println("leveldb\t-", "batch init", batch)

	data := []byte("hihi")
	if err := batch.Load(data); err != nil {
		log.Println("leveldb\t-", "batch load error", err.Error())
	} else {
		log.Println("leveldb\t-", "batch load success", batch)
	}
}

// 大概是将数据打包，后续不能再写入，而是只能读取
func SimpleConpactRange() {
	db := newDb()
	db.Close()

	r := util.Range{[]byte("k"), []byte("3")}
	db.CompactRange(r)
}

// 查询db相关属性
func TestSimpleGetProperty(t *testing.T) {
	db := newDb()
	defer db.Close()

	// "leveldb.num-files-at-level{1}"

	// "leveldb.stats"
	/*
		2017/08/23 10:21:10 leveldb	- get property success Compactions
		Level |   Tables   |    Size(MB)   |    Time(sec)  |    Read(MB)   |   Write(MB)
		-------+------------+---------------+---------------+---------------+---------------
			0   |          5 |       0.00063 |       0.00000 |       0.00000 |       0.00000
	*/

	// "leveldb.sstables"
	/*
		2017/08/23 10:22:33 leveldb	- get property success --- level 0 ---
		20:137["k1,v15" .. "k1,v13"]
		17:144["k11,v9" .. "k31,v11"]
		8:139["k1,v5" .. "k3,v7"]
		5:119["id,v3" .. "id,v3"]
		2:119["id,v1" .. "id,v1"]
	*/

	// "leveldb.blockpool"
	/*
		2017/08/23 10:24:11 leveldb	- get property success BufferPool{B·4101 Z·[0 0 0 0 0] Zm·[0 0 0 0 0] Zh·[0 0 0 0 0] G·0 P·0 H·0 <·0 =·0 >·0 M·0}
	*/

	// "leveldb.cachedblock"
	/*
		2017/08/23 10:25:27 leveldb	- get property success 0
	*/

	// "leveldb.openedtables"
	/*
		2017/08/23 10:26:12 leveldb	- get property success 0
	*/

	// "leveldb.alivesnaps"
	/*
		2017/08/23 10:26:56 leveldb	- get property success 0
	*/

	// "leveldb.aliveiters"
	/*
		2017/08/23 10:27:38 leveldb	- get property success 0
	*/

	name := "leveldb.aliveiters"

	if property, err := db.GetProperty(name); err != nil {
		log.Println("leveldb\t-", "get property error", err.Error())
	} else {
		log.Println("leveldb\t-", "get property success", property)
	}
}

// 获取快照,通过快照查询key
func TestSimpleGetSnapshot(t *testing.T) {
	db := newDb()
	defer db.Close()

	if snap, err := db.GetSnapshot(); err != nil {
		log.Println("leveldb\t-", "get snapshot error", err.Error())
	} else {
		log.Println("leveldb\t-", "get snapshot success", snap.String())

		if value, err := snap.Get([]byte("k1"), nil); err != nil {
			log.Println("leveldb\t-", "snapshot get key error", err.Error())
		} else {
			log.Println("leveldb\t", "snapshot get key success", string(value))
		}
	}
}

// 建立基于db的迭代器，遍历db所有数据
func TestSimpleNewDBIterator(t *testing.T) {
	db := newDb()
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	defer iter.Release()

	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		log.Println("leveldb\t-", "iterator based on db", string(key), string(value))
	}
}

// 查询某个key
func TestSimpleDBIteratorSeek(t *testing.T) {
	db := newDb()
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	defer iter.Release()

	if exist := iter.Seek([]byte("k1")); exist {
		log.Println("leveldb\t-", "iterator seek success", exist)
	}

	if exist := iter.Next(); exist {
		log.Println("leveldb\t-", "iterator next key and value", string(iter.Key()), string(iter.Value()))
	}

	if ok := iter.Last(); ok {
		log.Println("leveldb\t-", "iterator last key and value", string(iter.Key()), string(iter.Value()))
	}

	if ok := iter.First(); ok {
		log.Println("leveldb\t-", "iterator first key and value", string(iter.Key()), string(iter.Value()))
	}

}

// 使用前缀
func TestSimpleIteratorWithPrefix(t *testing.T) {
	db := newDb()
	defer db.Close()

	batch := new(leveldb.Batch)
	batch.Put([]byte("hash_2"), []byte("1"))
	batch.Put([]byte("hash_1"), []byte("2"))
	batch.Put([]byte("hash_3"), []byte("3"))
	batch.Put([]byte("hash_4"), []byte("4"))

	if err := db.Write(batch, nil); err != nil {
		log.Println("leveldb\t-", "iterator write with prefix error", err.Error())
	}

	iter := db.NewIterator(util.BytesPrefix([]byte("hash_")), nil)
	for iter.Next() {
		key := string(iter.Key())
		value := string(iter.Value())
		log.Println("leveldb\t-", "iterator seek with prefix", key, value)
	}
}

// 使用filter,不起作用
func TestSimpleFilter(t *testing.T) {
	opt := &opt2.Options{
		Filter: filter.NewBloomFilter(3)}

	db, err := leveldb.OpenFile("demo-filter-db", opt)
	defer db.Close()

	if err != nil {
		log.Println("leveldb\t-", "OpenFile with filter error", err.Error())
	}

	batch := new(leveldb.Batch)
	batch.Put([]byte("ssd-1"), []byte("v11111"))
	batch.Put([]byte("ssd-2"), []byte("v21111"))
	batch.Put([]byte("ssd-3"), []byte("v31111"))
	db.Write(batch, nil)

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := string(iter.Key())
		value := string(iter.Value())
		log.Println("leveldb\t-", "filter and iterator seek", key, value)
	}
}

func UseOptions() {

}
