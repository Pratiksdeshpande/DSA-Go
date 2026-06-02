# Print Numbers from 1 to N Sequentially Using Two Goroutines

## Problem Statement

Given an integer `n`, print numbers from `1` to `n` in sequential order using **two goroutines**:

* One goroutine is responsible for printing **odd numbers**.
* Another goroutine is responsible for printing **even numbers**.

The output must always be printed in the correct sequence.

### Example

#### Input

```text
n = 10
```

#### Output

```text
1
2
3
4
5
6
7
8
9
10
```

### Constraints

* Only two goroutines should be used:

  * Odd Number Goroutine
  * Even Number Goroutine
* Output must be sequential.
* Goroutines must coordinate with each other to maintain ordering.

---

# Solution Approach

The challenge is not printing odd and even numbers separately; the challenge is ensuring they are printed in the correct order.

If both goroutines execute independently, the Go scheduler may run them in any order, producing outputs such as:

```text
1 3 5 7 9 2 4 6 8 10
```

or

```text
2 4 6 8 10 1 3 5 7 9
```

which does not satisfy the requirement.

To solve this problem, we need a synchronization mechanism that allows the goroutines to take turns.

## Using Channels for Synchronization

We use two channels:

```go
oddTurn
evenTurn
```

### Responsibilities

* `oddTurn` gives permission to the odd goroutine to print.
* `evenTurn` gives permission to the even goroutine to print.

### Execution Flow

1. Main function signals the odd goroutine first.
2. Odd goroutine prints an odd number.
3. Odd goroutine signals the even goroutine.
4. Even goroutine prints an even number.
5. Even goroutine signals the odd goroutine.
6. This process continues until all numbers are printed.

The communication pattern becomes:

```text
Odd -> Even -> Odd -> Even -> Odd -> Even
```

This guarantees sequential output.

---

# Solution Code

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 10

	oddTurn := make(chan struct{})
	evenTurn := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	// Odd Number Goroutine
	go func() {
		defer wg.Done()

		for i := 1; i <= n; i += 2 {
			<-oddTurn

			fmt.Println(i)

			if i+1 <= n {
				evenTurn <- struct{}{}
			}
		}
	}()

	// Even Number Goroutine
	go func() {
		defer wg.Done()

		for i := 2; i <= n; i += 2 {
			<-evenTurn

			fmt.Println(i)

			if i+1 <= n {
				oddTurn <- struct{}{}
			}
		}
	}()

	// Start execution with odd goroutine
	oddTurn <- struct{}{}

	wg.Wait()
}
```

---

# Explanation

## Step 1: Create Synchronization Channels

```go
oddTurn := make(chan struct{})
evenTurn := make(chan struct{})
```

These channels act as signals between goroutines.

At any point, only the goroutine that receives a signal can continue execution.

---

## Step 2: Start Odd Goroutine

```go
go func() {
	for i := 1; i <= n; i += 2 {
		<-oddTurn
		fmt.Println(i)

		if i+1 <= n {
			evenTurn <- struct{}{}
		}
	}
}()
```

### What Happens?

For every odd number:

1. Wait for permission from `oddTurn`.
2. Print the current odd number.
3. Signal the even goroutine.

Example:

```text
Receive signal
Print 1
Signal even goroutine
```

---

## Step 3: Start Even Goroutine

```go
go func() {
	for i := 2; i <= n; i += 2 {
		<-evenTurn
		fmt.Println(i)

		if i+1 <= n {
			oddTurn <- struct{}{}
		}
	}
}()
```

### What Happens?

For every even number:

1. Wait for permission from `evenTurn`.
2. Print the current even number.
3. Signal the odd goroutine.

Example:

```text
Receive signal
Print 2
Signal odd goroutine
```

---

## Step 4: Kick Off Execution

```go
oddTurn <- struct{}{}
```

Initially, both goroutines are blocked waiting for signals.

This statement allows the odd goroutine to start first.

Execution begins as:

```text
Main
  |
  v
Odd Goroutine
```

---

## Dry Run

Assume:

```text
n = 6
```

### Initial State

```text
oddTurn <- signal
```

### Iteration 1

Odd Goroutine:

```text
Receives signal
Prints 1
Signals even goroutine
```

Output:

```text
1
```

---

### Iteration 2

Even Goroutine:

```text
Receives signal
Prints 2
Signals odd goroutine
```

Output:

```text
1
2
```

---

### Iteration 3

Odd Goroutine:

```text
Receives signal
Prints 3
Signals even goroutine
```

Output:

```text
1
2
3
```

---

### Iteration 4

Even Goroutine:

```text
Receives signal
Prints 4
Signals odd goroutine
```

Output:

```text
1
2
3
4
```

---

### Iteration 5

Odd Goroutine:

```text
Receives signal
Prints 5
Signals even goroutine
```

Output:

```text
1
2
3
4
5
```

---

### Iteration 6

Even Goroutine:

```text
Receives signal
Prints 6
No further signal required
```

Final Output:

```text
1
2
3
4
5
6
```

---

# Time Complexity

Each number is printed exactly once.

```text
Time Complexity: O(n)
```

---

# Space Complexity

Two channels and one wait group are used.

```text
Space Complexity: O(1)
```

---

# Key Learnings

* Goroutines execute concurrently and do not guarantee execution order.
* Channels can be used as synchronization primitives.
* Unbuffered channels naturally block until a sender and receiver are ready.
* Coordinating goroutines through channels helps maintain deterministic execution.
* This problem is a common Go concurrency interview question used to evaluate understanding of goroutines, channels, and synchronization.
