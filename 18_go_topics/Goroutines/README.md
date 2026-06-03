# For a Google L5-level interview, the expectation is usually:

* You can write correct concurrent code quickly
* You understand runtime internals beyond syntax
* You can reason about production failures
* You understand performance tradeoffs
* You can explain *why* Go behaves a certain way
* You can identify hidden bugs from partial code
* You can discuss scheduler/runtime/GC interactions confidently

---

# The Roadmap (In Correct Order)

The order matters because each topic builds on runtime mental models from previous topics.

## Phase 1 — Core Concurrency Model

These are foundational.

01. [Concurrency vs Parallelism](concurrency_vs_parallelism/README.md) ✅
02. [Goroutines in Go](Goroutines-in-Go/README.md) ✅
03. [Go Scheduler Internals (G-M-P model)](Go-Scheduler-Internals-G-M-P-Model/README.md) ✅
04. [Complete Goroutine Lifecycle with G-M-P](Complete-Goroutine-Lifecycle-with-G-M-P/README.md) ✅
05. [File-IO-vs-Network-IO-in-Goroutines](File-IO-vs-Network-IO-in-Goroutines/README.md) ✅
06. [Semaphore in Go](Semaphore/README.md) ✅

---

## Phase 2 — Correctness & Failure Modes

07. [Channels Internals](Channels-Internals/README.md) ✅
08. [Select Statement](Go-Select-Statement/README.md) ✅
09. [Backpressure](Backpressure/README.md) ✅
10. [Race Conditions](Race-Conditions/README.md) ✅
11. [Deadlocks](Deadlocks/README.md) ✅
12. Goroutine Leaks
13. Graceful Shutdown
14. Context Propagation & Cancellation

---

## Phase 3 — Production Concurrency Patterns

15. Worker Pool Pattern
16. Fan-Out / Fan-In
17. Pipelines
18. Semaphore Pattern
19. Bounded Concurrency

---

## Phase 4 — Runtime & Performance Internals

20. Escape Analysis
21. Stack vs Heap in Go
22. Memory Management
23. Garbage Collector Internals
24. Scheduler + GC Interaction
25. Allocation Optimization

---

## Phase 5 — L5-Level Advanced Topics

26. sync.Mutex Internals
27. sync.RWMutex Tradeoffs
28. sync.Cond
29. sync.Pool
30. Atomic Operations & Memory Ordering
31. False Sharing & Cache Lines
32. Profiling & Debugging Concurrent Programs
33. Production Failure Scenarios

---

# How We Will Study

For every topic, we will go through:

1. Intuition
2. Internal Runtime Behavior
3. Real Production Use Cases
4. Common Interview Questions
5. Hidden Traps
6. Tradeoffs
7. Whiteboard-style reasoning
8. Production-grade code examples
9. Performance considerations
10. Senior-level follow-up questions

You should aim to be able to:

* explain the concept,
* implement it,
* debug it,
* optimize it,
* and discuss tradeoffs.

That is the actual L5 bar.

---
