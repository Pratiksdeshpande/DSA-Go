# Complete Goroutine Lifecycle with G-M-P model Through One Production Example

> Goal: Understand Goroutines, Scheduler, G-M-P Model, Work Stealing, Channel Blocking, Syscalls, Netpoller, Memory Allocation, Garbage Collection, Goroutine Leaks, and Runtime Internals using one realistic example.

---

# Table of Contents

* [Example Overview](#example-overview)
* [Complete Example](#complete-example)
* [Runtime Startup Sequence](#runtime-startup-sequence)
* [G0 vs G1](#g0-vs-g1)
* [Goroutine Creation](#goroutine-creation)
* [Execution With One OS Thread](#execution-with-one-os-thread)
* [Execution With Two OS Threads](#execution-with-two-os-threads)
* [Channel Blocking](#channel-blocking)
* [Network I/O and Netpoller](#network-io-and-netpoller)
* [File I/O and Syscalls](#file-io-and-syscalls)
* [P Detaching From M](#p-detaching-from-m)
* [Work Stealing](#work-stealing)
* [Memory Allocation](#memory-allocation)
* [Escape Analysis](#escape-analysis)
* [Garbage Collection](#garbage-collection)
* [Goroutine Leaks](#goroutine-leaks)
* [Complete Lifecycle Timeline](#complete-lifecycle-timeline)
* [Interview Questions](#interview-questions)

---

# Example Overview

This example simulates a production backend service.

Each worker:

1. Receives a URL from a channel.
2. Performs an HTTP request.
3. Reads the response.
4. Writes response data into a file.
5. Sends results back through another channel.
6. Exits gracefully when work is completed.

This single example touches almost every important Go runtime concept.

---

# Complete Example

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func worker(
	ctx context.Context,
	id int,
	jobs <-chan string,
	results chan<- string,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {

		case <-ctx.Done():
			return

		case url, ok := <-jobs:

			if !ok {
				return
			}

			resp, err := http.Get(url)

			if err != nil {
				continue
			}

			body := make([]byte, 1024)

			resp.Body.Read(body)

			file, _ := os.Create(
				fmt.Sprintf("worker-%d.txt", id),
			)

			file.Write(body)

			results <- fmt.Sprintf(
				"worker %d processed %s",
				id,
				url,
			)

			file.Close()
			resp.Body.Close()
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(
		context.Background(),
	)
	defer cancel()

	jobs := make(chan string)

	results := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {

		wg.Add(1)

		go worker(
			ctx,
			i,
			jobs,
			results,
			&wg,
		)
	}

	go func() {

		urls := []string{
			"https://a.com",
			"https://b.com",
			"https://c.com",
			"https://d.com",
		}

		for _, url := range urls {
			jobs <- url
		}

		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
```

---

# Runtime Startup Sequence

Before your code executes, the Go runtime starts first.

When:

```bash
go run main.go
```

is executed, the operating system creates:

```text
Process
└── Thread-1
```

Go runtime takes ownership of this first thread.

The runtime initializes:

```text
G0
M0
P0
```

Where:

```text
G = Goroutine
M = Machine (OS Thread)
P = Processor (Scheduler Context)
```

Initial state:

```text
G0
 |
M0
 |
P0
```

At this moment your code has not started.

The runtime is still initializing:

* Scheduler
* Memory Allocator
* Garbage Collector
* Timer System
* Netpoller

---

# G0 vs G1

This is one of the most misunderstood concepts.

Many developers believe:

```text
main() runs on G0
```

This is incorrect.

---

## G0

G0 is a special runtime goroutine.

Its purpose is:

* Scheduling
* Context Switching
* Runtime Bookkeeping
* Stack Management
* GC Coordination

User code never executes on G0.

Think of G0 as:

```text
Scheduler Stack
```

rather than:

```text
Application Goroutine
```

---

## G1

After runtime initialization:

```text
runtime.main()
```

creates:

```text
G1
```

G1 becomes:

```go
func main()
```

State:

```text
             G0
              |
             M0
              |
             P0

Runnable:

G1
```

Scheduler chooses G1.

Now application execution begins.

---

# Goroutine Creation

When:

```go
go worker(...)
```

executes,

the compiler transforms it approximately into:

```go
runtime.newproc(...)
```

Runtime allocates:

```text
G2
G3
G4
G5
```

for the four workers.

Each new goroutine receives:

```text
~2 KB Initial Stack
```

State:

```text
Runnable
```

Then each goroutine is placed into:

```text
P Local Queue
```

Example:

```text
P0 Queue

[G2 G3 G4 G5]
```

Important:

Creating a goroutine does not create an OS thread.

This is why goroutines are extremely cheap.

---

# Execution With One OS Thread

Assume:

```go
runtime.GOMAXPROCS(1)
```

Runtime:

```text
P0
M0
```

Only one goroutine can execute at a time.

State:

```text
M0
 |
P0
 |
G1
```

When G1 blocks:

```go
result := <-results
```

Scheduler runs:

```text
G2
```

Then:

```text
G3
```

Then:

```text
G4
```

and so on.

This is concurrency.

Not parallelism.

---

# Execution With Two OS Threads

Assume:

```go
runtime.GOMAXPROCS(2)
```

Runtime now creates:

```text
P0
P1
```

and may create:

```text
M0
M1
```

State:

```text
M0 -> P0 -> G2

M1 -> P1 -> G3
```

Now two goroutines execute simultaneously.

This is parallelism.

---

# Channel Blocking

Worker executes:

```go
url := <-jobs
```

Suppose:

```text
jobs channel is empty
```

Runtime changes goroutine state:

```text
Runnable
    ↓
Waiting
```

Only goroutine is parked.

OS thread is not blocked.

Scheduler immediately executes another runnable goroutine.

This is why channel operations are efficient.

---

# Network I/O and Netpoller

Worker executes:

```go
http.Get(url)
```

Internally:

```text
Socket Connect
Socket Read
Socket Write
```

Instead of blocking a thread, Go uses:

```text
Netpoller
```

Netpoller uses:

Linux  -> epoll

Mac    -> kqueue

Windows -> IOCP

Runtime registers socket.

Goroutine becomes:

```text
Waiting(Network)
```

Thread becomes available for other work.

When network response arrives:

```text
Kernel
   ↓
Netpoller
   ↓
Goroutine Runnable
```

The goroutine is placed back into a run queue.

This is one reason Go handles massive numbers of network connections efficiently.

---

# File I/O and Syscalls

Worker executes:

```go
file.Write(...)
```

Eventually:

```text
write()
```

system call is executed.

Unlike sockets, file operations often block.

OS thread may enter kernel mode.

State:

```text
G2
 |
M0 (blocked)
```

Without special handling:

```text
Entire scheduler would stop.
```

Go runtime avoids this.

---

# P Detaching From M

Suppose:

```text
M0
```

enters a blocking syscall.

Runtime performs:

```text
Detach P0 from M0
```

Result:

```text
Before

M0 -> P0

After

M0 (blocked)

P0 (free)
```

Runtime attaches:

```text
P0 -> M1
```

Now other goroutines continue executing.

This is one of the most important scheduler optimizations.

---

# Work Stealing

Assume:

```text
P0 Queue

[G2 G3 G4 G5 G6 G7]
```

while:

```text
P1 Queue

[]
```

P1 becomes idle.

Its scheduler executes:

```text
Steal Half
```

Result:

```text
Before

P0 -> [G2 G3 G4 G5 G6 G7]

P1 -> []

After

P0 -> [G2 G3 G4]

P1 -> [G5 G6 G7]
```

Work stealing is performed by the idle P's scheduler.

Purpose:

```text
Load Balancing
```

across processors.

---

# Memory Allocation

Worker executes:

```go
body := make([]byte, 1024)
```

Memory is allocated.

Allocation may happen:

```text
Stack
or
Heap
```

depending on escape analysis.

Runtime allocator manages these allocations efficiently using per-P caches.

---

# Escape Analysis

Compiler decides:

Can object stay on stack?

or

Must move to heap?

Example:

```go
func f() *int {

	x := 10

	return &x
}
```

x escapes.

Compiler moves x to heap.

Result:

```text
More GC work
```

Understanding escape analysis is important for performance-sensitive Go services.

---

# Garbage Collection

Objects allocated on heap are tracked by GC.

Modern Go uses:

```text
Concurrent Tri-Color Mark-Sweep
```

Phases:

```text
Mark Roots

Concurrent Mark

Mark Termination

Concurrent Sweep
```

GC scans:

* Global Variables
* Heap Objects
* Goroutine Stacks

Unreachable objects are reclaimed.

GC runs concurrently with application goroutines.

---

# Goroutine Leaks

One of the most common production issues.

Example:

```go
func leak() {

	ch := make(chan int)

	go func() {
		<-ch
	}()

}
```

Nobody sends on channel.

Goroutine waits forever.

Leak.

---

Common Causes

### Blocked Receive

```go
<-ch
```

forever.

### Blocked Send

```go
ch <- value
```

forever.

### Missing Context Cancellation

```go
for {
	select {
	default:
	}
}
```

never exits.

---

Always ensure goroutines have a clear termination path.

---

# Complete Lifecycle Timeline

```text
Process Starts

↓

Runtime Initialization

↓

G0 Created

↓

P0 Created

↓

runtime.main()

↓

G1 Created

↓

main() Starts

↓

go worker()

↓

runtime.newproc()

↓

G2 Created

↓

Runnable

↓

Enqueued Into P Queue

↓

Scheduler Picks G2

↓

Running

↓

Network Call

↓

Waiting(Netpoller)

↓

Network Event Arrives

↓

Runnable Again

↓

File Write

↓

Syscall

↓

M Blocks

↓

P Detached

↓

P Attached To Another M

↓

Execution Continues

↓

Worker Completes

↓

wg.Done()

↓

Goroutine Dead

↓

Stack Reclaimed

↓

GC Eventually Reclaims Heap Objects

↓

Program Exit
```

---

# Interview Questions

### What happens when `go func()` is executed?

1. Compiler generates runtime.newproc().
2. Runtime allocates G structure.
3. Initial stack is allocated.
4. Goroutine becomes runnable.
5. Added to current P run queue.
6. Scheduler eventually executes it.

---

### What is G0?

Runtime scheduler goroutine.

Never executes user code.

---

### What is G1?

Main application goroutine.

Executes func main().

---

### When does work stealing happen?

When a P becomes idle and has no runnable goroutines.

It steals approximately half the runnable goroutines from another busy P.

---

### What happens during a blocking syscall?

The runtime detaches P from blocked M and attaches it to another M.

---

### What is Netpoller?

Runtime component that manages network I/O readiness events without blocking threads.

---

### What causes goroutine leaks?

Blocked sends, blocked receives, missing cancellation, and goroutines waiting forever on resources.

---

### Why are goroutines cheap?

* ~2KB initial stack
* User-space scheduling
* Growable stacks
* Reuse of OS threads
* Efficient scheduler

```
```
