package main

import "fmt"

type Node struct {
	key   int
	value int
	freq  int
	prev  *Node
	next  *Node
}

type DLL struct {
	size int
	head *Node
	tail *Node
}

type LFUCache struct {
	capacity  int
	size      int
	minFreq   int
	keyToNode map[int]*Node
	freqToDLL map[int]*DLL
}

func NewDLL() *DLL {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &DLL{
		size: 0,
		head: head,
		tail: tail,
	}
}

func (dll *DLL) addNode(node *Node) {

	node.next = dll.head.next
	node.prev = dll.head

	dll.head.next.prev = node
	dll.head.next = node

	dll.size++
}

func (dll *DLL) removeNode(node *Node) {
	prevNode := node.prev
	nextNode := node.next

	prevNode.next = nextNode
	nextNode.prev = prevNode
	dll.size--
}

func (dll *DLL) removeLastNode() *Node {
	if dll.size == 0 {
		return nil
	}
	lastNode := dll.tail.prev
	dll.removeNode(lastNode)
	return lastNode
}

func (lfuCache *LFUCache) updateFrequency(node *Node) {

	oldFreq := node.freq

	oldDLL := lfuCache.freqToDLL[oldFreq]
	oldDLL.removeNode(node)

	if lfuCache.minFreq == oldFreq && oldDLL.size == 0 {
		lfuCache.minFreq++
	}

	node.freq++

	newDLL, exists := lfuCache.freqToDLL[node.freq]
	if !exists {
		newDLL = NewDLL()
		lfuCache.freqToDLL[node.freq] = newDLL
	}
	newDLL.addNode(node)

}

func (lfuCache *LFUCache) Get(key int) int {
	node, exists := lfuCache.keyToNode[key]

	if !exists {
		return -1
	}

	lfuCache.updateFrequency(node)
	return node.value
}

func (lfuCache *LFUCache) Put(key, value int) {
	if lfuCache.capacity == 0 {
		return
	}

	if node, exists := lfuCache.keyToNode[key]; exists {
		node.value = value
		lfuCache.updateFrequency(node)
		return
	}

	if lfuCache.size == lfuCache.capacity {
		dll := lfuCache.freqToDLL[lfuCache.minFreq]
		nodeToRemove := dll.removeLastNode().key
		delete(lfuCache.keyToNode, nodeToRemove)
		lfuCache.size--
	}

	newNode := &Node{
		key:   key,
		value: value,
		freq:  1,
	}

	lfuCache.keyToNode[key] = newNode

	dll, exists := lfuCache.freqToDLL[1]
	if !exists {
		dll = NewDLL()
		lfuCache.freqToDLL[1] = dll
	}

	dll.addNode(newNode)
	lfuCache.minFreq = 1
	lfuCache.size++
}

func main() {

	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	fmt.Println(cache.Get(1))

	cache.Put(3, 3)

	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))

	cache.Put(4, 4)

	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))

}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:  capacity,
		keyToNode: make(map[int]*Node),
		freqToDLL: make(map[int]*DLL),
	}
}
