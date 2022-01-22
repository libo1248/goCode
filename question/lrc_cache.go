package main

import "fmt"

/*
LRU(最近最少使用) Cache是一种缓存机制的实现，采用丢弃缓存中最近最少使用的内容块的策略来实现高速度低占用的缓存系统。
例如，现有数据序列ABCDEDF, 该系统运行效果
访问A：[A][ ][ ][ ] - A最近最少使用
访问B：[A][B][ ][ ] - A最近最少使用，其次B
访问C：[A][B][C][ ] - A最近最少使用，其次B→C
访问D：[A][B][C][D] - A最近最少使用，其次B->C→D
访问E：[E][B][C][D] - B最近最少使用，其次C->D→E
访问D：[E][B][C][D] - B最近最少使用，其次C->E→D
访问F：[E][F][C][D] - C最近最少使用，其次C->E->D→F
Get(key) 获取key对应的存在cache中的值，如没有请返回-1
Put（key，value）遵循最近最少使用策略，将需要缓存的值更新到cache中。
Note：缓存数据为正整数，请以此设计容量为4的LRU Cache
*/

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*lruNode
	head, tail *lruNode
}

type lruNode struct {
	key, value int
	prev, next *lruNode
}

func newLruNode(key, value int) *lruNode {
	return &lruNode{
		key:   key,
		value: value,
	}
}

func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		cache:    map[int]*lruNode{},
		capacity: capacity,
		head:     newLruNode(0, 0),
		tail:     newLruNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head

	return lru
}

func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	node := lru.cache[key]
	lru.moveToHead(node)
	return node.value
}

func (lru *LRUCache) Put(key int, value int) {
	node, ok := lru.cache[key]
	if !ok {
		node = newLruNode(key, value)
		lru.cache[key] = node
		lru.addToHead(node)
		lru.size++
		if lru.size > lru.capacity {
			removed := lru.removeTail()
			delete(lru.cache, removed.key)
			lru.size--
		}
	} else {
		node.value = value
		lru.moveToHead(node)
	}
}

func (lru *LRUCache) addToHead(node *lruNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) removeNode(node *lruNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) moveToHead(node *lruNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeTail() *lruNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}

func (lru *LRUCache) Print() {
	node := lru.head.next
	fmt.Println("----------------")
	for node != lru.tail {
		fmt.Println(node.key, node.value)
		node = node.next
	}
}
