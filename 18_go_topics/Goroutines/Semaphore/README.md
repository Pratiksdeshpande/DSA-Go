# Semaphore in Go

## Definition

A **Semaphore** is a concurrency control mechanism used to limit the number of goroutines accessing a resource or executing a task simultaneously.

In Go, a semaphore is commonly implemented using a **buffered channel**.

---

## Example

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

---

## Semaphore Logic

```go
semaphore := make(chan struct{}, 10)
```

* Creates a buffered channel with capacity `10`
* Maximum `10` goroutines can run concurrently

### Acquire Slot

```go
semaphore <- struct{}{}
```

* Adds a token into the channel
* Blocks if channel is full

### Release Slot

```go
<-semaphore
```

* Removes token from the channel
* Frees space for another goroutine

This ensures controlled concurrency.

---

## Why Semaphore is Required

In production systems, processing too many files concurrently can overload the system.

### Example Use Case: File Handling

Suppose an application processes thousands of files:

* Reading large files
* Uploading files to cloud storage
* Parsing CSV/JSON logs
* Generating reports

If 1000 goroutines start together:

* Memory usage can spike
* Disk I/O becomes overloaded
* Too many open file descriptors may occur
* Application performance degrades

Using a semaphore limits concurrent file processing, making the system stable and resource-efficient.

Example:

```go
semaphore := make(chan struct{}, 10)
```

Only 10 files are processed at the same time.
