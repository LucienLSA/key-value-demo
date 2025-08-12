package toy_kv

// 定义Map接口

type IMap[Key, Value any] interface {
	// Load 从KV中读取一个Key，获取它的Value以及是否存在
	Load(key Key) (Value, bool)

	// Store 向KV中写入一个Key-Value
	Store(key Key, value Value)

	// Delete 删除KV中的一个Key
	Delete(key Key)

	// Range 遍历KV中的所有Key-Value
	// 遍历过程中，如果f返回false，遍历会停止
	Range(f func(key Key, value Value) bool)
}
