# LFU Cache (Least Frequently Used Cache)

## Problem Statement

Design and implement a data structure for a **Least Frequently Used (LFU) cache**.

Implement the `LFUCache` class:

- `LFUCache(int capacity)` - Initializes the object with the capacity of the data structure.
- `int get(int key)` - Gets the value of the key if it exists in the cache. Otherwise, returns `-1`.
- `void put(int key, int value)` - Updates the value of the key if present, or inserts the key if not already present. When the cache reaches its capacity, it should invalidate and remove the **least frequently used** key before inserting a new item. For this problem, when there is a **tie** (i.e., two or more keys with the same frequency), the **least recently used** key would be invalidated.

**Important:** The functions `get` and `put` must each run in **O(1)** average time complexity.

---

## LFU vs LRU: Key Difference

| Aspect | LRU (Least Recently Used) | LFU (Least Frequently Used) |
|--------|---------------------------|----------------------------|
| **Eviction Criteria** | Time since last access | Number of accesses (frequency) |
| **Tracks** | Access order | Access count |
| **Tie Breaker** | N/A (order is unique) | Use LRU among same frequency |
| **Use Case** | Recent data is important | Popular data is important |

---

## Input/Output Examples

### Example 1:

```
Input:
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]

Output:
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

Explanation:
// cnt(x) = the use counter for key x
// cache=[] will show the last used order for tiebreakers (leftmost = most recent)

LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
lfu.get(1);      // return 1, cache=[1,2], cnt(1)=2, cnt(2)=1
lfu.put(3, 3);   // 2 is the LFU, invalidate 2, cache=[3,1], cnt(3)=1, cnt(1)=2
lfu.get(2);      // return -1 (not found)
lfu.get(3);      // return 3, cache=[3,1], cnt(3)=2, cnt(1)=2
lfu.put(4, 4);   // Both 1 and 3 have same cnt, but 1 is LRU, invalidate 1
                 // cache=[4,3], cnt(4)=1, cnt(3)=2
lfu.get(1);      // return -1 (not found)
lfu.get(3);      // return 3, cache=[3,4], cnt(3)=3, cnt(4)=1
lfu.get(4);      // return 4, cache=[4,3], cnt(4)=2, cnt(3)=3
```

### Example 2:

```
Input:
["LFUCache", "put", "put", "get", "get", "put", "get", "get"]
[[2], [1, 10], [2, 20], [1], [2], [3, 30], [2], [3]]

Output:
[null, null, null, 10, 20, null, 20, 30]

Explanation:
LFUCache lfu = new LFUCache(2);
lfu.put(1, 10);  // cache={1:10}, freq(1)=1
lfu.put(2, 20);  // cache={1:10, 2:20}, freq(1)=1, freq(2)=1
lfu.get(1);      // return 10, freq(1)=2
lfu.get(2);      // return 20, freq(2)=2
lfu.put(3, 30);  // Both have freq=2, evict LRU (node 1 was accessed before node 2)
                 // cache={2:20, 3:30}, freq(2)=2, freq(3)=1
lfu.get(2);      // return 20, freq(2)=3
lfu.get(3);      // return 30, freq(3)=2
```

---

## Constraints

- `0 <= capacity <= 10^4`
- `0 <= key <= 10^5`
- `0 <= value <= 10^9`
- At most `2 * 10^5` calls will be made to `get` and `put`

---

## Solution Approach

### Challenge

To achieve **O(1)** for both operations, we need:
1. O(1) key lookup
2. O(1) frequency update
3. O(1) finding the minimum frequency
4. O(1) eviction of LFU (and LRU among ties)

### Data Structures Used

```
┌─────────────────────────────────────────────────────────────────┐
│                         LFU Cache                               │
├─────────────────────────────────────────────────────────────────┤
│  1. keyToNode: map[int]*Node    → O(1) key lookup               │
│  2. freqToDLL: map[int]*DLL     → O(1) access to frequency list │
│  3. minFreq: int                → O(1) find minimum frequency   │
│  4. capacity, size: int         → Track capacity                │
└─────────────────────────────────────────────────────────────────┘
```

### Architecture Diagram

```
keyToNode (HashMap):
┌─────┬─────────┐
│ Key │  Node*  │
├─────┼─────────┤
│  1  │  ──────►│──┐
│  2  │  ──────►│──┼──┐
│  3  │  ──────►│──┼──┼──┐
└─────┴─────────┘  │  │  │
                   │  │  │
freqToDLL (HashMap of Doubly Linked Lists):
┌──────┬─────────────────────────────────────┐
│ Freq │            DLL (LRU order)          │
├──────┼─────────────────────────────────────┤
│  1   │ head <-> [Node3] <-> tail           │◄──┘
│      │          (MRU)    (LRU)             │
├──────┼─────────────────────────────────────┤
│  2   │ head <-> [Node2] <-> [Node1] <-> tail│◄──┘
│      │          (MRU)       (LRU)          │
└──────┴─────────────────────────────────────┘
         ▲
         │
    minFreq = 1 (points to frequency with LFU nodes)
```

