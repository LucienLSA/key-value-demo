package toy_kv

import "sync"

type CommonMap[Key, Value any] struct {
	// 内部使用一个map来存储数据
	m  map[any]Value
	mu *sync.Mutex
}

func (c *CommonMap[Key, Value]) Load(key Key) (Value, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.m[key]
	return v, ok
}

func (c *CommonMap[Key, Value]) Store(key Key, value Value) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

func (c *CommonMap[Key, Value]) Delete(key Key) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}

func (c *CommonMap[Key, Value]) Range(f func(key Key, value Value) bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.m {
		if !f(k.(Key), v) {
			return
		}
	}
}

func NewCommonMap[Key, Value any]() IMap[Key, Value] {
	return &CommonMap[Key, Value]{
		mu: &sync.Mutex{},
		m:  make(map[any]Value),
	}
}
