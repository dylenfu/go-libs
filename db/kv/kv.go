package kv

import (
	_ "fmt"
	"sync"
	"time"
)

type KVCacheItem struct {
	expiredAt time.Duration
	value     interface{}
}

func (item *KVCacheItem) isExpired(unixNanoPtr *time.Duration) bool {
	if item.expiredAt < 0 {
		return false
	}
	var unixNano time.Duration
	if unixNanoPtr != nil {
		unixNano = *unixNanoPtr
	} else {
		unixNano = time.Duration(time.Now().UnixNano())
	}
	return unixNano > item.expiredAt
}

type KVCache struct {
	sync.RWMutex
	cache map[interface{}]KVCacheItem
}

type KVCacheHandler func(key, oldValue interface{}, oldIsExpired bool) (newValue interface{})

func NewKVCache() *KVCache {
	return &KVCache{
		cache: make(map[interface{}]KVCacheItem),
	}
}

func (c *KVCache) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.cache)
}

func (c *KVCache) Set(key interface{}, value interface{}, lifetime time.Duration) {
	c.Lock()
	defer c.Unlock()
	unixNano := time.Duration(time.Now().UnixNano())
	var expiredAt time.Duration
	if lifetime > 0 {
		expiredAt = unixNano + lifetime
	} else {
		expiredAt = -1
	}
	c.cache[key] = KVCacheItem{
		expiredAt: expiredAt,
		value:     value,
	}
}

func (c *KVCache) Get(key interface{}) (result interface{}) {
	c.RLock()
	defer c.RUnlock()
	unixNano := time.Duration(time.Now().UnixNano())
	item, existing := c.cache[key]
	result = nil
	if !existing {
		return
	}
	if !item.isExpired(&unixNano) {
		result = item.value
		return
	}
	return
}

func (c *KVCache) Del(key interface{}) {
	c.Lock()
	defer c.Unlock()
	delete(c.cache, key)
}

func (c *KVCache) Incr(key interface{}, value int, lifetime time.Duration) int {
	incrFunc := func(key, oldValue interface{}, oldIsExpired bool) (newValue interface{}) {
		newValue = value + oldValue.(int)
		return
	}
	return c.SetWithCallback(key, lifetime, incrFunc).(int)
}

func (c *KVCache) SetWithCallback(key interface{}, lifetime time.Duration, callback KVCacheHandler) interface{} {
	c.Lock()
	defer c.Unlock()
	unixNano := time.Duration(time.Now().UnixNano())
	var expiredAt time.Duration
	if lifetime > 0 {
		expiredAt = unixNano + lifetime
	} else {
		expiredAt = -1
	}
	var oldValue interface{}
	var oldIsExpired bool
	if oldItem, ok := c.cache[key]; ok {
		oldValue = oldItem.value
		oldIsExpired = oldItem.isExpired(&unixNano)
	}
	newValue := callback(key, oldValue, oldIsExpired)
	c.cache[key] = KVCacheItem{
		expiredAt: expiredAt,
		value:     newValue,
	}
	return newValue
}

func (c *KVCache) Walk(callback KVCacheHandler) {
	var unlocked = false
	c.RLock()
	defer func() {
		if !unlocked {
			c.RUnlock()
		}
	}()
	cache := make([][2]interface{}, len(c.cache))
	i := 0
	for key, item := range c.cache {
		cache[i][0] = key
		cache[i][1] = item.value
		i++
	}
	c.RUnlock()
	unlocked = true

	for _, item := range cache {
		callback(item[0], item[1], false)
	}
}

func (c *KVCache) GetWithCleanCallback(key interface{}, cleanCallback KVCacheHandler) (result interface{}) {
	c.RLock()
	unixNano := time.Duration(time.Now().UnixNano())
	item, existing := c.cache[key]
	result = nil
	if !existing {
		c.RUnlock()
		return
	}
	if !item.isExpired(&unixNano) {
		result = item.value
		c.RUnlock()
		return
	}
	c.RUnlock()

	// cache expired ?
	unixNano = time.Duration(time.Now().UnixNano())
	c.Lock()
	item, existing = c.cache[key]
	if !existing { // already cleaned by other goroutine?
		c.Unlock()
		return
	}
	isExpired := item.isExpired(&unixNano)
	if !isExpired { // already updated by other goroutine?
		result = item.value
		c.Unlock()
		return
	}
	delete(c.cache, key)
	c.Unlock()
	go cleanCallback(key, item.value, isExpired)
	return
}

func (c *KVCache) DelWithCallback(key interface{}, cleanCallback KVCacheHandler) bool {
	c.Lock()
	item, existing := c.cache[key]
	if !existing { // already cleaned by other goroutine?
		c.Unlock()
		return false
	}
	delete(c.cache, key)
	c.Unlock()
	go cleanCallback(key, item.value, true /* fake value */)
	return true
}

func (c *KVCache) TruncateWithCallback(cleanCallback KVCacheHandler) {
	c.Lock()
	defer c.Unlock()
	keys := []interface{}{}
	for k, _ := range c.cache {
		keys = append(keys, k)
	}
	for _, k := range keys {
		cleanCallback(k, c.cache[k].value, true)
		delete(c.cache, k)
	}
}
