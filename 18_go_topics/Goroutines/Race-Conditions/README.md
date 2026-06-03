# Race Conditions in Go

> A deep dive into Race Conditions, Data Races, Go Memory Model, Happens-Before Relationships, Synchronization Primitives, Race Detection, Production Failures, and Google L5 Interview Preparation.

---

# Table of Contents

* [1. Introduction](#1-introduction)
* [2. What is a Race Condition?](#2-what-is-a-race-condition)
* [3. Data Race vs Race Condition](#3-data-race-vs-race-condition)
* [4. Why Race Conditions Occur](#4-why-race-conditions-occur)
* [5. Read-Modify-Write Problem](#5-read-modify-write-problem)
* [6. Race Condition Example](#6-race-condition-example)
* [7. Internal CPU Operations](#7-internal-cpu-operations)
* [8. Scheduler and Race Conditions](#8-scheduler-and-race-conditions)
* [9. Why Data Races Are Dangerous](#9-why-data-races-are-dangerous)
* [10. Go Memory Model](#10-go-memory-model)
* [11. Memory Visibility Problems](#11-memory-visibility-problems)
* [12. Happens-Before Relationship](#12-happens-before-relationship)
* [13. Synchronization Mechanisms](#13-synchronization-mechanisms)
* [14. Mutexes](#14-mutexes)
* [15. RWMutex](#15-rwmutex)
* [16. Channels and Synchronization](#16-channels-and-synchronization)
* [17. Atomic Operations](#17-atomic-operations)
* [18. WaitGroup Synchronization](#18-waitgroup-synchronization)
* [19. Concurrent Map Access](#19-concurrent-map-access)
* [20. Race Detector Internals](#20-race-detector-internals)
* [21. Using go test -race](#21-using-go-test--race)
* [22. Production Failures Caused By Races](#22-production-failures-caused-by-races)
* [23. Common Hidden Traps](#23-common-hidden-traps)
* [24. Performance Tradeoffs](#24-performance-tradeoffs)
* [25. Whiteboard Reasoning](#25-whiteboard-reasoning)
* [26. Production-Grade Examples](#26-production-grade-examples)
* [27. Google L5 Interview Questions](#27-google-l5-interview-questions)
* [28. Key Takeaways](#28-key-takeaways)

---

# 1. Introduction

Race conditions are one of the most common causes of production bugs in concurrent systems.

Unlike deadlocks, race conditions usually:

* do not crash immediately
* are difficult to reproduce
* depend on timing
* may appear only under load

These characteristics make race conditions among the most dangerous classes of concurrency bugs.

For senior-level interviews, understanding race conditions requires knowledge of:

* scheduler behavior
* memory visibility
* synchronization primitives
* CPU instruction ordering
* Go memory model

---

# 2. What is a Race Condition?

A race condition occurs when:

> Program correctness depends on the relative timing or ordering of concurrent operations.

The key phrase is:

```text
Correctness depends on execution order.
```

If different execution orders produce different outcomes, a race condition exists.

---

# 3. Data Race vs Race Condition

Many engineers use these terms interchangeably.

This is incorrect.

---

## Data Race

A data race occurs when:

```text
Two goroutines access same memory
AND
At least one access is a write
AND
No synchronization exists
```

Example:

```go
var counter int

go func() {
	counter++
}()

go func() {
	counter++
}()
```

This is a data race.

---

## Race Condition

Race condition is broader.

Example:

```text
G1 Sends Welcome Email
G2 Deletes Account
```

Depending on execution order:

```text
Email before delete -> OK
Delete before email -> Wrong
```

No shared memory exists.

No data race exists.

Race condition still exists.

---

## Important Interview Statement

```text
Every Data Race is a Race Condition.

Not Every Race Condition is a Data Race.
```

---

# 4. Why Race Conditions Occur

Concurrency introduces:

```text
Non-deterministic execution order
```

Meaning:

```text
Run #1
G1 → G2

Run #2
G2 → G1

Run #3
G2 → G1 → G2
```

Different schedules produce different outcomes.

---

# 5. Read-Modify-Write Problem

Most race conditions involve:

```text
Read
Modify
Write
```

Example:

```go
counter++
```

Many developers think:

```text
Single operation
```

Reality:

```text
LOAD counter
ADD 1
STORE counter
```

Three separate operations.

---

# 6. Race Condition Example

Initial value:

```text
counter = 5
```

Two goroutines execute:

```go
counter++
```

Timeline:

```text
G1 Load 5

G2 Load 5

G1 Store 6

G2 Store 6
```

Expected:

```text
7
```

Actual:

```text
6
```

Lost update occurred.

---

# 7. Internal CPU Operations

At CPU level:

```text
counter++
```

expands into:

```text
MOV counter -> register
ADD 1
MOV register -> counter
```

These instructions are interruptible.

The scheduler may switch goroutines between any of these steps.

---

# 8. Scheduler and Race Conditions

Go scheduler controls:

```text
Which goroutine runs
When it runs
How long it runs
```

Because scheduling is non-deterministic:

```text
Program behavior becomes timing dependent.
```

Timing-dependent correctness creates races.

---

# 9. Why Data Races Are Dangerous

Possible outcomes:

```text
Correct Result
Incorrect Result
Corrupted State
Stale Read
Unexpected Behavior
```

The Go Memory Model treats programs with data races as:

```text
Undefined Behavior
```

You lose correctness guarantees.

---

# 10. Go Memory Model

The memory model defines:

> When writes performed by one goroutine become visible to another goroutine.

Without synchronization:

```text
Visibility is not guaranteed.
```

---

# 11. Memory Visibility Problems

Example:

```go
var ready bool

go func() {
	ready = true
}()

for !ready {
}
```

This program is broken.

Reason:

```text
No synchronization
```

Reader may never observe:

```text
ready = true
```

---

# 12. Happens-Before Relationship

Most important concept in race-condition discussions.

Definition:

> A happens-before relationship guarantees visibility and ordering between goroutines.

If:

```text
Operation A happens-before Operation B
```

Then:

```text
B sees effects of A
```

---

## Happens-Before Created By

### Mutex

```go
mu.Unlock()
mu.Lock()
```

---

### Channels

```go
ch <- value
value := <-ch
```

---

### Atomics

```go
atomic.Store()
atomic.Load()
```

---

### WaitGroups

```go
wg.Done()
wg.Wait()
```

---

# 13. Synchronization Mechanisms

Go provides:

```text
Mutex
RWMutex
Channels
Atomics
WaitGroups
Cond Variables
```

All create ordering guarantees.

---

# 14. Mutexes

Most common race prevention tool.

```go
var mu sync.Mutex

mu.Lock()
counter++
mu.Unlock()
```

Guarantees:

```text
Mutual Exclusion
Memory Visibility
Ordering
```

---

# 15. RWMutex

Allows:

```text
Multiple Readers
Single Writer
```

Useful for read-heavy workloads.

---

# 16. Channels and Synchronization

Channels do more than communication.

They also create happens-before relationships.

Example:

```go
done := make(chan struct{})

go func() {
	value = 42
	close(done)
}()

<-done
fmt.Println(value)
```

Safe.

---

# 17. Atomic Operations

Example:

```go
atomic.AddInt64(&counter, 1)
```

Benefits:

```text
Lock-free
Fast
Thread-safe
```

Limitations:

```text
Works best on simple values
```

Not a replacement for mutexes.

---

# 18. WaitGroup Synchronization

Example:

```go
wg.Done()
wg.Wait()
```

Guarantees:

```text
All goroutine writes complete
before Wait returns
```

---

# 19. Concurrent Map Access

Maps are not thread-safe.

Unsafe:

```go
go func() {
	m["a"] = 1
}()

go func() {
	fmt.Println(m["a"])
}()
```

May cause:

```text
Race
Corruption
Panic
```

---

# 20. Race Detector Internals

Go race detector instruments:

```text
Reads
Writes
Synchronization Events
```

It builds a happens-before graph.

If conflicting accesses occur without synchronization:

```text
WARNING: DATA RACE
```

is reported.

---

# 21. Using go test -race

```bash
go test -race ./...
```

or

```bash
go run -race main.go
```

The race detector should always be part of CI pipelines.

---

# 22. Production Failures Caused By Races

Examples:

* Incorrect account balances
* Double billing
* Duplicate order processing
* Corrupted caches
* Lost updates
* Incorrect metrics
* Invalid user states

---

# 23. Common Hidden Traps

## Trap 1

Reading without locking.

```go
if counter > 0 {
}
```

Still a race.

---

## Trap 2

Concurrent map access.

---

## Trap 3

Assuming atomics protect complex objects.

---

## Trap 4

Using race detector as the only defense.

---

## Trap 5

Double-checked locking mistakes.

---

# 24. Performance Tradeoffs

| Mechanism | Performance         | Complexity | Safety |
| --------- | ------------------- | ---------- | ------ |
| Mutex     | Medium              | Low        | High   |
| RWMutex   | High Read Workloads | Medium     | High   |
| Channel   | Medium              | Medium     | High   |
| Atomic    | Highest             | High       | Medium |

---

# 25. Whiteboard Reasoning

When debugging a race:

Step 1:

```text
Identify Shared State
```

Step 2:

```text
Identify Readers
```

Step 3:

```text
Identify Writers
```

Step 4:

```text
Identify Synchronization
```

Step 5:

```text
Verify Happens-Before Relationship
```

If missing:

```text
Race Exists
```

---

# 26. Production-Grade Examples

## Mutex Protected Counter

```go
type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.n++
}
```

---

## Atomic Counter

```go
var counter int64

atomic.AddInt64(&counter, 1)
```

---

## Safe Map

```go
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}
```

---

# 27. Google L5 Interview Questions

## Q1. What is a race condition?

### Answer

A race condition occurs when program correctness depends on the relative timing or ordering of concurrent operations. Different execution schedules may produce different outcomes.

---

## Q2. What is a data race?

### Answer

A data race occurs when two goroutines access the same memory location concurrently, at least one access is a write, and no synchronization mechanism establishes a happens-before relationship.

---

## Q3. Difference between race condition and data race?

### Answer

A data race is a specific memory-access problem involving unsynchronized shared state.

A race condition is broader and includes any correctness issue caused by timing or ordering dependencies.

Every data race is a race condition, but not every race condition is a data race.

---

## Q4. Why is counter++ not atomic?

### Answer

Because it consists of multiple CPU operations:

```text
Load
Add
Store
```

Context switches may occur between these operations, causing lost updates.

---

## Q5. What is a happens-before relationship?

### Answer

A happens-before relationship guarantees ordering and memory visibility between goroutines. If operation A happens-before operation B, then B must observe all effects of A.

---

## Q6. How do channels prevent races?

### Answer

Channel send and receive operations create happens-before relationships. This guarantees visibility of writes performed before the send.

---

## Q7. Why can a race-free program still be incorrect?

### Answer

Because race detector only finds shared-memory races.

Example:

```text
Delete Account
Send Email
```

Protected by locks.

No data race.

Wrong business ordering.

Still incorrect.

---

## Q8. How does Go race detector work?

### Answer

It instruments memory accesses and synchronization operations, builds a happens-before graph, and detects conflicting accesses lacking synchronization.

---

## Q9. Why are atomics faster than mutexes?

### Answer

Atomics use CPU-level synchronization instructions and avoid lock acquisition, blocking, and wakeup overhead.

---

## Q10. When would you choose atomics over mutexes?

### Answer

For simple independent values:

* counters
* flags
* sequence numbers

Not for protecting complex object graphs.

---

## Q11. Can concurrent reads race?

### Answer

No.

Multiple reads are safe.

A data race requires at least one write.

---

## Q12. Why are Go maps not thread-safe?

### Answer

Map internals may resize buckets, relocate entries, and modify metadata. Concurrent modification can corrupt internal structures.

---

## Q13. What production metrics might indicate a race condition?

### Answer

* Non-deterministic failures
* Inconsistent counts
* Duplicate records
* Sporadic crashes
* Flaky tests
* Corrupted state

---

## Q14. What is memory visibility?

### Answer

Memory visibility refers to whether updates performed by one goroutine become observable by another goroutine.

---

## Q15. What is the most important concept behind race prevention?

### Answer

Establishing happens-before relationships through synchronization primitives.

---

# 28. Key Takeaways

* Race conditions are timing-dependent correctness failures.
* Data races are unsynchronized shared-memory accesses.
* Go programs require synchronization for visibility guarantees.
* Happens-before relationships are the foundation of concurrency correctness.
* Mutexes, channels, atomics, and WaitGroups create ordering guarantees.
* The race detector finds data races but not all race conditions.
* Concurrent map access is unsafe without protection.
* Atomics are powerful but limited.
* Correctness comes before performance.
* Senior engineers reason about ordering, visibility, and synchronization—not just locking.

---
