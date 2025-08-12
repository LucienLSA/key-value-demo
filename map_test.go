package toy_kv

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

// 普通Mutex锁
func CommonMapTestFunc() {
	wg := &sync.WaitGroup{}
	wg.Add(100 * 10000)
	m := NewCommonMap[int, int]()
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				m.Store(i, i)
				m.Load(i)
				wg.Done()
			}
		}()
	}
	wg.Wait()
}
func BenchmarkCommonMapTestFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CommonMapTestFunc()
	}
}

// RWMutex读写锁

func SimpleMapTestFunc() {
	wg := &sync.WaitGroup{}
	wg.Add(100 * 10000)
	m := NewSimpleMap[int, int]()
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				m.Store(i, i)
				m.Load(i)
				wg.Done()
			}
		}()
	}
	wg.Wait()
}
func BenchmarkSimpleMapTestFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleMapTestFunc()
	}
}

func TestFune(t *testing.T) {
	x := 0
	uptr := unsafe.Pointer(&x)
	bptr := *(*[]byte)(uptr)
	fmt.Printf("len of bptr: %v\n", len(bptr))
}

// 读写均衡
func RWMapTestFunc() {
	wg := &sync.WaitGroup{}
	wg.Add(100 * 10000)
	m := NewRWMap[int, int](1007)
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				m.Store(i, i)
				m.Load(i)
				wg.Done()
			}
		}()
	}
	wg.Wait()
}

func BenchmarkRWMapTestFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RWMapTestFunc()
	}
}

// sync.Map

func SyncMapTestFunc() {
	wg := &sync.WaitGroup{}
	wg.Add(100 * 10000)
	m := sync.Map{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				m.Store(i, i)
				m.Load(i)
				wg.Done()
			}
		}()
	}
	wg.Wait()
}

func BenchmarkSyncMapTestFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SyncMapTestFunc()
	}
}
