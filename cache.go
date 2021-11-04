// Package cache provides ...
package cache

import "sync"

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Del(key string)
	DelOldest()
	Len() int
}

// 默认允许占用的最大内存
const DefaultMaxBytes = 1 << 29

// 并发安全缓存
type safeCache struct {
	m          sync.RWMutex
	cache      Cache
	nget, nhit int
}

// 新建缓存
func newSafeCache(cache Cache) *safeCache {
	return &safeCache{
		cache: cache,
	}
}

// 设置
func (sc *safeCache) set(key string, value interface{}) {
	sc.m.Lock()
	defer sc.m.Unlock()
	sc.cache.Set(key, value)
}

//获取
func (sc *safeCache) get(key string) interface{} {
	sc.m.RLock()
	defer sc.m.RUnlock()
	sc.nget++
	if sc.cache == nil {
		return nil
	}
	v := sc.cache.Get(key)
	if v != nil {
		sc.nhit++
	}
	return v
}

// 状态
func (sc *safeCache) stat() *Stat {
	sc.m.RLock()
	defer sc.m.RUnlock()
	return &Stat{
		NHit: sc.nhit,
		NGet: sc.nget,
	}
}

type Stat struct {
	NHit, NGet int
}
