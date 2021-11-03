// Package fifo provides ...
package fifo

import (
	"container/list"

	"github.com/Garfield247/go_cache/cache"
)

//先进先出
// fifo 是一个FIFO cache,他不是并发安全的.
type fifo struct {
	//缓存的最大容量,单位字节
	// groupcache 使用的是最大存放entry个数
	maxBytes int

	// 当一个entry从缓存移除是调用该回调函数,默认为nil
	// groupcache中的key是任意的可比较类型;value是interface{}
	onEvicted func(key string, value interface{})

	//已使用字节数
	usedBytes int

	ll *list.List

	cache map[string]*list.Element
}

type entry struct {
	key   string
	value interface{}
}

func (e *entry) Len() int {
	return cache.CalcLen(e.value)
}
