# Got it — you’re asking for **DSA/coding round questions** that are most commonly asked for a **Senior Golang Developer (5 years experience)**, not Go theory questions.

At this experience level, companies usually ask:

* medium-to-hard LeetCode style problems
* concurrency-oriented coding
* backend-centric data structure problems
* optimization-focused implementations
* machine coding / low-level design

Here are the 10 most frequently asked coding problems.

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

# 2. Producer Consumer using Goroutines & Channels

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

# 3. Merge Intervals

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

# 4. Top K Frequent Elements

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

# 5. Design a Rate Limiter

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

# 6. Implement Worker Pool

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

# 7. Longest Substring Without Repeating Characters

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

# 8. Detect Cycle in Linked List

### Expected Solution

Floyd’s Cycle Detection Algorithm

### Concepts

* Fast & slow pointers
* Space optimization

### Follow-up

* Find cycle starting node

---

# 9. Concurrent Safe In-Memory Key Value Store

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

# 10. Design URL Shortener

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
2. Worker Pool
3. Producer Consumer
4. Rate Limiter
5. Top K Frequent
6. Merge Intervals
7. Sliding Window
8. URL Shortener
9. Concurrent Map
10. Heap-based problems
