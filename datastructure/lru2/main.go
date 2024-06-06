package main

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

	var victim any
	ent := &entry{key, value}
	elm := l.evictions.PushFront(ent)
	l.items[key] = elm

	if l.needEvict() {
		victim = l.evictions.Back()
		l.removeOldest()
	}

	return victim
}

func (l *LRU) needEvict() bool {
	return l.evictions.Len() > l.capacity
}

func (l *LRU) removeOldest() {
	elm := l.evictions.Back()
	if elm != nil {
		l.evictions.Remove(elm)
		delete(l.items, elm.Value.(*entry).key)
	}
}
