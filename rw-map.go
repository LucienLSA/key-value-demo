package toy_kv

import (
	"sync"
	"unsafe"
)

type RWMap[Key, Value any] struct {
	tables map[uint]*Table[Key, Value]
	length uint
}

type Table[Key, Value any] struct {
	mu    *sync.RWMutex
	lines map[any]Value
}

func (t *Table[Key, Value]) get(key Key) (Value, bool) {
	v, ok := t.lines[key]
	return v, ok
}

func (t *Table[Key, Value]) set(key Key, value Value) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.lines[key] = value

}

func (t *Table[Key, Value]) del(key Key) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.lines, key)
}

func (t *Table[Key, Value]) tRange(f func(key Key, value Value) bool) bool {
	for k, v := range t.lines {
		if !f(k.(Key), v) {
			return false
		}
	}
	return true
}

func hash[Key any](key Key) uint {
	gc := &key
	start := uintptr(unsafe.Pointer(gc))
	offset := unsafe.Sizeof(key)
	sizeofByte := unsafe.Sizeof(byte(0))
	hashSum := uint(0)
	for ptr := start; ptr < start+offset; ptr += sizeofByte {
		h := *(*byte)(unsafe.Pointer(ptr))
		// 哈希算法
		hashSum = uint(h) + (hashSum << 6) + (hashSum << 16) - hashSum
	}
	return hashSum
}

func (R *RWMap[Key, Value]) Load(key Key) (Value, bool) {
	index := hash(key) % R.length
	return R.tables[index].get(key)

}

func (R *RWMap[Key, Value]) Store(key Key, value Value) {
	index := hash(key) % R.length
	R.tables[index].set(key, value)

}

func (R *RWMap[Key, Value]) Delete(key Key) {
	index := hash(key) % R.length
	R.tables[index].del(key)
}

func (R *RWMap[Key, Value]) Range(f func(key Key, value Value) bool) {
	for _, table := range R.tables {
		if !table.tRange(f) {
			break
		}
	}
}
func NewRWMap[Key, Value any](length uint) IMap[Key, Value] {
	rwMap := &RWMap[Key, Value]{
		tables: make(map[uint]*Table[Key, Value]),
		length: length,
	}
	for i := uint(0); i < length; i++ {
		rwMap.tables[i] = &Table[Key, Value]{
			mu:    &sync.RWMutex{},
			lines: make(map[any]Value),
		}
	}
	return rwMap

}
