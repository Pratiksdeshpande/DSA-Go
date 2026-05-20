# LRU Cache (Least Recently Used Cache)

## Problem Statement

Design a data structure that follows the constraints of a **Least Recently Used (LRU) cache**.

Implement the `LRUCache` class:

- `LRUCache(int capacity)` - Initialize the LRU cache with **positive** size capacity.
- `int get(int key)` - Return the value of the key if it exists, otherwise return `-1`.
- `void put(int key, int value)` - Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, **evict** the least recently used key.

**Important:** The functions `get` and `put` must each run in **O(1)** average time complexity.

---

## Input/Output Examples

### Example 1:

```
Input:
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]

Output:
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation:
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1);  // cache is {1=1}
lRUCache.put(2, 2);  // cache is {1=1, 2=2}
lRUCache.get(1);     // return 1, cache is {2=2, 1=1} (1 becomes most recently used)
lRUCache.put(3, 3);  // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);     // returns -1 (not found)
lRUCache.put(4, 4);  // LRU key was 1, evicts key 1, cache is {3=3, 4=4}
lRUCache.get(1);     // return -1 (not found)
lRUCache.get(3);     // return 3, cache is {4=4, 3=3}
lRUCache.get(4);     // return 4, cache is {3=3, 4=4}
```

### Example 2:

```
Input:
["LRUCache", "put", "put", "put", "put", "get", "get"]
[[2], [2, 1], [1, 1], [2, 3], [4, 1], [1], [2]]

Output:
[null, null, null, null, null, -1, 3]

Explanation:
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(2, 1);  // cache is {2=1}
lRUCache.put(1, 1);  // cache is {2=1, 1=1}
lRUCache.put(2, 3);  // Update existing key 2, cache is {1=1, 2=3}
lRUCache.put(4, 1);  // LRU key was 1, evicts key 1, cache is {2=3, 4=1}
lRUCache.get(1);     // return -1 (not found, was evicted)
lRUCache.get(2);     // return 3
```

---

## Constraints

- `1 <= capacity <= 3000`
- `0 <= key <= 10^4`
- `0 <= value <= 10^5`
- At most `2 * 10^5` calls will be made to `get` and `put`

---

## Solution Approach

### Data Structure Choice

To achieve **O(1)** time complexity for both `get` and `put` operations, we need:

1. **Hash Map** - For O(1) key lookup
2. **Doubly Linked List** - For O(1) insertion and deletion at any position

### Why Doubly Linked List?

| Operation | Singly Linked List | Doubly Linked List |
|-----------|-------------------|-------------------|
| Remove node (given pointer) | O(n) - need to find prev | O(1) - have prev pointer |
| Add to front | O(1) | O(1) |
| Remove from tail | O(n) | O(1) |

### Architecture

```
HashMap: key -> Node pointer (for O(1) access)

Doubly Linked List (for maintaining access order):

    head <-> Node1 <-> Node2 <-> Node3 <-> tail
    (dummy)   (MRU)              (LRU)    (dummy)

- Head side = Most Recently Used (MRU)
- Tail side = Least Recently Used (LRU)
```

### Algorithm

#### `get(key)`:
1. If key doesn't exist in map → return -1
2. Move the node to front (mark as most recently used)
3. Return the value

#### `put(key, value)`:
1. If key exists:
   - Update the value
   - Move node to front
   - **Return** (important!)
2. If key doesn't exist:
   - Create new node
   - Add to hash map
   - Add node to front of list
   - If capacity exceeded:
     - Remove LRU node (from tail)
     - Delete from hash map

### Helper Functions

| Function | Purpose | Time Complexity |
|----------|---------|-----------------|
| `addToFront(node)` | Add node right after head | O(1) |
| `removeNode(node)` | Unlink node from list | O(1) |
| `moveToFront(node)` | Remove + Add to front | O(1) |
| `removeLRU()` | Remove node before tail | O(1) |

---

## Time & Space Complexity

| Operation | Time | Space |
|-----------|------|-------|
| `get()` | O(1) | O(1) |
| `put()` | O(1) | O(1) |
| **Overall Space** | - | O(capacity) |

---

## Points to Remember for Senior Level Interviews

### 1. **Design Decisions to Discuss**

- **Why dummy head and tail nodes?**
  - Eliminates edge cases for empty list operations
  - No null checks needed when adding/removing nodes
  - Cleaner, less error-prone code

