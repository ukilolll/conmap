package src

import "sync"

type IoMap[K comparable, V any] struct {
	Map map[K]V
	Mu  sync.RWMutex
}

func (b *IoMap[K, V]) Read(key K) (V, bool) {
	b.Mu.RLock()
	defer b.Mu.RUnlock()
	val, ok := b.Map[key]
	return val, ok
}

func (b *IoMap[K, V]) Write(key K, value V) {
	b.Mu.Lock()
	defer b.Mu.Unlock()
	b.Map[key] = value
}

func (b *IoMap[K, V]) Edit(callback func(m *map[K]V)) {
	b.Mu.Lock()
	defer b.Mu.Unlock()
	callback(&b.Map)
}

func (b *IoMap[K, V]) Range(callback func(key K, value V)) {
	b.Mu.RLock()
	defer b.Mu.RUnlock()
	for k, v := range b.Map {
		callback(k, v)
	}
}

func NewIoMap[K comparable, V any]() *IoMap[K, V] {
	return &IoMap[K, V]{Map: make(map[K]V)}
}