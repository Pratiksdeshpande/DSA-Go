# Goroutines in Go

> Deep dive into Goroutines, runtime scheduling, stack management, concurrency behavior, performance implications, and production-grade patterns in Go.

---

# Table of Contents

* [1. Introduction](#1-introduction)
* [2. What is a Goroutine?](#2-what-is-a-goroutine)
* [3. Why Goroutines Exist](#3-why-goroutines-exist)
* [4. Goroutines vs OS Threads](#4-goroutines-vs-os-threads)
* [5. Creating Goroutines](#5-creating-goroutines)
* [6. Internal Runtime Behavior](#6-internal-runtime-behavior)
* [7. The G-M-P Scheduler Model](#7-the-g-m-p-scheduler-model)
* [8. Goroutine Lifecycle](#8-goroutine-lifecycle)
* [9. Goroutine Stack Management](#9-goroutine-stack-management)
* [10. Dynamic Stack Growth](#10-dynamic-stack-growth)
* [11. Parking and Unparking](#11-parking-and-unparking)
* [12. Blocking Behavior](#12-blocking-behavior)
* [13. Real Production Use Cases](#13-real-production-use-cases)
* [14. Goroutine Synchronization](#14-goroutine-synchronization)
* [15. Common Goroutine Patterns](#15-common-goroutine-patterns)
* [16. Common Interview Questions](#16-common-interview-questions)
* [17. Hidden Traps](#17-hidden-traps)
* [18. Goroutine Leaks](#18-goroutine-leaks)
* [19. Performance Considerations](#19-performance-considerations)
* [20. Debugging Goroutines](#20-debugging-goroutines)
* [21. Best Practices](#21-best-practices)
* [22. Key Takeaways](#22-key-takeaways)
* [23. Summary](#23-summary)

---

# 1. Introduction

Goroutines are the foundational concurrency primitive in Go.

They enable Go applications to efficiently handle:

* APIs
* distributed systems
* streaming pipelines
* event-driven systems
* background jobs
* concurrent processing

Unlike traditional threads, goroutines are lightweight and managed by the Go runtime.

---

# 2. What is a Goroutine?

A goroutine is:

> A lightweight independently executing function managed by the Go runtime.

---

## Example

```go
go processOrder()
```

This creates a new goroutine.

---

# 3. Why Goroutines Exist

Traditional OS threads are expensive because they require:

* kernel management
* large stacks
* expensive context switching

Goroutines solve this by being:

* lightweight
* runtime-managed
* dynamically scheduled

This allows Go to efficiently support massive concurrency.

---

# 4. Goroutines vs OS Threads

| Feature           | Goroutines          | OS Threads        |
| ----------------- | ------------------- | ----------------- |
| Managed By        | Go Runtime          | Operating System  |
| Stack Size        | Small Dynamic Stack | Large Fixed Stack |
| Creation Cost     | Cheap               | Expensive         |
| Context Switching | Lightweight         | Heavyweight       |
| Scalability       | Millions Possible   | Limited           |

---

# 5. Creating Goroutines

## Basic Example

```go
package main

import "fmt"

func task() {
    fmt.Println("Running task")
}

func main() {
    go task()
}
```

---

# 6. Internal Runtime Behavior

When a goroutine is created:

1. Runtime allocates goroutine structure
2. Small stack is allocated
3. Goroutine added to run queue
4. Scheduler eventually executes it

---

# 7. The G-M-P Scheduler Model

Go scheduler uses:

| Component | Meaning             |
| --------- | ------------------- |
| G         | Goroutine           |
| M         | Machine (OS Thread) |
| P         | Processor           |

Go uses M:N scheduling:

* many goroutines
* mapped onto fewer threads

---

# 8. Goroutine Lifecycle

Common goroutine states:

| State    | Meaning             |
| -------- | ------------------- |
| Runnable | Ready to run        |
| Running  | Currently executing |
| Waiting  | Blocked             |
| Dead     | Completed           |

---

# 9. Goroutine Stack Management

Goroutines use very small initial stacks.

Approximate initial size:

* ~2 KB

Stacks grow dynamically when needed.

---

# 10. Dynamic Stack Growth

When stack becomes full:

1. runtime allocates larger stack
2. existing stack copied
3. pointers updated

Modern Go uses contiguous dynamically growing stacks.

---

# 11. Parking and Unparking

When goroutine blocks:

* scheduler parks it
* OS thread reused elsewhere

When work becomes available:

* scheduler unparks goroutine

This is critical for scalable I/O systems.

---

# 12. Blocking Behavior

Goroutines block during:

* channel operations
* mutex locking
* network I/O
* syscalls
* sleep operations

Blocked goroutines do not waste CPU resources.

---

# 13. Real Production Use Cases

## API Servers

Each HTTP request handled concurrently.

---

## Background Jobs

Examples:

* metrics collection
* async notifications
* cache refresh

---

## Streaming Systems

Used in:

* Kafka consumers
* event processors
* ETL pipelines

---

# 14. Goroutine Synchronization

Common synchronization mechanisms:

* channels
* WaitGroup
* Mutex
* Context cancellation

---

## WaitGroup Example

```go
package main

import (
    "sync"
)

func worker(wg *sync.WaitGroup) {
    defer wg.Done()
}

func main() {
    var wg sync.WaitGroup

    wg.Add(1)

    go worker(&wg)

    wg.Wait()
}
```

---

# 15. Common Goroutine Patterns

## Worker Pool

Bounded concurrent processing.

---

## Fan-Out / Fan-In

Parallel task execution and aggregation.

---

## Pipeline

Multi-stage processing systems.

---

## Bounded Concurrency with semaphores.

```go
package main

import (
    "sync"
)

func process(job int) {}

func main() {
    jobs := make([]int, 1000)

    semaphore := make(chan struct{}, 10)

    var wg sync.WaitGroup

    for _, job := range jobs {
        semaphore <- struct{}{}

        wg.Add(1)

        go func(job int) {
            defer wg.Done()
            defer func() { <-semaphore }()

            process(job)
        }(job)
    }

    wg.Wait()
}
```

This prevents unbounded goroutine explosion

Read more about [Semaphore](../../semaphore/README.md)

---

# 16. Common Interview Questions

## Why are goroutines lightweight?

Because they use:

* small stacks
* runtime scheduling
* lightweight context switching

---

## Are goroutines parallel?

Not always.

Depends on:

* CPU cores
* scheduler
* GOMAXPROCS

---

## What happens when goroutine blocks?

Scheduler parks it and schedules another goroutine.

---

# 17. Hidden Traps

## Loop Variable Capture

Incorrect:

```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
```

Correct:

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        i := i

        go func() {
            fmt.Println(i)
		}()
    }
}
```

---

## Main Exits Early

Need synchronization to avoid premature program termination.

---

# 18. Goroutine Leaks

Leaked goroutines remain blocked forever.

Example:

```go
func worker(ch chan int) {
    <-ch
}
```

If nobody sends to channel:

* goroutine leaks

---

# 19. Performance Considerations

## Too Many Goroutines

Can cause:

* scheduler pressure
* memory growth
* GC overhead
* contention

---

## Blocking Syscalls

May increase thread creation.

---

## Stack Growth Costs

Stack copying introduces runtime overhead.

---

# 20. Debugging Goroutines

Useful tools:

* pprof
* runtime.NumGoroutine()
* go tool trace
* stack dumps

---

## Example

```go
fmt.Println(runtime.NumGoroutine())
```

---

# 21. Best Practices

* Avoid unbounded goroutines
* Use worker pools
* Use context cancellation
* Prevent goroutine leaks
* Avoid blocking indefinitely
* Use synchronization correctly

---

# 22. Key Takeaways

* Goroutines are lightweight concurrent execution units.
* Managed by Go runtime, not directly by OS.
* Scheduled using G-M-P model.
* Use dynamic stacks for scalability.
* Efficient for I/O-bound systems.
* Require careful synchronization and lifecycle management.

---

# 23. Summary

Goroutines are one of the most important innovations in Go.

Understanding goroutines deeply requires knowledge of:

* runtime internals
* scheduler behavior
* memory management
* synchronization
* performance engineering

Mastering goroutines is essential for:

* senior Go interviews
* scalable system design
* production-grade backend engineering

---
