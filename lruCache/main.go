// https://leetcode.com/problems/lru-cache
package main

import (
	"fmt"
)

type Node struct {
	key        int
	value      int
	next, prev *Node
}

type LRUCache struct {
	head, tail *Node
	addresses  map[int]*Node
	capacity   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		addresses: make(map[int]*Node),
		capacity:  capacity,
	}
}

func (this *LRUCache) evict(key int) *Node {
	node, ok := this.addresses[key]
	if !ok {
		return nil
	}

	if this.head == node {
		this.head = node.next
	}
	if this.tail == node {
		this.tail = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	delete(this.addresses, node.key)

	node.next = nil
	node.prev = nil

	return node
}

func (this *LRUCache) Get(key int) int {
	node := this.evict(key)
	if node == nil {
		return -1
	}

	this.Put(node.key, node.value)

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	this.evict(key)
	if len(this.addresses) >= this.capacity {
		this.evict(this.tail.value)
	}

	node := &Node{
		key:   key,
		value: value,
	}
	if this.head != nil {
		this.head.prev = node
	}
	node.next = this.head
	this.head = node
	if this.tail == nil {
		this.tail = node
	}
	this.addresses[node.key] = node
}

func main() {
	cache := Constructor(2)

	fmt.Println(cache.Get(2))
	cache.Put(2, 6)
	fmt.Println(cache.Get(1))
	cache.Put(1, 5)
	cache.Put(1, 2)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}
