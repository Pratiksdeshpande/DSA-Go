# Fibonacci Sequence

## Problem Statement

The **Fibonacci sequence** is a series of numbers where each number is the sum of the two preceding ones, starting from 0 and 1.

```
F(0) = 0
F(1) = 1
F(n) = F(n-1) + F(n-2) for n > 1
```

**Sequence:** 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...

Given an integer `n`, calculate `F(n)`.

---

## Input/Output Examples

### Example 1:
```
Input: n = 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1
```

### Example 2:
```
Input: n = 3
Output: 2
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2
```

### Example 3:
```
Input: n = 10
Output: 55
Explanation: F(10) = 55
```

### Example 4:
```
Input: n = 0
Output: 0
```

---

## Constraints

- `0 <= n <= 30`

---

## Solution Approaches

### Approach 1: Simple Recursion (Brute Force)

#### Idea
Directly translate the mathematical definition into code. For each `F(n)`, recursively compute `F(n-1)` and `F(n-2)`.

#### Code
```go
func fibWithRecursion(n int) int {
    if n <= 1 {
        return n
    }
    return fibWithRecursion(n-1) + fibWithRecursion(n-2)
}
```

#### Recursion Tree for F(5)
```
                         fib(5)
                        /      \
                   fib(4)      fib(3)
                  /     \      /     \
              fib(3)  fib(2) fib(2) fib(1)
              /   \    /  \   /  \
          fib(2) fib(1) ...  ...  ...
          /   \
      fib(1) fib(0)
```

#### Problem
- **Overlapping subproblems**: `fib(3)` is computed twice, `fib(2)` is computed three times, etc.
- Exponential time complexity!

#### Complexity
| Time | Space |
|------|-------|
| O(2ⁿ) | O(n) - recursion stack depth |

---

### Approach 2: Memoization (Top-Down DP)

#### Idea
Use recursion but **cache** the results of subproblems to avoid redundant calculations.

#### Code
```go
func fibWithDP(n int) int {
    cache := make(map[int]int, n)
    return dynamicProgramming(n, cache)
}

func dynamicProgramming(n int, cache map[int]int) int {
    // Base case
    if n <= 1 {
        return n
    }

    // Check if already computed
    if val, ok := cache[n]; ok {
        return val
    }

    // Compute and store in cache
    cache[n] = dynamicProgramming(n-1, cache) + dynamicProgramming(n-2, cache)

    return cache[n]
}
```

#### How It Works
```
fib(5) with memoization:

Call fib(5) → needs fib(4) and fib(3)
  Call fib(4) → needs fib(3) and fib(2)
    Call fib(3) → needs fib(2) and fib(1)
      Call fib(2) → needs fib(1) and fib(0)
        fib(1) = 1, fib(0) = 0
        cache[2] = 1
      fib(1) = 1
      cache[3] = 2
    fib(2) = cache[2] = 1  ← Cache hit!
    cache[4] = 3
  fib(3) = cache[3] = 2  ← Cache hit!
  cache[5] = 5
```

#### Complexity
| Time | Space |
|------|-------|
| O(n) | O(n) - cache + recursion stack |

---

### Approach 3: Tabulation (Bottom-Up DP)

#### Idea
Build the solution iteratively from the base cases up to `n`, storing results in an array.

#### Code
```go
func fibTabulation(n int) int {
    if n <= 1 {
        return n
    }

    dp := make([]int, n+1)
    dp[0] = 0
    dp[1] = 1

    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }

    return dp[n]
}
```

#### Visualization
```
n:    0   1   2   3   4   5   6   7   8   9   10
dp:   0   1   1   2   3   5   8  13  21  34   55
              ↑
          dp[2] = dp[1] + dp[0] = 1 + 0 = 1
```

#### Complexity
| Time | Space |
|------|-------|
| O(n) | O(n) - dp array |

---

### Approach 4: Space Optimized (Iterative)

#### Idea
Since we only need the last two values to compute the next one, we don't need to store the entire array.

#### Code
```go
func fibOptimized(n int) int {
    if n <= 1 {
        return n
    }

    prev2 := 0  // F(n-2)
    prev1 := 1  // F(n-1)

    for i := 2; i <= n; i++ {
        current := prev1 + prev2
        prev2 = prev1
        prev1 = current
    }

    return prev1
}
```

#### Visualization
```
i=2: prev2=0, prev1=1 → current=1 → prev2=1, prev1=1
i=3: prev2=1, prev1=1 → current=2 → prev2=1, prev1=2
i=4: prev2=1, prev1=2 → current=3 → prev2=2, prev1=3
i=5: prev2=2, prev1=3 → current=5 → prev2=3, prev1=5
...
```

#### Complexity
| Time | Space |
|------|-------|
| O(n) | O(1) ← Best! |

---

### Approach 5: Matrix Exponentiation (Advanced)

#### Idea
Use the matrix identity:
```
| F(n+1)  F(n)   |   | 1  1 |ⁿ
| F(n)    F(n-1) | = | 1  0 |
```

Use fast exponentiation to compute the matrix power in O(log n) time.