### Node Structure

```go
type Node struct {
    key   int      // Key (needed for deletion from keyToNode)
    value int      // Value
    freq  int      // Current frequency count
    prev  *Node    // Previous node in DLL
    next  *Node    // Next node in DLL
}
```

### Why This Design?

| Requirement | Solution | Complexity |
|-------------|----------|------------|
| Find node by key | `keyToNode` HashMap | O(1) |
| Get all nodes with same frequency | `freqToDLL` HashMap | O(1) |
| Find LRU among same frequency | DLL (tail = LRU) | O(1) |
| Find minimum frequency | `minFreq` variable | O(1) |
| Update frequency | Remove from old DLL, add to new DLL | O(1) |

---

## Algorithm

### `Get(key)`:

```
1. If key not in keyToNode → return -1
2. Get the node from keyToNode
3. Update frequency:
   a. Remove node from current frequency's DLL
   b. If current freq == minFreq AND DLL is now empty → minFreq++
   c. Increment node's frequency
   d. Add node to new frequency's DLL (create if needed)
4. Return node's value
```

### `Put(key, value)`:

```
1. If capacity == 0 → return
2. If key exists in keyToNode:
   a. Update the value
   b. Call updateFrequency (same as Get)
   c. Return
3. If size == capacity (cache full):
   a. Get DLL for minFreq
   b. Remove the last node (LRU) from that DLL
   c. Delete from keyToNode
   d. Decrement size
4. Create new node with freq=1
5. Add to keyToNode
6. Add to freqToDLL[1] (create if needed)
7. Set minFreq = 1 (new node always has freq 1)
8. Increment size
```

### `updateFrequency(node)`:

```
1. Get old frequency
2. Remove node from old frequency's DLL
3. If oldFreq == minFreq AND old DLL is empty:
   → minFreq++ (no more nodes at this frequency)
4. Increment node's frequency
5. Add node to new frequency's DLL (create if needed)
```

---

## Time & Space Complexity

| Operation | Time Complexity | Space Complexity |
|-----------|-----------------|------------------|
| `Get()` | O(1) | O(1) |
| `Put()` | O(1) | O(1) |
| **Overall Space** | - | O(capacity) |

### Why O(1)?

- **HashMap lookups**: O(1) average
- **DLL operations**: Add to head O(1), Remove node O(1), Remove tail O(1)
- **minFreq update**: Simple increment when DLL becomes empty

---

## Common Bugs and Edge Cases

### Bug 1: Wrong Condition for Incrementing minFreq

```go
// ❌ WRONG - checking overall cache size
if oldFreq == lfuCache.minFreq && lfuCache.size == 0 {
    lfuCache.minFreq++
}

// ✅ CORRECT - checking if old frequency's DLL is empty
if oldFreq == lfuCache.minFreq && oldDLL.size == 0 {
    lfuCache.minFreq++
}
```

### Bug 2: Forgetting to Reset minFreq on New Insert

```go
// When adding a new node, always set minFreq to 1
// because new nodes have frequency 1
lfuCache.minFreq = 1
```

### Bug 3: Not Handling Zero Capacity

```go
func (lfuCache *LFUCache) Put(key, value int) {
    if lfuCache.capacity == 0 {
        return  // Handle edge case
    }
    // ...
}
```

### Bug 4: Incorrect minFreq Update After Eviction

```go
// ❌ WRONG - incrementing minFreq after eviction
if lfuCache.size == lfuCache.capacity {
    // ... evict node
    lfuCache.minFreq++  // Wrong! New node will have freq=1
}

// ✅ CORRECT - minFreq will be set to 1 when new node is added
```

---

## Key Interview Points

### 1. **Design Decisions**

- **Why use HashMap of DLLs instead of a single DLL?**
  - Need O(1) access to all nodes of a particular frequency
  - Single DLL would require O(n) to find all nodes with minFreq

- **Why track minFreq separately?**
  - Avoid O(n) search through all frequencies to find minimum
  - Only updates in two cases: DLL becomes empty, or new node added

- **Why store key in Node?**
  - When evicting, need to delete from `keyToNode` HashMap
  - Without key in node, would need O(n) search

### 2. **Comparison with LRU**

