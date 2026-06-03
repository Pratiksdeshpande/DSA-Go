# File I/O vs Network I/O in Goroutines

## Table of Contents

* [Introduction](#introduction)
* [The Core Difference](#the-core-difference)
* [How File I/O Works](#how-file-io-works)
* [How Network I/O Works](#how-network-io-works)
* [Go Netpoller Internals](#go-netpoller-internals)
* [Why Netpoller Works for Sockets but Not Files](#why-netpoller-works-for-sockets-but-not-files)
* [Scheduler Behavior During File I/O](#scheduler-behavior-during-file-io)
* [Scheduler Behavior During Network I/O](#scheduler-behavior-during-network-io)
* [File I/O Can Cause Thread Growth](#file-io-can-cause-thread-growth)
* [Why Go Can Handle Millions of Network Connections](#why-go-can-handle-millions-of-network-connections)
* [Common Misconceptions](#common-misconceptions)
* [Production Considerations](#production-considerations)
* [Summary Table](#summary-table)
* [L5 Interview Questions and Answers](#l5-interview-questions-and-answers)

---

# Introduction

Many Go developers know that:

```text
Network I/O scales well.
File I/O scales less efficiently.
```

However, the reason is not simply that:

```text
Network = Fast
File = Slow
```

The real reason lies inside the Go runtime.

Understanding this distinction is important for senior-level interviews because it requires knowledge of:

* Goroutines
* Scheduler (G-M-P Model)
* Syscalls
* Netpoller
* OS Thread Management
* Runtime Scalability

---

# The Core Difference

The fundamental difference is:

```text
Network I/O
    Uses Netpoller
    ↓
    Goroutine parks
    OS Thread remains available

File I/O
    Usually bypasses Netpoller
    ↓
    OS Thread enters blocking syscall
    ↓
    Runtime detaches P
```

Mental model:

```text
Network I/O

G blocks
↓
Netpoller waits
↓
M continues
↓
P continues


File I/O

G blocks
↓
M blocks
↓
P detached
↓
Another M created/reused
```

---

# How File I/O Works

Examples:

```go
file.Read(...)
file.Write(...)
os.ReadFile(...)
os.WriteFile(...)
```

Typical execution flow:

```text
G1
 |
 v
file.Read()

 |
 v

M1 enters syscall

 |
 v

Kernel performs disk operation

 |
 v

M1 blocked
```

Since the OS thread is blocked:

```text
Runtime:
Detach P from M1
```

Diagram:

```text
Before

P1 → M1 → G1

After File Syscall

P1

M1 (Blocked in Kernel)

P1 → M2 → G2
```

Once the syscall completes:

```text
Kernel wakes M1

Runtime:
G1 becomes runnable again
```

---

# How Network I/O Works

Examples:

```go
conn.Read(...)
conn.Write(...)
http.Get(...)
http.Post(...)
```

Network sockets are configured as:

```text
Non-Blocking
```

Execution:

```text
G1 → conn.Read()
```

If data is unavailable:

```text
read() returns immediately
```

Instead of blocking:

```text
G1 parked
```

Runtime registers the socket with Netpoller.

```text
Socket FD
     |
     v
Netpoller
     |
     v
epoll/kqueue/IOCP
```

When data arrives:

```text
Kernel
  |
  v
Netpoller
  |
  v
G1 Runnable
```

---

# Go Netpoller Internals

Go implements an event-driven I/O subsystem called:

```text
Netpoller
```

Built on top of:

| OS        | Backend |
| --------- | ------- |
| Linux     | epoll   |
| macOS/BSD | kqueue  |
| Windows   | IOCP    |

Instead of blocking a thread:

```text
Wait here until data arrives
```

Go asks the OS:

```text
Notify me when this socket becomes ready
```

Workflow:

```text
Socket Not Ready
      |
      v
Park Goroutine
      |
      v
Register FD
      |
      v
epoll_wait()
      |
      v
Event Arrives
      |
      v
Goroutine Runnable
```

---

# Why Netpoller Works for Sockets but Not Files

A common interview question.

Why not use epoll for files?

Because sockets change readiness state.

Example:

```text
Socket

No Data
   ↓
Data Arrives
```

The kernel can notify the runtime.

Regular files are different.

Example:

```text
file.Read()
```

Most disk files are considered:

```text
Always Readable
```

The kernel sees:

```text
Ready
Ready
Ready
Ready
Ready
```

There is no meaningful readiness transition.

Therefore:

```text
epoll(file_fd)
```

provides little value.

This is why Go uses Netpoller for sockets but not for normal disk files.

---

# Scheduler Behavior During File I/O

When a goroutine enters a blocking syscall:

```go
file.Read(...)
```

The runtime executes:

```text
runtime.entersyscall()
```

Flow:

```text
G1
 |
 v
Syscall
 |
 v
M1 Blocked
 |
 v
Detach P
 |
 v
Assign P to M2
 |
 v
Continue Scheduling
```

Diagram:

```text
Before

P1 → M1 → G1

After

P1 → M2 → G2

M1 Blocked
```

This prevents the entire scheduler from stopping.

---

# Scheduler Behavior During Network I/O

Network I/O follows a different path.

```text
G1 → conn.Read()
```

If socket is not ready:

```text
G1 parked
```

Important:

```text
M is NOT blocked
P is NOT detached
```

Diagram:

```text
P1 → M1

G1 waiting

G2 running
G3 running
G4 running
```

When data arrives:

```text
Netpoller
    |
    v
G1 Runnable
```

The scheduler picks G1 again.

---

# File I/O Can Cause Thread Growth

Consider:

```go
for _, file := range files {
    go process(file)
}
```

Assume:

```text
100,000 files
```

Every goroutine executes:

```go
os.ReadFile(...)
```

Many threads may become blocked:

```text
M1 blocked
M2 blocked
M3 blocked
M4 blocked
...
```

The runtime keeps creating threads to maintain progress.

Result:

```text
More Threads
More Memory
More Context Switching
Higher Scheduler Overhead
```

---

# Why Go Can Handle Millions of Network Connections

Without Netpoller:

```text
1 Connection
=
1 Thread
```

This does not scale.

With Netpoller:

```text
1 Million Connections

↓

1 Million Goroutines

↓

Few OS Threads
```

Diagram:

```text
1,000,000 Goroutines

        |
        v

     Netpoller

        |
        v

   Few Threads
```

This is the primary reason Go is widely used for:

* API Gateways
* Reverse Proxies
* Load Balancers
* Message Brokers
* Streaming Systems

---

# Common Misconceptions

## Misconception 1

```text
Network I/O never blocks
```

Wrong.

The goroutine waits.

The difference is:

```text
Thread does not wait
```

---

## Misconception 2

```text
File I/O always blocks
```

Not entirely true.

Modern operating systems provide:

* io_uring
* AIO
* Async File APIs

However:

```text
Go runtime does not currently integrate
file operations into Netpoller
the same way it does sockets.
```

---

## Misconception 3

```text
epoll makes network faster
```

Wrong.

epoll improves:

```text
Scalability
```

not raw network speed.

---

# Production Considerations

For network-heavy services:

```text
Goroutines scale extremely well.
```

For file-heavy services:

Use:

* Worker Pools
* Bounded Concurrency
* Semaphores

Example:

```go
sem := make(chan struct{}, 100)

for _, file := range files {
    sem <- struct{}{}

    go func(f string) {
        defer func() {
            <-sem
        }()

        processFile(f)

    }(file)
}
```

Benefits:

```text
Bounded Threads
Controlled Memory Usage
Predictable Throughput
```

---

# Summary Table

| Aspect                 | File I/O    | Network I/O |
| ---------------------- | ----------- | ----------- |
| Uses Netpoller         | Usually No  | Yes         |
| Goroutine Suspended    | Yes         | Yes         |
| OS Thread Blocked      | Usually Yes | Usually No  |
| P Detached             | Yes         | No          |
| Thread Growth Possible | Yes         | Rare        |
| Scales to Millions     | Difficult   | Excellent   |
| Uses epoll/kqueue/IOCP | No          | Yes         |

---

# L5 Interview Questions and Answers

## Q1. Why does Go use Netpoller for sockets but not regular files?

Because sockets have readiness transitions:

```text
Not Ready
→
Ready
```

which epoll/kqueue can monitor.

Regular files are generally always readable from the kernel's perspective and do not generate useful readiness events.

---

## Q2. What happens when a goroutine executes file.Read()?

Flow:

```text
G enters syscall
M enters kernel
M blocks
Runtime detaches P
P assigned elsewhere
Scheduler continues
```

---

## Q3. What is runtime.entersyscall()?

It informs the scheduler that:

```text
This thread may block.
```

The runtime can then detach the processor (P) and keep the scheduler running.

---

## Q4. Why can Go support millions of idle connections?

Because waiting connections consume:

```text
Goroutines
```

instead of:

```text
OS Threads
```

Netpoller multiplexes many sockets onto a small thread pool.

---

## Q5. Can excessive file I/O cause thread explosion?

Yes.

Many blocked file syscalls may force the runtime to create additional threads to keep work progressing.

---

## Q6. Why is epoll ineffective for disk files?

Disk files are generally always reported as ready.

There is no readiness transition for epoll to monitor.

---

## Q7. What is the relationship between Netpoller and Scheduler?

Netpoller:

```text
Detects I/O readiness
```

Scheduler:

```text
Executes runnable goroutines
```

Netpoller feeds runnable goroutines back into the scheduler.

---

## Q8. How does Go keep making progress when a thread blocks in a syscall?

The runtime detaches:

```text
P from blocked M
```

and reassigns the P to another available thread.

---

## Q9. How do worker pools help with file I/O?

They limit concurrent blocking syscalls and prevent excessive thread creation.

---

## Q10. Compare scalability of File I/O vs Network I/O.

Network I/O scales primarily through:

```text
Netpoller + Goroutine Parking
```

File I/O scales through:

```text
Scheduler Detaching P From Blocked M
```

Network I/O therefore achieves much higher concurrency with fewer OS threads.
