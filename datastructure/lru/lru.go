package lru

import (
	"container/list"
	"sync"
)

type LRU struct {
	capacity  int
	evictions *list.List
	items     map[any]*list.Element
	mutex     sync.RWMutex
}

type entry struct {
	key   any
	value any
}

func NewLRU(c int) *LRU {
	return &LRU{
		capacity:  c,
		evictions: list.New(),
		items:     make(map[any]*list.Element),
	}
}

func (l *LRU) Insert(key, value any) any {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	var victim any // 満杯だったときに削除される値
	ent := &entry{key, value}
	elm := l.evictions.PushFront(ent)
	l.items[key] = elm

	if l.needEvict() {
		victim = l.evictions.Back()
		l.removeOldest()
	}

	return victim
}

func (l *LRU) Get(key any) any {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	if elm, ok := l.items[key]; ok {
		l.evictions.MoveToFront(elm)
		return elm.Value.(*entry).value
	}

	return nil
}

func (l *LRU) Len() int {
	return l.evictions.Len()
}

func (l *LRU) removeOldest() {
	elm := l.evictions.Back()
	if elm != nil {
		l.evictions.Remove(elm)
		delete(l.items, elm.Value.(*entry).key)
	}
}

func (l *LRU) needEvict() bool {
	return l.Len() > l.capacity
}

func (l *LRU) GetAll() []any {
	items := make([]any, 0, len(l.items))
	for _, item := range l.items {
		items = append(items, item.Value.(*entry).value)
	}

	return items
}