| Aspect | LRU | LFU |
|--------|-----|-----|
| Data Structures | 1 HashMap + 1 DLL | 2 HashMaps + Multiple DLLs |
| Complexity | Simpler | More complex |
| Eviction | Oldest access | Lowest frequency (then oldest) |
| Use Case | Temporal locality | Frequency-based popularity |

### 3. **Production Considerations**

| Aspect | Consideration |
|--------|---------------|
| **Thread Safety** | Need mutex/lock for concurrent access |
| **Memory Overhead** | More memory than LRU (freq tracking) |
| **Cache Pollution** | One-time access items stay if freq is low |
| **Frequency Decay** | May need aging mechanism for long-running systems |

### 4. **Follow-up Questions**

1. **"How would you implement frequency decay?"**
   - Periodically halve all frequencies
   - Use time-weighted frequency

2. **"What if you need O(1) for finding k-th most frequent?"**
   - Use additional data structures (e.g., indexed priority queue)

3. **"LFU vs LRU - when to use which?"**
   - LFU: When popular items should stay (e.g., CDN, trending content)
   - LRU: When recent items are more relevant (e.g., browser cache)

4. **"How to handle frequency overflow?"**
   - Use modular arithmetic or periodic reset

### 5. **Real-World Applications**

- **CDN (Content Delivery Network)**: Cache popular content
- **Database Query Cache**: Keep frequently accessed queries
- **DNS Caching**: Popular domain lookups
- **CPU Cache**: Frequency-based replacement policies
- **Recommendation Systems**: Track popular items

---

## Code Implementation (Go)

```go
package main

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

func Constructor(capacity int) LFUCache {
    return LFUCache{
        capacity:  capacity,
        keyToNode: make(map[int]*Node),
        freqToDLL: make(map[int]*DLL),
    }
}

func NewDLL() *DLL {
    head := &Node{}
    tail := &Node{}
    head.next = tail
    tail.prev = head
    return &DLL{size: 0, head: head, tail: tail}
}

func (dll *DLL) addNode(node *Node) {
    node.next = dll.head.next
    node.prev = dll.head
    dll.head.next.prev = node
    dll.head.next = node
    dll.size++
}

func (dll *DLL) removeNode(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
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

func (lfu *LFUCache) updateFrequency(node *Node) {
    oldFreq := node.freq
    oldDLL := lfu.freqToDLL[oldFreq]
    oldDLL.removeNode(node)

    // Update minFreq if we emptied the minimum frequency list
    if oldFreq == lfu.minFreq && oldDLL.size == 0 {
        lfu.minFreq++
    }

    node.freq++

    newDLL, exists := lfu.freqToDLL[node.freq]
    if !exists {
        newDLL = NewDLL()
        lfu.freqToDLL[node.freq] = newDLL
    }
    newDLL.addNode(node)
}

func (lfu *LFUCache) Get(key int) int {
    node, exists := lfu.keyToNode[key]
    if !exists {
        return -1
    }
    lfu.updateFrequency(node)
    return node.value
}

func (lfu *LFUCache) Put(key, value int) {
    if lfu.capacity == 0 {
        return
    }

    // Key exists - update value and frequency
    if node, exists := lfu.keyToNode[key]; exists {
        node.value = value
        lfu.updateFrequency(node)
        return
    }

    // Cache is full - evict LFU (and LRU among ties)
    if lfu.size == lfu.capacity {
        dll := lfu.freqToDLL[lfu.minFreq]
        nodeToRemove := dll.removeLastNode()
        delete(lfu.keyToNode, nodeToRemove.key)
        lfu.size--
    }

    // Add new node with freq = 1
    newNode := &Node{key: key, value: value, freq: 1}
    lfu.keyToNode[key] = newNode

    dll, exists := lfu.freqToDLL[1]
    if !exists {
        dll = NewDLL()
        lfu.freqToDLL[1] = dll
    }
    dll.addNode(newNode)

    lfu.minFreq = 1  // New node always has freq 1
    lfu.size++
}
```

---

## Complexity Summary Table

| Component | Purpose | Time | Space |
|-----------|---------|------|-------|
| `keyToNode` | Key → Node lookup | O(1) | O(n) |
| `freqToDLL` | Freq → DLL lookup | O(1) | O(n) |
| `DLL.addNode` | Add to front | O(1) | O(1) |
| `DLL.removeNode` | Remove specific node | O(1) | O(1) |
| `DLL.removeLastNode` | Remove LRU | O(1) | O(1) |
| `minFreq` | Track minimum freq | O(1) | O(1) |

---

## LeetCode Reference

- **Problem**: [460. LFU Cache](https://leetcode.com/problems/lfu-cache/)
- **Difficulty**: Hard
- **Tags**: Hash Table, Linked List, Design, Doubly-Linked List