- **Why Doubly Linked List over Singly Linked List?**
  - O(1) deletion when we have the node pointer
  - Singly linked list would require O(n) to find previous node

### 2. **Common Bugs to Avoid**

```go
// BUG 1: Forgetting to return after updating existing key
if node, exists := lru.cache[key]; exists {
    node.value = val
    lru.moveToFront(node)
    return  // ← Don't forget this!
}

// BUG 2: Wrong pointer reassignment in removeNode
// WRONG:
node.prev = prevNode
node.next = nextNode

// CORRECT:
prevNode.next = nextNode
nextNode.prev = prevNode

// BUG 3: Using moveToFront for new nodes (causes nil pointer)
// For NEW nodes: use addToFront() only
// For EXISTING nodes: use moveToFront() (which does remove + add)
```

### 3. **Thread Safety Discussion**

In production, discuss how to make it thread-safe:
- **Read-Write Locks (RWMutex)**: Allow concurrent reads, exclusive writes
- **Sharding**: Partition cache into multiple segments with separate locks
- **Lock-free structures**: Using CAS operations (advanced)

```go
type ThreadSafeLRUCache struct {
    sync.RWMutex
    LRUCache
}

func (lru *ThreadSafeLRUCache) Get(key int) int {
    lru.RLock()
    defer lru.RUnlock()
    return lru.get(key)  // Note: This is still not fully correct!
}
```

**Caveat**: Even `get()` modifies the list (moves to front), so simple RWMutex doesn't work perfectly. You'd need full mutex or more sophisticated synchronization.

### 4. **Production Considerations**

| Aspect | Consideration |
|--------|---------------|
| **TTL (Time-To-Live)** | Add expiration timestamp to nodes |
| **Size-based eviction** | Track memory usage, not just count |
| **Metrics** | Hit rate, miss rate, eviction count |
| **Persistence** | WAL (Write-Ahead Log) for durability |
| **Distributed Cache** | Consistent hashing for sharding |

### 5. **Follow-up Questions to Expect**

1. **"How would you implement LFU (Least Frequently Used) cache?"**
   - Need additional frequency tracking
   - Use HashMap + LinkedHashSet per frequency

2. **"How would you handle cache stampede?"**
   - Locking per key
   - Probabilistic early expiration
   - Background refresh

3. **"What if values are large objects?"**
   - Store pointers/references
   - Consider size-based eviction instead of count-based

4. **"How to test this implementation?"**
   - Unit tests for edge cases (empty cache, single capacity)
   - Concurrency tests
   - Benchmark tests for performance

### 6. **Real-World Applications**

- **Database query caching** (MySQL query cache)
- **Web page caching** (CDN, browser cache)
- **DNS caching**
- **CPU cache replacement policy**
- **Session management**
- **Redis/Memcached internal implementation**

---

## Code Implementation (Go)

```go
type Node struct {
    key   int
    value int
    prev  *Node
    next  *Node
}

type LRUCache struct {
    capacity int
    cache    map[int]*Node
    head     *Node  // dummy head
    tail     *Node  // dummy tail
}

func Constructor(capacity int) LRUCache {
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

func (lru *LRUCache) Get(key int) int {
    if node, exists := lru.cache[key]; exists {
        lru.moveToFront(node)
        return node.value
    }
    return -1
}

func (lru *LRUCache) Put(key, value int) {
    if node, exists := lru.cache[key]; exists {
        node.value = value
        lru.moveToFront(node)
        return
    }

    newNode := &Node{key: key, value: value}
    lru.cache[key] = newNode
    lru.addToFront(newNode)

    if len(lru.cache) > lru.capacity {
        lruNode := lru.removeLRU()
        delete(lru.cache, lruNode.key)
    }
}

func (lru *LRUCache) addToFront(node *Node) {
    node.next = lru.head.next
    node.prev = lru.head
    lru.head.next.prev = node
    lru.head.next = node
}

func (lru *LRUCache) removeNode(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (lru *LRUCache) moveToFront(node *Node) {
    lru.removeNode(node)
    lru.addToFront(node)
}

func (lru *LRUCache) removeLRU() *Node {
    lruNode := lru.tail.prev
    lru.removeNode(lruNode)
    return lruNode
}
```

---

## LeetCode Reference

- **Problem**: [146. LRU Cache](https://leetcode.com/problems/lru-cache/)
- **Difficulty**: Medium
- **Tags**: Hash Table, Linked List, Design, Doubly-Linked List

