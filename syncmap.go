package syncmap

import "sync"

// SyncMap is a wrapper around sync.Map which uses generics to make accessing the map more convenient
type SyncMap[K comparable, V any] struct {
	syncMap sync.Map
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.syncMap.Load(key)
	if !ok {
		var v2 V
		return v2, false
	}
	return v.(V), ok
}

// Store sets the value for a key.
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.syncMap.Store(key, value)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, l := m.syncMap.LoadOrStore(key, value)
	return a.(V), l
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, l := m.syncMap.LoadAndDelete(key)
	if !l {
		var v2 V
		return v2, false
	}
	return v.(V), l
}

// Delete deletes the value for a key.
func (m *SyncMap[K, V]) Delete(key K) {
	m.syncMap.Delete(key)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.syncMap.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

// Len returns the number of items in the map
func (m *SyncMap[K, V]) Len() int {
	l := 0
	m.syncMap.Range(func(key, value any) bool {
		l++
		return true
	})
	return l
}

// Map returns the current contents of the map as a standard Go map
func (m *SyncMap[K, V]) Map() map[K]V {
	newMap := make(map[K]V)
	m.Range(func(key K, value V) bool {
		newMap[key] = value
		return true
	})
	return newMap
}

// Keys returns a slice containing the keys in the map
func (m *SyncMap[K, V]) Keys() []K {
	var keys []K
	m.syncMap.Range(func(key, value any) bool {
		keys = append(keys, key.(K))
		return true
	})
	return keys
}

// Items returns a slice containing the items in the map
func (m *SyncMap[K, V]) Items() []V {
	var items []V
	m.syncMap.Range(func(key, value any) bool {
		items = append(items, value.(V))
		return true
	})
	return items
}
