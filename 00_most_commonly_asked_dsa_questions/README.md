# **Most commonly asked Golang DSA questions**.

At this experience level, companies usually ask:

* medium-to-hard LeetCode style problems
* concurrency-oriented coding
* backend-centric data structure problems
* optimization-focused implementations
* machine coding / low-level design

Here are the 11 most frequently asked coding problems.

---

## Table of Contents

| #  | Problem                                                                                             | Key Concepts                                            | Solution |
|----|-----------------------------------------------------------------------------------------------------|---------------------------------------------------------|----------|
| 1  | [LRU Cache](#1-lru-cache)                                                                           | HashMap, Doubly Linked List, O(1) Design                | ✅ [Solution](./2.LRU-Cache/) |
| 2  | [LFU Cache](#2-lfu-cache)                                                                           | HashMap, Multiple DLLs, Frequency Tracking, O(1) Design | ✅ [Solution](./1.LFU-Cache/) |
| 3  | [Producer Consumer using Goroutines & Channels](#3-producer-consumer-using-goroutines--channels)    | Concurrency, Channels, WaitGroups                       | ⬜ |
| 4  | [Merge Intervals](#4-merge-intervals)                                                               | Sorting, Array Manipulation                             | ⬜ |
| 5  | [Top K Frequent Elements](#5-top-k-frequent-elements)                                               | Heap, HashMap, Optimization                             | ⬜ |
| 6  | [Design a Rate Limiter](#6-design-a-rate-limiter)                                                   | Token Bucket, Concurrency Safety                        | ⬜ |
| 7  | [Implement Worker Pool](#7-implement-worker-pool)                                                   | Goroutines, Channels, Synchronization                   | ⬜ |
| 8  | [Longest Substring Without Repeating Characters](#8-longest-substring-without-repeating-characters) | Sliding Window, Two Pointers                            | ⬜ |
| 9  | [Detect Cycle in Linked List](#9-detect-cycle-in-linked-list)                                       | Floyd's Algorithm, Fast & Slow Pointers                 | ⬜ |
| 10 | [Concurrent Safe In-Memory Key Value Store](#10-concurrent-safe-in-memory-key-value-store)          | Mutex, RWMutex, Map Internals                           | ⬜ |
| 11 | [Design URL Shortener](#11-design-url-shortener)                                                    | Base62 Encoding, Scalability                            | ⬜ |

**Other Sections:**
- [Most Frequently Asked Algorithms for Senior Go Roles](#most-frequently-asked-algorithms-for-senior-go-roles)
- [What Companies Usually Ask](#what-companies-usually-ask)
- [Most Important Preparation Areas](#most-important-preparation-areas)
- [Best Practice for Preparation](#best-practice-for-preparation)
- [Highest ROI Problems](#highest-roi-problems)

---

# 1. LRU Cache

### Why it is asked

Tests:

* HashMap + Doubly Linked List
* O(1) design
* Real backend cache understanding

### Expected Complexity

* `get()` → O(1)
* `put()` → O(1)

### Common Follow-up

* Thread-safe LRU cache in Go
* Distributed cache discussion

### Core Concepts

* Map
* Linked List
* Pointer manipulation

---

# 2. LFU Cache

### Why it is asked

Tests:

* Advanced cache eviction strategy
* Two HashMaps + Multiple Doubly Linked Lists
* O(1) design with frequency tracking
* Understanding of LFU vs LRU trade-offs

### Expected Complexity

* `get()` → O(1)
* `put()` → O(1)

### Data Structures Required

* `keyToNode` → HashMap for O(1) key lookup
* `freqToDLL` → HashMap of DLLs for each frequency
* `minFreq` → Track minimum frequency for O(1) eviction

### Key Implementation Details

* When frequency ties occur, evict LRU among same frequency
* Update `minFreq` only when:
  - Old frequency's DLL becomes empty
  - New node is added (reset to 1)

### Common Follow-up

* LFU vs LRU: When to use which?
* Thread-safe LFU cache
* Frequency decay for long-running systems

### Core Concepts

* Multiple Maps
* Multiple Linked Lists
* Frequency tracking
* Tie-breaking with recency

---

# 3. Producer Consumer using Goroutines & Channels

### Why it is asked

Extremely common for Go interviews.

### What interviewers evaluate

* Concurrency handling
* Channel usage
* Deadlock prevention
* Graceful shutdown

### Expected Features

* Multiple workers
* Buffered channels
* Context cancellation
* WaitGroups

### Senior-Level Extension

* Rate limiting
* Retry queue
* Worker pool

---

# 4. Merge Intervals

### Example

```text
[[1,3],[2,6],[8,10],[15,18]]
→ [[1,6],[8,10],[15,18]]
```

### Why it is asked

Tests:

* Sorting
* Array manipulation
* Edge case handling

### Complexity

* O(n log n)

---

# 5. Top K Frequent Elements

### Example

```text
Input: [1,1,1,2,2,3], k=2
Output: [1,2]
```

### Concepts Tested

* Heap/Priority Queue
* HashMap
* Optimization

### Senior Follow-up

* Streaming version
* Huge dataset version

---

# 6. Design a Rate Limiter

This is VERY common for senior backend Go roles.

### Variants

* Token Bucket
* Leaky Bucket
* Sliding Window

### Interview Focus

* Concurrency safety
* Scalability
* Time complexity
* Production use cases

### Senior-Level Discussion

* Redis distributed rate limiter
* API gateway integration

---

# 7. Implement Worker Pool

### Typical Problem

Process jobs concurrently with fixed workers.

### Concepts Tested

* Goroutines
* Channels
* Synchronization
* Backpressure

### Expected Additions

* Error handling
* Cancellation
* Panic recovery

---

# 8. Longest Substring Without Repeating Characters

### Example

```text
"abcabcbb" → 3
```

### Concepts

* Sliding Window
* HashMap
* Two pointers

### Why it matters

Tests optimization thinking.

### Expected Complexity

* O(n)

---

# 9. Detect Cycle in Linked List

### Expected Solution

Floyd's Cycle Detection Algorithm

### Concepts

* Fast & slow pointers
* Space optimization

### Follow-up

* Find cycle starting node

---

# 10. Concurrent Safe In-Memory Key Value Store

Very common in Go backend interviews.

### Features expected

```text
PUT(key, value)
GET(key)
DELETE(key)
```

### What is tested

* Mutex/RWMutex
* Map internals
* Concurrency safety

### Senior Extensions

* TTL expiration
* Sharding
* Persistence
* LRU eviction

---

# 11. Design URL Shortener

Classic senior backend coding/design round.

### Components

* Unique ID generation
* Base62 encoding
* Database schema
* Cache layer

### What interviewers evaluate

* Scalability
* Collision handling
* Distributed systems thinking

### Senior Follow-up

* 100M URLs/day scaling
* Analytics tracking
* Hot key handling

---

# Most Frequently Asked Algorithms for Senior Go Roles

## Arrays & Strings

* Sliding Window
* Two Pointers
* Prefix Sum

## Concurrency

* Worker Pool
* Producer Consumer
* Rate Limiter

## Design Problems

* LRU Cache
* KV Store
* URL Shortener

## Heap Problems

* Top K elements
* Merge K sorted lists

## Graph Problems

* BFS/DFS
* Dependency resolution
* Shortest path

---

# What Companies Usually Ask

## Product Companies

Examples:

* Google
* Uber
* Amazon
* Flipkart

Focus:

* DSA hard problems
* Concurrency
* System design

---

## FinTech / High Scale Backend

Examples:

* Razorpay
* PhonePe
* CRED

Focus:

* Rate limiter
* Caching
* Distributed locking
* Goroutine safety

---

## Startups

Focus:

* Practical backend coding
* APIs
* DB optimization
* Concurrent processing

---

# Most Important Preparation Areas

For 5-year Go developers, prioritize:

1. Concurrency coding
2. LLD coding rounds
3. Cache design
4. Goroutine/channel problems
5. Heap + sliding window
6. Backend-focused DSA
7. Thread-safe implementations

---

# Best Practice for Preparation

## Must Solve At Least

* 15 sliding window problems
* 10 heap problems
* 10 concurrency problems
* 5 cache/design problems
* 5 machine coding problems

---

# Highest ROI Problems

If time is limited, master these first:

1. LRU Cache
2. LFU Cache
3. Worker Pool
4. Producer Consumer
5. Rate Limiter
6. Top K Frequent
7. Merge Intervals
8. Sliding Window
9. URL Shortener
10. Concurrent Map
11. Heap-based problems
