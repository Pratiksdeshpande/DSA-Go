# Goroutine Leaks in Go

> A deep dive into Goroutine Leaks, runtime behavior, detection, prevention techniques, debugging strategies, and Google L5 interview preparation.

---

# Table of Contents

* [1. Introduction](#1-introduction)
* [2. What is a Goroutine Leak?](#2-what-is-a-goroutine-leak)
* [3. Why Goroutine Leaks Are Dangerous](#3-why-goroutine-leaks-are-dangerous)
* [4. Goroutine Leak vs Memory Leak](#4-goroutine-leak-vs-memory-leak)
* [5. Runtime Internals](#5-runtime-internals)

  * [5.1 Goroutine Lifecycle](#51-goroutine-lifecycle)
  * [5.2 What Happens When a Goroutine Leaks](#52-what-happens-when-a-goroutine-leaks)
  * [5.3 Scheduler Impact](#53-scheduler-impact)
  * [5.4 GC Impact](#54-gc-impact)
* [6. Common Goroutine Leak Patterns](#6-common-goroutine-leak-patterns)

  * [6.1 Channel Receive Leak](#61-channel-receive-leak)
  * [6.2 Channel Send Leak](#62-channel-send-leak)
  * [6.3 Forgotten Context Cancellation](#63-forgotten-context-cancellation)
  * [6.4 HTTP Calls Without Timeouts](#64-http-calls-without-timeouts)
  * [6.5 Worker Pool Leaks](#65-worker-pool-leaks)
  * [6.6 Pipeline Leaks](#66-pipeline-leaks)
* [7. Real Production Scenarios](#7-real-production-scenarios)
* [8. Leak Detection Techniques](#8-leak-detection-techniques)
* [9. Preventing Goroutine Leaks](#9-preventing-goroutine-leaks)
* [10. Production-Grade Patterns](#10-production-grade-patterns)
* [11. Performance Considerations](#11-performance-considerations)
* [12. Whiteboard Mental Models](#12-whiteboard-mental-models)
* [13. Debugging Goroutine Leaks in Production](#13-debugging-goroutine-leaks-in-production)
* [14. Google L5 Interview Questions](#14-google-l5-interview-questions)
* [15. Key Takeaways](#15-key-takeaways)

---

# 1. Introduction

Goroutine leaks are one of the most common causes of production instability in Go services.

Unlike crashes, goroutine leaks usually grow slowly and silently.

Symptoms often appear much later as:

* Increased memory usage
* Increased GC pauses
* Increased goroutine count
* Increased scheduler overhead
* Service degradation
* OOM crashes

Understanding goroutine leaks is essential for:

* Backend services
* Distributed systems
* Event-driven architectures
* Streaming systems
* Microservices

---

# 2. What is a Goroutine Leak?

A goroutine leak occurs when:

> A goroutine remains alive indefinitely after it has stopped serving useful business purpose.

The goroutine never exits.

The runtime cannot reclaim it.

---

## Example

```go
func main() {
	ch := make(chan int)

	go func() {
		<-ch
	}()

	time.Sleep(time.Hour)
}
```

The goroutine waits forever.

No sender exists.

The goroutine never terminates.

Leak.

---

# 3. Why Goroutine Leaks Are Dangerous

Many developers think:

```text
Goroutines are cheap.
```

True.

But:

```text
1 Goroutine      → Fine
100 Goroutines   → Fine
10,000 Goroutines → Problem
1,000,000 Goroutines → Disaster
```

A leaked goroutine consumes:

* Stack memory
* Runtime metadata
* Scheduler resources
* GC scanning resources

---

## Production Growth Example

```text
100 Requests/sec
      ↓
1 Leaked Goroutine per Request
      ↓
360,000 Leaked Goroutines per Hour
      ↓
8.6 Million per Day
```

Eventually:

```text
Memory Exhaustion
GC Pressure
OOM Crash
```

---

# 4. Goroutine Leak vs Memory Leak

| Memory Leak              | Goroutine Leak                     |
| ------------------------ | ---------------------------------- |
| Memory never released    | Goroutine never exits              |
| Heap grows continuously  | Goroutine count grows continuously |
| GC cannot reclaim memory | Runtime cannot reclaim goroutine   |
| Usually memory-centric   | Memory + scheduler + GC impact     |

---

## Important

Goroutine leaks often create memory leaks.

Because leaked goroutines hold references to heap objects.

---

# 5. Runtime Internals

---

# 5.1 Goroutine Lifecycle

When Go executes:

```go
go worker()
```

Runtime allocates:

```text
G Structure
Stack
Scheduler Metadata
```

---

## Lifecycle

```text
Created
   ↓
Runnable
   ↓
Running
   ↓
Blocked (optional)
   ↓
Running Again
   ↓
Completed
   ↓
Destroyed
```

---

# 5.2 What Happens When a Goroutine Leaks

Instead of reaching:

```text
Completed
   ↓
Destroyed
```

It remains:

```text
Blocked Forever
```

Example:

```go
<-ch
```

No sender exists.

---

## Result

Runtime keeps:

```text
G Structure
Stack
References
Scheduling Metadata
```

alive forever.

---

# 5.3 Scheduler Impact

Go scheduler tracks every goroutine.

Large goroutine counts increase:

```text
Run Queue Management
Work Stealing
Scheduling Overhead
Runtime Bookkeeping
```

Even blocked goroutines are not completely free.

---

# 5.4 GC Impact

Garbage Collector scans goroutine stacks.

Example:

```go
buffer := make([]byte, 50*1024*1024)
```

If a leaked goroutine references:

```text
50 MB Buffer
```

GC cannot reclaim it.

---

## Consequences

```text
More Goroutines
      ↓
More Stack Scanning
      ↓
Longer GC Cycles
      ↓
Higher Latency
```

---

# 6. Common Goroutine Leak Patterns

---

# 6.1 Channel Receive Leak

```go
func main() {
	ch := make(chan int)

	go func() {
		<-ch
	}()

	time.Sleep(time.Hour)
}
```

---

## Problem

```text
Receiver waits forever.
```

---

## Diagram

```text
Goroutine
     ↓
Waiting on Channel
     ↓
No Sender
     ↓
Forever
```

---

# 6.2 Channel Send Leak

```go
func main() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	time.Sleep(time.Hour)
}
```

---

## Problem

```text
No receiver exists.
```

Sender blocks forever.

---

# 6.3 Forgotten Context Cancellation

```go
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			doWork()
		}
	}
}
```

---

## Problem

Caller forgets:

```go
cancel()
```

Worker runs forever.

---

# 6.4 HTTP Calls Without Timeouts

Bad:

```go
resp, err := http.Get(url)
```

Remote server hangs forever.

---

## Better

```go
client := http.Client{
	Timeout: 5 * time.Second,
}
```

---

# 6.5 Worker Pool Leaks

```go
for {
	job := <-jobs
	process(job)
}
```

Application exits.

Workers still wait forever.

Leak.

---

## Correct

```go
select {
case job := <-jobs:
	process(job)

case <-ctx.Done():
	return
}
```

---

# 6.6 Pipeline Leaks

Example:

```go
for i := 0; i < 100; i++ {
	go worker(results)
}
```

Consumer exits early.

Workers continue sending.

Workers become permanently blocked.

---

# 7. Real Production Scenarios

---

## Metrics Publisher Leak

```go
func HandleRequest() {
	go publishMetrics()
}
```

Inside:

```go
func publishMetrics() {
	for {
		sendMetric()
		time.Sleep(time.Second)
	}
}
```

---

## Traffic

```text
1000 Requests / Second
```

Creates:

```text
1000 New Infinite Goroutines / Second
```

Eventually:

```text
Millions of Goroutines
```

Service crashes.

---

# 8. Leak Detection Techniques

---

# runtime.NumGoroutine()

```go
fmt.Println(runtime.NumGoroutine())
```

Monitor over time.

Example:

```text
100
250
500
2000
10000
```

Leak suspected.

---

# pprof Goroutine Dump

Enable:

```go
import _ "net/http/pprof"
```

Then:

```bash
curl http://localhost:6060/debug/pprof/goroutine?debug=2
```

---

## Example Output

```text
goroutine 101:
chan receive

goroutine 102:
net/http

goroutine 103:
sync.Mutex
```

Extremely useful.

---

# Goroutine Profile

```bash
go tool pprof
```

Analyze:

```text
Who Created Goroutines?
Where Are They Blocked?
How Many Exist?
```

---

# 9. Preventing Goroutine Leaks

---

## Rule #1

Every goroutine must have an exit path.

---

## Rule #2

Use context cancellation.

```go
ctx, cancel := context.WithCancel(...)
defer cancel()
```

---

## Rule #3

Use timeouts.

```go
context.WithTimeout(...)
```

---

## Rule #4

Close channels properly.

```go
close(ch)
```

---

## Rule #5

Bound concurrency.

Never create unlimited goroutines.

---

# 10. Production-Grade Patterns

---

# Context-Aware Worker

```go
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			process()
		}
	}
}
```

---

# Graceful Channel Consumer

```go
for value := range jobs {
	process(value)
}
```

Automatically exits when channel closes.

---

# WaitGroup Coordination

```go
var wg sync.WaitGroup
```

Ensures workers terminate before shutdown.

---

# Bounded Worker Pool

Good:

```text
10 Workers
10000 Jobs
```

Bad:

```text
10000 Workers
10000 Jobs
```

---

# 11. Performance Considerations

---

## Are Blocked Goroutines Free?

No.

They still consume:

* Stack memory
* Runtime metadata
* GC scanning cost
* Scheduler bookkeeping

---

## Memory Cost

Suppose:

```text
2 KB Initial Stack
```

One million leaked goroutines:

```text
~2 GB
```

before stack growth.

---

## GC Cost

```text
More Goroutines
      ↓
More Stack Scanning
      ↓
Longer GC Cycles
```

---

## Scheduler Cost

```text
More Goroutines
      ↓
More Runtime Overhead
```

---

# 12. Whiteboard Mental Models

Whenever you see:

```go
go func() {
	...
}()
```

Ask:

```text
How Does This Goroutine Die?
```

---

## Leak Safety Checklist

```text
1. What is the exit condition?
2. Can it block forever?
3. Who cancels it?
4. Who closes channels?
5. What happens during shutdown?
```

If you cannot answer all five:

```text
Potential Goroutine Leak
```

---

# 13. Debugging Goroutine Leaks in Production

Step 1:

```go
runtime.NumGoroutine()
```

Monitor trends.

---

Step 2:

Collect goroutine dump.

```bash
curl /debug/pprof/goroutine?debug=2
```

---

Step 3:

Group blocked states.

```text
Channel Receive
Channel Send
Mutex Wait
Network Read
Select Wait
```

---

Step 4:

Find creator stack traces.

Determine:

```text
Who Started Them?
Why Didn't They Exit?
```

---

Step 5:

Fix root cause.

Common fixes:

* Context cancellation
* Timeout
* Channel closure
* Worker pool shutdown

---

# 14. Google L5 Interview Questions

---

## Q1. What is a goroutine leak?

A goroutine that remains alive indefinitely after it stops serving useful business purpose.

---

## Q2. Why are goroutine leaks dangerous?

They consume:

* Memory
* Scheduler resources
* GC resources

and may eventually crash the service.

---

## Q3. Can blocked goroutines be leaks?

Yes.

If blocked forever unexpectedly.

---

## Q4. Why do goroutine leaks increase GC work?

GC must scan their stacks and referenced objects.

---

## Q5. How do you detect goroutine leaks?

* runtime.NumGoroutine()
* pprof
* stack dumps
* monitoring metrics

---

## Q6. Why are goroutine leaks harder to detect than memory leaks?

Because memory growth may be slow while goroutine count quietly increases.

---

## Q7. How would you design a leak-safe worker pool?

Use:

* Fixed worker count
* Context cancellation
* Channel closure
* WaitGroup synchronization
* Graceful shutdown

---

## Q8. Can a goroutine leak without consuming CPU?

Yes.

Blocked goroutines often consume little CPU but still consume memory and GC resources.

---

## Q9. What production metrics would you monitor?

* Goroutine count
* Memory usage
* GC pause times
* Heap growth
* Request latency

---

## Q10. What is the first question you ask when reviewing concurrent code?

```text
How Does This Goroutine Exit?
```

---

# 15. Key Takeaways

* Every goroutine must have a termination path.
* Blocked goroutines are not free.
* Goroutine leaks often become memory leaks.
* Context cancellation is the primary defense.
* Timeouts are mandatory for external calls.
* Monitor goroutine count in production.
* Use pprof to investigate leaks.
* Always ask: "How does this goroutine die?"
* Bounded concurrency prevents runaway growth.
* Senior Go engineers design systems to be leak-safe from day one.

---
