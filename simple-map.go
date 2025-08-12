package toy_kv

import "sync"

type SimpleMap[Key, Value any] struct {
	// 内部使用一个map来存储数据
	m  map[any]Value
	mu *sync.RWMutex
}

func (s *SimpleMap[Key, Value]) Load(key Key) (Value, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.m[key]
	return v, ok
}

func (s *SimpleMap[Key, Value]) Store(key Key, value Value) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SimpleMap[Key, Value]) Delete(key Key) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
}

func (s *SimpleMap[Key, Value]) Range(f func(key Key, value Value) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for k, v := range s.m {
		if !f(k.(Key), v) {
			return
		}
	}
}

func NewSimpleMap[Key, Value any]() IMap[Key, Value] {
	return &SimpleMap[Key, Value]{
		mu: &sync.RWMutex{},
		m:  make(map[any]Value),
	}
}
