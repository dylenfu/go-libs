package kv

import (
	"hash/fnv"
	"sync"
	"time"
)

const USERS_SHARD_COUNT int = 1024

type KVCacheShardFunc func(key interface{}, shardCnt int) uint
type KVCacheLoader func(key interface{}) interface{}

type KVCachePool struct {
	slots     []*KVCache
	mutexes   []sync.Mutex
	shardCnt  int
	shardFunc KVCacheShardFunc
	lifeTime  time.Duration
}

func NewKVCachePool(lifeTime time.Duration, shardCnt int, shardFunc KVCacheShardFunc) *KVCachePool {
	pool := KVCachePool{
		slots:     make([]*KVCache, shardCnt),
		shardCnt:  shardCnt,
		shardFunc: shardFunc,
		lifeTime:  lifeTime,
	}
	for i := 0; i < shardCnt; i += 1 {
		pool.slots[i] = NewKVCache()
	}
	return &pool
}

func (c *KVCachePool) shardIndex(key interface{}) int {
	shardIndex := int(c.shardFunc(key, c.shardCnt))
	if shardIndex < 0 {
		shardIndex = -shardIndex
	}
	shardIndex = shardIndex % c.shardCnt
	return shardIndex
}

func (c *KVCachePool) Get(key interface{}) interface{} {
	return c.slots[c.shardIndex(key)].Get(key)
}

func (c *KVCachePool) Set(key interface{}, value interface{}, lifetime time.Duration) {
	c.slots[c.shardIndex(key)].Set(key, value, lifetime)
}

func (c *KVCachePool) Del(key interface{}) {
	c.slots[c.shardIndex(key)].Del(key)
}

func (c *KVCachePool) Incr(key interface{}, value int, lifetime time.Duration) int {
	return c.slots[c.shardIndex(key)].Incr(key, value, lifetime)
}

func (c *KVCachePool) SetWithCallback(key interface{}, lifetime time.Duration, callback KVCacheHandler) interface{} {
	return c.slots[c.shardIndex(key)].SetWithCallback(key, lifetime, callback)
}

func (c *KVCachePool) Walk(callback KVCacheHandler) {
	for i := 0; i < len(c.slots); i++ {
		if c.slots[i].Len() < 1 {
			continue
		}
		c.slots[i].Walk(callback)
	}
}

func (c *KVCachePool) GetWithCleanCallback(key interface{}, cleanCallback KVCacheHandler) (result interface{}) {
	return c.slots[c.shardIndex(key)].GetWithCleanCallback(key, cleanCallback)
}

func (c *KVCachePool) DelWithCallback(key interface{}, cleanCallback KVCacheHandler) bool {
	return c.slots[c.shardIndex(key)].DelWithCallback(key, cleanCallback)
}

func (c *KVCachePool) TruncateWithCallback(cleanCallback KVCacheHandler) {
	var wg sync.WaitGroup
	for i := 0; i < len(c.slots); i++ {
		if c.slots[i].Len() < 1 {
			continue
		}
		wg.Add(1)
		go func(i int) {
			c.slots[i].TruncateWithCallback(cleanCallback)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Shard(key interface{}, shardCnt int) uint {
	return hash(key) % uint(shardCnt)
}

func hash(s interface{}) uint {
	str, ok := s.(string)
	if !ok {
		return 0
	}
	h := fnv.New32a()
	h.Write([]byte(str))
	return uint(h.Sum32())
}
