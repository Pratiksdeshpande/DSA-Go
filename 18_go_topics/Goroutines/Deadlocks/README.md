# Deadlocks in Go (Golang)

> A deep dive into Deadlocks in Go covering runtime internals, scheduler interaction, channels, mutexes, WaitGroups, production debugging, prevention strategies, and Google L5 interview preparation.

---

# Table of Contents

* [1. Introduction](#1-introduction)
* [2. What is a Deadlock?](#2-what-is-a-deadlock)
* [3. Why Deadlocks Matter](#3-why-deadlocks-matter)
* [4. Coffman Conditions](#4-coffman-conditions)
* [5. Types of Deadlocks in Go](#5-types-of-deadlocks-in-go)
* [6. Channel Deadlocks](#6-channel-deadlocks)

  * [6.1 Receive Without Sender](#61-receive-without-sender)
  * [6.2 Send Without Receiver](#62-send-without-receiver)
  * [6.3 Buffered Channel Deadlock](#63-buffered-channel-deadlock)
  * [6.4 Range on Never Closed Channel](#64-range-on-never-closed-channel)
* [7. WaitGroup Deadlocks](#7-waitgroup-deadlocks)
* [8. Mutex Deadlocks](#8-mutex-deadlocks)

  * [8.1 Self Deadlock](#81-self-deadlock)
  * [8.2 Circular Deadlock](#82-circular-deadlock)
* [9. Runtime Deadlock Detection](#9-runtime-deadlock-detection)
* [10. Scheduler Interaction](#10-scheduler-interaction)
* [11. Internal Runtime Flow](#11-internal-runtime-flow)
* [12. Distributed Deadlocks](#12-distributed-deadlocks)
* [13. Deadlock vs Livelock](#13-deadlock-vs-livelock)
* [14. Deadlock Prevention Techniques](#14-deadlock-prevention-techniques)
* [15. Production Debugging](#15-production-debugging)
* [16. Performance Considerations](#16-performance-considerations)
* [17. Common Interview Questions](#17-common-interview-questions)
* [18. Google L5 Follow-up Questions](#18-google-l5-follow-up-questions)
* [19. Interview Cheat Sheet](#19-interview-cheat-sheet)
* [20. Key Takeaways](#20-key-takeaways)

---

# 1. Introduction

Deadlocks are one of the most important concurrency correctness failures.

Unlike race conditions, which may produce incorrect results, deadlocks completely stop progress.

A deadlocked system may:

* stop processing requests
* stop consuming messages
* stop accepting traffic
* become permanently blocked

Deadlocks are frequently tested in senior Go interviews because they reveal understanding of:

* channels
* goroutines
* mutexes
* scheduler behavior
* runtime internals

---

# 2. What is a Deadlock?

A deadlock occurs when one or more goroutines are waiting forever for events that can never occur.

Formal definition:

> A set of goroutines are blocked forever because each is waiting for another goroutine to perform an action that will never happen.

---

## Simple Visualization

```text
Goroutine A
Waiting for B

Goroutine B
Waiting for A
```

Neither can proceed.

System progress becomes impossible.

---

# 3. Why Deadlocks Matter

Deadlocks cause:

* Infinite request latency
* Hung services
* Failed deployments
* Resource starvation
* Distributed system outages

Production impact:

```text
Requests arrive
      ↓
Workers blocked
      ↓
Queue grows
      ↓
Timeouts increase
      ↓
Outage
```

---

# 4. Coffman Conditions

A deadlock requires all four conditions.

---

## 1. Mutual Exclusion

Resource can only be held by one goroutine.

Example:

```go
var mu sync.Mutex
```

Only one goroutine owns the lock.

---

## 2. Hold and Wait

A goroutine holds one resource while waiting for another.

```text
G1:
Holding Lock A
Waiting Lock B
```

---

## 3. No Preemption

Resources cannot be forcibly taken.

```go
mu.Lock()
```

Another goroutine cannot steal it.

---

## 4. Circular Wait

Circular dependency exists.

```text
G1 -> waiting for G2

G2 -> waiting for G3

G3 -> waiting for G1
```

---

## Important Interview Point

Breaking ANY one Coffman condition prevents deadlocks.

Most production systems break:

* Circular Wait
* Hold and Wait

---

# 5. Types of Deadlocks in Go

```text
Deadlocks
│
├── Channel Deadlocks
│
├── WaitGroup Deadlocks
│
├── Mutex Deadlocks
│
└── Distributed Deadlocks
```

---

# 6. Channel Deadlocks

Most common category.

---

# 6.1 Receive Without Sender

```go
func main() {
	ch := make(chan int)

	<-ch
}
```

Output:

```text
fatal error:
all goroutines are asleep - deadlock!
```

---

## Runtime Flow

```text
Receive
   ↓
Buffer Empty?
   ↓ YES
Sender Waiting?
   ↓ NO
Park Goroutine
```

---

## Internal State

```text
Main Goroutine
      |
      v
Waiting on Receive

Channel Empty

No Sender Exists
```

Deadlock.

---

# 6.2 Send Without Receiver

```go
func main() {
	ch := make(chan int)

	ch <- 42
}
```

---

Runtime:

```text
Send
 ↓
Receiver Exists?
 ↓ NO
Park Goroutine
```

Deadlock.

---

# 6.3 Buffered Channel Deadlock

```go
func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	ch <- 3
}
```

---

Visualization

```text
Buffer

[1][2]

Capacity = 2

Third Send
    ↓
Blocks Forever
```

---

Important:

Buffered channels delay blocking.

They do NOT eliminate deadlocks.

---

# 6.4 Range on Never Closed Channel

```go
func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
```

---

Output

```text
1

fatal error:
all goroutines are asleep - deadlock!
```

---

Why?

```text
range waits for:

1. new value
OR
2. channel close

Neither happens
```

Receiver waits forever.

---

# 7. WaitGroup Deadlocks

Extremely common.

---

## Missing Done()

```go
var wg sync.WaitGroup

wg.Add(1)

go func() {
	return
}()

wg.Wait()
```

---

Visualization

```text
Counter = 1

Wait()
  ↓

Waiting Forever
```

---

Correct

```go
go func() {
	defer wg.Done()
}()
```

---

## Add/Done Mismatch

```go
wg.Add(3)

wg.Done()
wg.Done()

wg.Wait()
```

Deadlock.

Counter never reaches zero.

---

# 8. Mutex Deadlocks

---

# 8.1 Self Deadlock

```go
var mu sync.Mutex

mu.Lock()
mu.Lock()
```

---

Why?

Go mutexes are NOT reentrant.

Visualization:

```text
G1 owns mutex
      ↓
G1 requests mutex
      ↓
Wait Forever
```

---

# 8.2 Circular Deadlock

```go
var muA sync.Mutex
var muB sync.Mutex
```

---

Goroutine 1

```go
muA.Lock()
muB.Lock()
```

---

Goroutine 2

```go
muB.Lock()
muA.Lock()
```

---

Possible Execution

```text
G1 acquires A

G2 acquires B

G1 waits B

G2 waits A
```

Deadlock.

---

Visualization

```text
       waits
G1 ------------> B

^               |
|               |
|               v

A <------------ G2
       waits
```

---

# 9. Runtime Deadlock Detection

Go runtime can detect global deadlocks.

---

Condition:

```text
No runnable goroutines

AND

No future progress possible
```

---

Runtime Output

```text
fatal error:
all goroutines are asleep - deadlock!
```

---

# Runtime Detection Flow

```text
Scheduler
    ↓
Runnable Goroutines?
    ↓ NO
Timers Pending?
    ↓ NO
Network Events Pending?
    ↓ NO
Deadlock Panic
```

---

# 10. Scheduler Interaction

Blocked goroutines move into waiting state.

---

State Transition

```text
_Grunnable
      ↓
channel receive

_Gwaiting
```

---

Scheduler Run Queue

```text
Before

P Run Queue
------------
G1
G2
G3

After Blocking

P Run Queue
------------
(empty)
```

---

When every goroutine becomes:

```text
_Gwaiting
```

Scheduler detects deadlock.

---

# 11. Internal Runtime Flow

Example:

```go
<-ch
```

Runtime executes:

```text
chanrecv()
```

---

Simplified Internal Logic

```text
Lock Channel
      ↓
Check Buffer
      ↓
Check Send Queue
      ↓
No Sender Found
      ↓
Park Goroutine
      ↓
_Gwaiting
```

---

Later:

```text
Sender Arrives
     ↓
Wake Receiver
     ↓
_Grunnable
```

If sender never arrives:

Deadlock.

---

# 12. Distributed Deadlocks

Most dangerous production deadlocks.

---

Example

```text
Service A
Locks Resource X
       ↓
Calls Service B

Service B
Calls Service A

Service A
Needs Resource X
```

---

Visualization

```text
Service A
   ↓
Waiting B

Service B
   ↓
Waiting A
```

Distributed deadlock.

Runtime cannot detect this.

---

# 13. Deadlock vs Livelock

Many interviewers ask this.

---

## Deadlock

Nobody moves.

```text
A waits B

B waits A
```

Progress = 0

---

## Livelock

Everyone moves.

Nobody progresses.

Example:

```text
Person A steps left

Person B steps left

Person A steps right

Person B steps right
```

Both active.

No progress.

---

# 14. Deadlock Prevention Techniques

---

## Lock Ordering

Always acquire locks in same order.

Correct:

```text
A → B → C
```

Never:

```text
B → A
```

---

## Timeouts

```go
select {
case v := <-ch:
	return v

case <-time.After(time.Second):
	return timeout
}
```

---

## Context Cancellation

```go
ctx, cancel := context.WithTimeout(...)
```

Prevents permanent waiting.

---

## Avoid Nested Locks

Bad:

```go
A.Lock()
B.Lock()
```

Prefer:

```go
Acquire
Use
Release
```

---

## Bounded Concurrency

Prevent resource exhaustion.

Use:

* worker pools
* semaphores
* backpressure

---

# 15. Production Debugging

---

## Stack Dump

```bash
kill -QUIT <pid>
```

Shows:

```text
goroutine 1
chan receive

goroutine 2
mutex lock

goroutine 3
waitgroup wait
```

---

## Goroutine Profile

```bash
go tool pprof
```

Look for:

```text
chan send
chan receive
sync.Mutex.Lock
sync.WaitGroup.Wait
```

---

## Scheduler Trace

```bash
GODEBUG=schedtrace=1000
```

Shows:

```text
runnable goroutines
idle processors
blocked goroutines
```

---

## Goroutine Leak Detection

Useful indicators:

```text
Increasing goroutine count

No CPU usage

Growing latency
```

---

# 16. Performance Considerations

Deadlock prevention strategies have tradeoffs.

---

## Coarse-Grained Locking

Pros

```text
Easy
Safe
```

Cons

```text
Less Concurrency
```

---

## Fine-Grained Locking

Pros

```text
Higher Throughput
```

Cons

```text
Higher Deadlock Risk
```

---

## Timeouts

Pros

```text
Avoid Infinite Waits
```

Cons

```text
Complexity
False Timeouts
```

---

# 17. Common Interview Questions

---

### Can buffered channels prevent deadlocks?

No.

They only postpone blocking.

---

### Why are Go mutexes non-reentrant?

To keep implementation simpler and faster.

The runtime does not track lock ownership.

---

### What state is a blocked goroutine in?

```text
_Gwaiting
```

---

### Can runtime detect all deadlocks?

No.

Only global deadlocks.

---

### Why is lock ordering important?

It eliminates circular wait.

---

# 18. Google L5 Follow-up Questions

---

## Q1. Why does Go detect some deadlocks but not all?

### Answer

Go runtime only detects:

```text
All goroutines blocked

AND

No future events can make progress
```

Examples detected:

```go
<-ch
```

with no sender.

---

Examples NOT detected:

```go
go worker()

select {}
```

Program hangs forever.

Runtime sees active goroutines and scheduler state, so no panic occurs.

Production deadlocks are often not detected.

---

## Q2. Explain Channel Deadlock Internals

### Answer

For:

```go
<-ch
```

Runtime executes:

```text
chanrecv()
```

Flow:

```text
Lock Channel
      ↓
Check Buffer
      ↓
Check Sender Queue
      ↓
No Sender
      ↓
Park Goroutine
      ↓
_Gwaiting
```

Scheduler removes goroutine from run queue.

If all goroutines reach waiting state:

```text
Deadlock Panic
```

---

## Q3. Difference Between Deadlock and Livelock?

### Deadlock

```text
No activity
No progress
```

---

### Livelock

```text
Activity exists
No progress
```

---

Example:

```text
Deadlock:
A waits B
B waits A

Livelock:
A retries
B retries
Forever
```

---

## Q4. Why Are Go Mutexes Non-Reentrant?

### Answer

Tracking lock ownership would require:

* additional metadata
* owner tracking
* recursion counters

This adds overhead.

Go prioritizes:

* simplicity
* speed
* minimal runtime cost

Therefore:

```go
mu.Lock()
mu.Lock()
```

deadlocks.

---

## Q5. How Would You Design a System That Guarantees Deadlock Freedom?

### Answer

Apply:

1. Global lock ordering
2. No nested locks
3. Context timeouts
4. Channel ownership rules
5. Bounded concurrency
6. Avoid circular dependencies

A common strategy:

```text
Acquire Resources
       ↓
Perform Work
       ↓
Release Resources
```

without nested acquisitions.

---

## Q6. How Do You Debug Deadlocks in Production?

### Answer

Tools:

```bash
kill -QUIT
```

```bash
go tool pprof
```

```bash
GODEBUG=schedtrace=1000
```

Inspect:

```text
Blocked Goroutines
Mutex Contention
WaitGroup Waits
Channel Receives
Channel Sends
```

Then reconstruct dependency chain.

---

## Q7. Can Buffered Channels Eliminate Deadlocks?

### Answer

No.

Example:

```go
ch := make(chan int, 2)

ch <- 1
ch <- 2
ch <- 3
```

Third send blocks.

Buffered channels only delay blocking.

---

## Q8. How Would You Enforce Lock Ordering Across Large Codebases?

### Answer

Senior-level approaches:

* documented lock hierarchy
* lock acquisition guidelines
* code reviews
* static analysis
* lock wrappers

Example:

```text
UserLock
   ↓
TenantLock
   ↓
GlobalLock
```

Never reverse order.

---

## Q9. What Runtime State Is a Blocked Goroutine In?

### Answer

```text
_Gwaiting
```

Possible reasons:

```text
Channel Send
Channel Receive
Mutex Lock
WaitGroup Wait
Network Poll Wait
```

Scheduler excludes it from execution.

---

## Q10. How Does The Scheduler Participate In Deadlock Detection?

### Answer

Scheduler maintains:

```text
Runnable Goroutines
Waiting Goroutines
Running Goroutines
```

Detection:

```text
Run Queue Empty
       ↓
No Timers
       ↓
No Network Events
       ↓
All Goroutines Waiting
       ↓
Deadlock Panic
```

The scheduler is the component that determines whether future progress is possible.

---

# 19. Interview Cheat Sheet

## Deadlock Definition

```text
Permanent waiting with no possibility of progress.
```

---

## Four Coffman Conditions

```text
Mutual Exclusion
Hold and Wait
No Preemption
Circular Wait
```

---

## Most Common Go Deadlocks

```text
Receive Without Sender
Send Without Receiver
Range Without Close
Missing WaitGroup Done
Circular Mutex Dependency
Self Mutex Lock
```

---

## Runtime Deadlock Panic

```text
fatal error:
all goroutines are asleep - deadlock!
```

---

## Goroutine Waiting State

```text
_Gwaiting
```

---

## Runtime Detects?

```text
Global Deadlocks
```

---

## Runtime Does NOT Detect?

```text
Distributed Deadlocks
Application-Level Deadlocks
Many Production Hangs
```

---

## Best Prevention Strategy

```text
Lock Ordering
Timeouts
Context Cancellation
Avoid Nested Locks
Bounded Concurrency
```

---

# 20. Key Takeaways

* Deadlocks are correctness failures.
* A deadlock means permanent waiting.
* Channel, Mutex, and WaitGroup deadlocks are common in Go.
* Go runtime only detects global deadlocks.
* Most production deadlocks are application-level and require debugging.
* Global lock ordering is the most effective prevention strategy.
* Buffered channels do not eliminate deadlocks.
* Senior engineers must understand scheduler interaction and runtime states.
* Deadlock debugging is a critical production skill.
* Google L5 candidates should be able to explain, detect, debug, prevent, and reason about deadlocks at runtime level.

---
