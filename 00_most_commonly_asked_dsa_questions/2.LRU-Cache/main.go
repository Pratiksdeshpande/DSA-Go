package main

import "fmt"

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func main() {
	lruCache := constructor(3)

	lruCache.put(1, 100)
	fmt.Println("line number 38:", lruCache.get(1))

	lruCache.put(2, 200)
	lruCache.put(3, 300)
	lruCache.put(4, 400)

	fmt.Println("line number 44:", lruCache.get(1))
	fmt.Println("line number 45:", lruCache.get(2))

	lruCache.put(5, 500)
	fmt.Println("line number 48:", lruCache.get(5))

}

func (lru *LRUCache) get(key int) int {
	node, exists := lru.cache[key]

	if !exists {
		return -1
	}

	lru.moveToFront(node)
	return node.value
}

func (lru *LRUCache) put(key, val int) {

	if node, exists := lru.cache[key]; exists {
		node.value = val
		lru.moveToFront(node)
		return
	}

	newNode := &Node{
		key:   key,
		value: val,
	}

	lru.cache[key] = newNode

	lru.addTOFront(newNode)

	if len(lru.cache) > lru.capacity {
		lruNode := lru.removeLRU()
		delete(lru.cache, lruNode.key)
	}
}

func (lru *LRUCache) addTOFront(node *Node) {

	node.next = lru.head.next
	node.prev = lru.head

	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) removeNode(node *Node) {
	prevNode := node.prev
	nextNode := node.next

	prevNode.next = nextNode
	nextNode.prev = prevNode
}

func (lru *LRUCache) moveToFront(node *Node) {
	lru.removeNode(node)
	lru.addTOFront(node)
}

func (lru *LRUCache) removeLRU() *Node {
	lruNode := lru.tail.prev
	lru.removeNode(lruNode)
	return lruNode
}