#### Code
```go
func fibMatrix(n int) int {
    if n <= 1 {
        return n
    }

    result := matrixPower([][]int{{1, 1}, {1, 0}}, n)
    return result[0][1]
}

func matrixPower(matrix [][]int, n int) [][]int {
    result := [][]int{{1, 0}, {0, 1}} // Identity matrix

    for n > 0 {
        if n%2 == 1 {
            result = multiplyMatrix(result, matrix)
        }
        matrix = multiplyMatrix(matrix, matrix)
        n /= 2
    }

    return result
}

func multiplyMatrix(a, b [][]int) [][]int {
    return [][]int{
        {a[0][0]*b[0][0] + a[0][1]*b[1][0], a[0][0]*b[0][1] + a[0][1]*b[1][1]},
        {a[1][0]*b[0][0] + a[1][1]*b[1][0], a[1][0]*b[0][1] + a[1][1]*b[1][1]},
    }
}
```

#### Complexity
| Time | Space |
|------|-------|
| O(log n) ← Best! | O(1) |

---

## Complexity Comparison

| Approach | Time Complexity | Space Complexity | Notes |
|----------|----------------|------------------|-------|
| Recursion (Brute Force) | O(2ⁿ) | O(n) | Too slow for large n |
| Memoization (Top-Down) | O(n) | O(n) | Easy to implement |
| Tabulation (Bottom-Up) | O(n) | O(n) | No recursion overhead |
| Space Optimized | O(n) | O(1) | Best for interviews |
| Matrix Exponentiation | O(log n) | O(1) | Best theoretical |

---

## When to Use Which Approach?

| Scenario | Recommended Approach |
|----------|---------------------|
| Quick implementation | Memoization |
| Interview (balance of clarity & efficiency) | Space Optimized O(1) |
| Very large n (millions) | Matrix Exponentiation |
| Teaching/Learning DP | Tabulation |

---

## Key Concepts for Interviews

### 1. **Understanding DP Patterns**
Fibonacci is the classic example to understand:
- **Overlapping Subproblems**: Same subproblems are solved multiple times
- **Optimal Substructure**: Solution depends on solutions of smaller subproblems

### 2. **Top-Down vs Bottom-Up**
| Top-Down (Memoization) | Bottom-Up (Tabulation) |
|------------------------|------------------------|
| Starts from F(n) → F(0) | Starts from F(0) → F(n) |
| Uses recursion + cache | Uses iteration + array |
| Only computes needed values | Computes all values |
| Risk of stack overflow | No stack overflow |

### 3. **Space Optimization Pattern**
When the current state only depends on a fixed number of previous states:
- Don't store entire DP table
- Only keep track of required previous values
- Common in 1D DP problems

### 4. **Common Follow-up Questions**

1. **"What if n can be negative?"**
   - Use the negafibonacci formula: F(-n) = (-1)^(n+1) × F(n)

2. **"What if n is very large (like 10^18)?"**
   - Use Matrix Exponentiation O(log n)

3. **"What about integer overflow?"**
   - Use modular arithmetic (return result % MOD)
   - Or use big integers

4. **"Can you compute F(n) mod M?"**
   - Yes, apply mod at each step: `(a + b) % M = ((a % M) + (b % M)) % M`

---

## Related Problems

| Problem | Similarity |
|---------|------------|
| Climbing Stairs | F(n) = F(n-1) + F(n-2) |
| House Robber | 1D DP with recurrence |
| Tribonacci | F(n) = F(n-1) + F(n-2) + F(n-3) |
| Min Cost Climbing Stairs | DP with optimization |
| Decode Ways | Counting problems |

---

## Code Implementation (Go)

```go
package main

import "fmt"

func main() {
    // Test all approaches
    n := 10
    
    fmt.Println("Recursion:", fibWithRecursion(n))    // O(2^n) time
    fmt.Println("Memoization:", fibWithDP(n))         // O(n) time, O(n) space
    fmt.Println("Tabulation:", fibTabulation(n))      // O(n) time, O(n) space
    fmt.Println("Optimized:", fibOptimized(n))        // O(n) time, O(1) space
}

// Approach 1: Simple Recursion - O(2^n) time, O(n) space
func fibWithRecursion(n int) int {
    if n <= 1 {
        return n
    }
    return fibWithRecursion(n-1) + fibWithRecursion(n-2)
}

// Approach 2: Memoization (Top-Down DP) - O(n) time, O(n) space
func fibWithDP(n int) int {
    cache := make(map[int]int, n)
    return dynamicProgramming(n, cache)
}

func dynamicProgramming(n int, cache map[int]int) int {
    if n <= 1 {
        return n
    }
    if val, ok := cache[n]; ok {
        return val
    }
    cache[n] = dynamicProgramming(n-1, cache) + dynamicProgramming(n-2, cache)
    return cache[n]
}

// Approach 3: Tabulation (Bottom-Up DP) - O(n) time, O(n) space
func fibTabulation(n int) int {
    if n <= 1 {
        return n
    }
    dp := make([]int, n+1)
    dp[0], dp[1] = 0, 1
    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}

// Approach 4: Space Optimized - O(n) time, O(1) space
func fibOptimized(n int) int {
    if n <= 1 {
        return n
    }
    prev2, prev1 := 0, 1
    for i := 2; i <= n; i++ {
        prev2, prev1 = prev1, prev1+prev2
    }
    return prev1
}
```

---

## LeetCode Reference

- **Problem**: [509. Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)
- **Difficulty**: Easy
- **Tags**: Math, Dynamic Programming, Recursion, Memoization

