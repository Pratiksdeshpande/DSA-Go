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

1. [Concurrency vs Parallelism](concurrency_vs_parallelism/README.md)
2. [Goroutines](Goroutines-in-Go/README.md)
3. [Go Scheduler Internals (G-M-P model)](Go-Scheduler-Internals-G-M-P-Model/README.md)
4. Channels Internals
5. Select Statement
6. Backpressure

---

## Phase 2 — Correctness & Failure Modes

7. Race Conditions
8. Deadlocks
9. Goroutine Leaks
10. Graceful Shutdown
11. Context Propagation & Cancellation

---

## Phase 3 — Production Concurrency Patterns

12. Worker Pool Pattern
13. Fan-Out / Fan-In
14. Pipelines
15. Semaphore Pattern
16. Bounded Concurrency

---

## Phase 4 — Runtime & Performance Internals

17. Escape Analysis
18. Stack vs Heap in Go
19. Memory Management
20. Garbage Collector Internals
21. Scheduler + GC Interaction
22. Allocation Optimization

---

## Phase 5 — L5-Level Advanced Topics

23. sync.Mutex Internals
24. sync.RWMutex Tradeoffs
25. sync.Cond
26. sync.Pool
27. Atomic Operations & Memory Ordering
28. False Sharing & Cache Lines
29. Profiling & Debugging Concurrent Programs
30. Production Failure Scenarios

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