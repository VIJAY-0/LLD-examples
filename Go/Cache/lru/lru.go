package lru

import (
	"fmt"
	"sync"
)

type Node struct {
	key   int
	next  *Node
	prev  *Node
	value int
}

type LRU struct {
	capacity int
	mp       map[int]*Node
	start    *Node
	end      *Node
	mu       sync.Mutex
}

func NewLruCache(cap int) *LRU {

	startNode := new(Node)
	endNode := new(Node)

	startNode.next = endNode
	endNode.prev = startNode

	return &LRU{
		capacity: cap,
		mp:       make(map[int]*Node),
		start:    startNode,
		end:      endNode,
	}
}

func (lru *LRU) printCahe() {
	startptr := lru.start
	fmt.Print("[ ")
	for startptr.next != lru.end {
		startptr = startptr.next
		fmt.Printf(" {%v:%v} ", startptr.key, startptr.value)
	}
	fmt.Print(" ]\n")
}

func (lru *LRU) get(key int) (int, bool) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if elem, found := lru.mp[key]; found {

		elem.prev.next = elem.next
		elem.next.prev = elem.prev

		elem.next = lru.start.next
		elem.next.prev = elem
		lru.start.next = elem
		elem.prev = lru.start

		return elem.value, true
	}

	return -1, false
}

func (lru *LRU) put(key int, val int) bool {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if elem, found := lru.mp[key]; found {
		elem.next.prev = elem.prev
		elem.prev.next = elem.next

		elem.next = lru.start.next
		elem.next.prev = elem
		elem.prev = lru.start
		elem.prev.next = elem

		elem.value = val

	} else {
		// fmt.Printf("in 2 and len of map is :%v\n", len(lru.mp))

		if len(lru.mp) == lru.capacity {

			node := lru.end.prev
			node.prev.next = node.next
			node.next.prev = node.prev

			node.next = lru.start.next
			node.prev = lru.start
			node.next.prev = node
			node.prev.next = node

			delete(lru.mp, node.key)
			node.key = key
			node.value = val
			lru.mp[key] = node
		} else {
			node := new(Node)
			node.value = val
			node.key = key
			lru.mp[key] = node

			node.next = lru.start.next
			node.next.prev = node
			node.prev = lru.start
			lru.start.next = node

		}
	}

	return true
}
