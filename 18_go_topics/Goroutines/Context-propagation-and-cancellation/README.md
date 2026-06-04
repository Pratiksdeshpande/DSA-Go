# Context Propagation & Cancellation in Go

> A deep dive into Go's Context package, cancellation propagation, deadlines, timeouts, goroutine lifecycle management, and production-grade distributed systems patterns.

---

# Table of Contents

* [1. Introduction](#1-introduction)
* [2. Why Context Exists](#2-why-context-exists)
* [3. What is Context?](#3-what-is-context)
* [4. Context Tree Model](#4-context-tree-model)
* [5. Context Interface Internals](#5-context-interface-internals)
* [6. How Cancellation Works Internally](#6-how-cancellation-works-internally)
* [7. Why Context Uses Channel Close](#7-why-context-uses-channel-close)
* [8. Types of Context](#8-types-of-context)
* [9. Context Propagation](#9-context-propagation)
* [10. Real Production Use Cases](#10-real-production-use-cases)
* [11. Context and Goroutine Leak Prevention](#11-context-and-goroutine-leak-prevention)
* [12. Context and Graceful Shutdown](#12-context-and-graceful-shutdown)
* [13. Common Context Patterns](#13-common-context-patterns)
* [14. Hidden Traps and Mistakes](#14-hidden-traps-and-mistakes)
* [15. Performance Considerations](#15-performance-considerations)
* [16. Whiteboard Interview Reasoning](#16-whiteboard-interview-reasoning)
* [17. L5 Interview Questions and Answers](#17-l5-interview-questions-and-answers)
* [18. Key Takeaways](#18-key-takeaways)

---

# 1. Introduction

Modern backend systems perform work across multiple layers:

```text
HTTP Request
      |
      v
API Service
      |
      +---- Database
      |
      +---- Cache
      |
      +---- External Service
      |
      +---- Background Workers
```

When a request is cancelled or times out, all downstream operations should stop.

Without a cancellation mechanism:

* resources are wasted
* goroutines continue running
* database queries continue executing
* external requests continue unnecessarily

Go solves this problem using the `context` package.

---

# 2. Why Context Exists

Imagine:

```text
User Request Started
       |
       v
Database Query
       |
       v
External API Call
```

After 200ms:

```text
User closes browser
```

Question:

Should all operations continue?

Answer:

No.

The system should notify all participating operations that the request is no longer relevant.

This is the purpose of Context.

---

# 3. What is Context?

Context is a mechanism for propagating:

* cancellation signals
* deadlines
* timeouts
* request-scoped metadata

across API boundaries and goroutines.

---

## Senior-Level Definition

> Context is a hierarchical cancellation mechanism used to coordinate the lifecycle of concurrent operations.

---

# 4. Context Tree Model

Contexts form a parent-child hierarchy.

```text
Root Context
      |
      +------ Child Context A
      |
      +------ Child Context B
      |
      +------ Child Context C
```

If the parent is cancelled:

```text
Root Context Cancelled
          |
          +------ Child A Cancelled
          |
          +------ Child B Cancelled
          |
          +------ Child C Cancelled
```

Cancellation automatically propagates downward.

---

# 5. Context Interface Internals

Internally, Context is defined as:

```go
type Context interface {
    Deadline() (time.Time, bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

---

## Most Important Method

```go
Done() <-chan struct{}
```

This channel is used to notify goroutines that cancellation has occurred.

---

# 6. How Cancellation Works Internally

Example:

```go
ctx, cancel := context.WithCancel(
    context.Background(),
)

cancel()
```

Conceptually:

```go
close(doneChannel)
```

is executed.

---

## Before Cancellation

```text
Worker 1
Worker 2
Worker 3

      waiting on

      ctx.Done()
```

---

## After Cancellation

```text
close(ctx.Done())
```

All waiting goroutines become aware of cancellation.

---

# 7. Why Context Uses Channel Close

A critical interview topic.

Many developers assume Context sends a value:

```go
done <- struct{}{}
```

That would wake only one receiver.

Instead Context closes the channel:

```go
close(done)
```

---

## Channel Close as Broadcast

```text
Worker A
Worker B
Worker C

waiting on

<-ctx.Done()
```

After:

```go
close(done)
```

All workers are notified.

---

## Important Distinction

Cancellation notification is immediate.

Execution is not.

```text
Waiting
   ↓
Runnable
```

The scheduler still decides when goroutines run.

---

# 8. Types of Context

---

## context.Background()

Root context.

Used when starting an application.

```go
ctx := context.Background()
```

---

## context.TODO()

Placeholder context.

```go
ctx := context.TODO()
```

Used when the proper context is not yet known.

---

## context.WithCancel()

Manual cancellation.

```go
ctx, cancel :=
    context.WithCancel(
        context.Background(),
    )
```

---

## context.WithTimeout()

Automatic cancellation after duration.

```go
ctx, cancel :=
    context.WithTimeout(
        context.Background(),
        5*time.Second,
    )
```

---

## context.WithDeadline()

Automatic cancellation at a specific timestamp.

```go
ctx, cancel :=
    context.WithDeadline(
        context.Background(),
        deadline,
    )
```

---

## context.WithValue()

Adds request-scoped metadata.

```go
ctx = context.WithValue(
    ctx,
    requestIDKey,
    requestID,
)
```

---

# 9. Context Propagation

Every function should pass Context downstream.

```go
func Handler(
    ctx context.Context,
) error {

    return service.Process(ctx)
}
```

```go
func Process(
    ctx context.Context,
) error {

    return repository.Save(ctx)
}
```

Propagation maintains:

* cancellation
* deadlines
* metadata

throughout the request lifecycle.

---

# 10. Real Production Use Cases

---

## HTTP Servers

```go
func handler(
    w http.ResponseWriter,
    r *http.Request,
) {
    ctx := r.Context()
}
```

If the client disconnects:

```text
Context automatically cancels
```

---

## Database Queries

```go
db.QueryContext(
    ctx,
    query,
)
```

Database work can stop when requests timeout.

---

## gRPC

Deadlines automatically propagate across services.

```go
ctx, cancel :=
    context.WithTimeout(...)
```

---

## HTTP Clients

```go
req = req.WithContext(ctx)
```

Outbound requests inherit cancellation.

---

## Worker Pools

Workers stop cleanly.

```go
select {
case <-ctx.Done():
    return
}
```

---

# 11. Context and Goroutine Leak Prevention

Without Context:

```go
go func() {
    <-someChannel
}()
```

Potential issue:

```text
Channel never receives
```

Goroutine waits forever.

---

Using Context:

```go
go func() {

    select {

    case <-someChannel:

    case <-ctx.Done():
        return
    }
}()
```

Now the goroutine can exit.

---

# 12. Context and Graceful Shutdown

During service shutdown:

```go
ctx, cancel :=
    context.WithTimeout(
        context.Background(),
        30*time.Second,
    )
```

Shutdown signal propagates.

```text
HTTP Handlers
      ↓
Workers
      ↓
Database Operations
      ↓
External Calls
```

All components receive cancellation.

---

# 13. Common Context Patterns

---

## Pattern 1: Timeout Protection

```go
ctx, cancel :=
    context.WithTimeout(
        parent,
        2*time.Second,
    )

defer cancel()
```

---

## Pattern 2: Worker Cancellation

```go
select {

case <-ctx.Done():
    return

case job := <-jobs:
    process(job)
}
```

---

## Pattern 3: Periodic Cancellation Checks

```go
for _, item := range items {

    select {

    case <-ctx.Done():
        return ctx.Err()

    default:
    }

    process(item)
}
```

---

# 14. Hidden Traps and Mistakes

---

## Mistake 1

Ignoring cancellation.

Bad:

```go
for {
    process()
}
```

---

Correct:

```go
for {

    select {

    case <-ctx.Done():
        return

    default:
        process()
    }
}
```

---

## Mistake 2

Forgetting cancel()

Bad:

```go
ctx, cancel :=
    context.WithTimeout(...)
```

Missing:

```go
cancel()
```

---

Correct:

```go
defer cancel()
```

---

## Mistake 3

Storing Context in Structs

Bad:

```go
type Service struct {
    ctx context.Context
}
```

Context should be request-scoped.

---

## Mistake 4

Using Context as a Parameter Bag

Bad:

```go
ctx = context.WithValue(
    ctx,
    "config",
    hugeConfig,
)
```

Use Context values only for request-scoped metadata.

---

# 15. Performance Considerations

---

## Context Creation

Cheap.

```go
context.WithCancel(...)
```

creates lightweight objects.

---

## Cancellation

Efficient.

Implemented using:

```go
close(doneChannel)
```

which acts as a broadcast signal.

---

## Large Context Trees

```text
Parent
  |
  +---- 100,000 Children
```

Cancellation must propagate to all descendants.

Usually acceptable but important to understand at scale.

---

## Avoid Large Values

Bad:

```go
ctx = context.WithValue(
    ctx,
    key,
    hugeObject,
)
```

Can increase memory retention.

---

# 16. Whiteboard Interview Reasoning

Request Flow:

```text
Client Request
      |
      +---- Database
      |
      +---- Cache
      |
      +---- Payment Service
```

Client disconnects.

Without Context:

```text
Database continues

Cache continues

Payment continues
```

Resources wasted.

---

With Context:

```text
Request Cancelled
        |
        +---- Database Cancelled
        |
        +---- Cache Cancelled
        |
        +---- Payment Cancelled
```

All operations receive notification.

---

# 17. L5 Interview Questions and Answers

---

## Q1. What problem does Context solve?

Context provides:

* cancellation propagation
* deadlines
* timeouts
* request-scoped metadata

across API boundaries.

---

## Q2. Why is Context passed as the first parameter?

Convention.

Makes request lifecycle management visible.

```go
func Process(
    ctx context.Context,
)
```

---

## Q3. Why does Context use channel close instead of sending a value?

Sending:

```go
done <- struct{}{}
```

wakes one receiver.

Closing:

```go
close(done)
```

notifies all receivers.

---

## Q4. Is Context cancellation immediate?

No.

Cancellation makes waiting goroutines runnable.

The scheduler still determines when they execute.

---

## Q5. Why is Context important for preventing goroutine leaks?

Blocked goroutines can receive cancellation signals and exit cleanly.

---

## Q6. Difference between WithCancel and WithTimeout?

### WithCancel

Manual cancellation.

```go
cancel()
```

---

### WithTimeout

Automatic cancellation after duration.

```go
2 seconds
↓
automatic cancellation
```

---

## Q7. Why should Context not be stored in structs?

Contexts are request-scoped and short-lived.

Storing them in long-lived objects causes lifecycle issues.

---

## Q8. What is the most important part of Context internals?

The Done channel.

```go
Done() <-chan struct{}
```

It is the foundation of cancellation propagation.

---

# 18. Key Takeaways

* Context controls request lifecycles.
* Context is primarily a cancellation propagation mechanism.
* Context forms a parent-child hierarchy.
* Cancellation propagates automatically to descendants.
* Internally, Context relies on a Done channel.
* Cancellation is implemented using channel close.
* Context prevents goroutine leaks.
* Context enables graceful shutdown.
* Always pass Context downstream.
* Always call cancel() when creating cancellable contexts.
* Avoid storing Context in structs.
* Avoid abusing Context values.

---

# Final Mental Model

Think of Context as:

```text
Request Lifecycle Controller
```

not:

```text
Timeout Utility
```

A senior Go engineer views Context as the mechanism that coordinates cancellation, deadlines, cleanup, and resource management across concurrent operations and distributed systems.
