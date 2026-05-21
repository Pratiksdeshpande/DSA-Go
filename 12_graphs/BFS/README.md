# Breadth-First Search (BFS) in Graphs using Go

---

# Table of Contents

1.  [Introduction to BFS](#1-introduction-to-bfs)
2.  [What is a Graph?](#2-what-is-a-graph)
3.  [BFS Definition](#3-bfs-definition)
4.  [Core Characteristics of BFS](#4-core-characteristics-of-bfs)
5.  [BFS Traversal Visualization](#5-bfs-traversal-visualization)
6.  [Key Terms You Must Know](#6-key-terms-you-must-know)
7.  [BFS Internal Working](#7-bfs-internal-working)
8.  [Queue in BFS](#8-queue-in-bfs)
9.  [Visited Tracking](#9-visited-tracking)
10. [BFS Algorithm](#10-bfs-algorithm)
11. [BFS Pseudocode](#11-bfs-pseudocode)
12. [BFS Implementation in Go](#12-bfs-implementation-in-go)
13. [Dry Run (Step-by-Step)](#13-dry-run-step-by-step)
14. [BFS Traversal Order](#14-bfs-traversal-order)
15. [BFS Time Complexity](#15-bfs-time-complexity)
16. [BFS Space Complexity](#16-bfs-space-complexity)
17. [BFS vs DFS](#17-bfs-vs-dfs)
18. [BFS in Directed vs Undirected Graphs](#18-bfs-in-directed-vs-undirected-graphs)
19. [BFS in Disconnected Graphs](#19-bfs-in-disconnected-graphs)
20. [BFS on Matrix/Grid](#20-bfs-on-matrixgrid)
21. [Multi-Source BFS](#21-multi-source-bfs)
22. [BFS Shortest Path Concept](#22-bfs-shortest-path-concept)
23. [Common BFS Patterns](#23-common-bfs-patterns)
24. [Common Mistakes](#24-common-mistakes)
25. [Senior-Level BFS Understanding](#25-senior-level-bfs-understanding)
26. [BFS Mental Model](#26-bfs-mental-model)
27. [Quick Revision Notes](#27-quick-revision-notes)
28. [Final Summary](#key-properties)

---

## 1. Introduction to BFS

Breadth-First Search (BFS) is one of the most fundamental graph traversal algorithms.

It explores a graph:

- Level by level
- Neighbor by neighbor
- Closest nodes first

BFS is extremely important because many advanced graph algorithms are built on top of it.

---

## 2. What is a Graph?

A graph is a data structure consisting of:

- Vertices (Nodes)
- Edges (Connections)

Example:

```text
    A
   / \
  B   C
 / \
D   E
```

### Vertices

```text
A, B, C, D, E
```

### Edges

```text
A-B
A-C
B-D
B-E
```

---

## 3. BFS Definition

### Definition

Breadth-First Search (BFS) is a graph traversal algorithm that explores nodes level-by-level starting from a source node using a queue data structure.

---

## 4. Core Characteristics of BFS

| Property | Description |
|---|---|
| Traversal Style | Level Order |
| Data Structure Used | Queue |
| Shortest Path | Yes (Unweighted Graph) |
| Traversal Nature | Iterative |
| Time Complexity | O(V + E) |
| Space Complexity | O(V) |

---

## 5. BFS Traversal Visualization

Consider:

```text
          0
        /   \
       1     2
      / \     \
     3   4     5
```

---

### BFS Traversal Order

Starting from `0`

### Level 0

```text
0
```

### Level 1

```text
1 2
```

### Level 2

```text
3 4 5
```

---

### Final BFS Order

```text
0 → 1 → 2 → 3 → 4 → 5
```

---

## 6. Key Terms You Must Know

---

### Vertex (Node)

A single entity in graph.

Example:

```text
0,1,2,3
```

---

### Edge

Connection between two nodes.

Example:

```text
0 → 1
```

---

### Neighbor / Adjacent Node

Directly connected node.

Example:

```text
0 -> 1,2
```

Here:
- 1 is neighbor of 0
- 2 is neighbor of 0

---

### Visited Node

A node already processed/traversed.

---

### Traversal

Process of visiting all graph nodes.

---

### Source Node

Starting node for BFS.

---

### Queue

FIFO data structure used in BFS.

FIFO:

```text
First In First Out
```

---

## 7. BFS Internal Working

BFS follows this pattern:

```text
1. Start from source node
2. Visit source node
3. Add source to queue
4. Remove node from queue
5. Visit all unvisited neighbors
6. Add neighbors to queue
7. Repeat until queue becomes empty
```

---

## 8. Queue in BFS

Queue is the backbone of BFS.

---

### Queue Visualization

```text
Front -> [1 2 3] <- Rear
```

Removal happens from front.

Insertion happens at rear.

---

### Why Queue?

Queue guarantees:

```text
Older nodes processed first
```

This naturally creates:

```text
Level-order traversal
```

---

## 9. Visited Tracking

Without visited tracking:

```text
0 <-> 1
```

Traversal becomes:

```text
0 → 1 → 0 → 1 → 0 ...
```

Infinite loop.

---

### Solution

Use:

```go

visited := make(map[int]bool)

```

---

## 10. BFS Algorithm

---

### High-Level Algorithm

```text
1. Create queue
2. Mark source visited
3. Push source into queue

4. While queue not empty:
    a. Pop front node
    b. Process node
    c. Visit all unvisited neighbors
    d. Push neighbors into queue
```

---

## 11. BFS Pseudocode

```text
BFS(graph, start):

    create visited set
    create queue

    mark start visited
    push start into queue

    while queue not empty:

        node = pop front

        process node

        for neighbor in graph[node]:

            if neighbor not visited:

                mark neighbor visited
                push neighbor into queue
```

---

## 12. BFS Implementation in Go

---

### Graph Representation (Adjacency List)

```go
graph := map[int][]int{
    0: {1, 2},
    1: {0, 3, 4},
    2: {0, 5},
    3: {1},
    4: {1},
    5: {2},
}
```

---

### Complete BFS Code

```go
package main

import "fmt"

func BFS(graph map[int][]int, start int) {

    visited := make(map[int]bool)

    queue := []int{}

    visited[start] = true
    queue = append(queue, start)

    for len(queue) > 0 {

        node := queue[0]
        queue = queue[1:]

        fmt.Print(node, " ")

        for _, neighbor := range graph[node] {

            if !visited[neighbor] {

                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
}

func main() {

    graph := map[int][]int{
        0: {1, 2},
        1: {0, 3, 4},
        2: {0, 5},
        3: {1},
        4: {1},
        5: {2},
    }

    BFS(graph, 0)
}
```

---

## 13. Dry Run (Step-by-Step)

---

### Initial State

```text
Queue   = [0]
Visited = {0}
```

---

### Iteration 1

Pop:

```text
0
```

Visit neighbors:

```text
1,2
```

Queue:

```text
[1,2]
```

Visited:

```text
0,1,2
```

---

### Iteration 2

Pop:

```text
1
```

Visit neighbors:

```text
3,4
```

Queue:

```text
[2,3,4]
```

---

### Iteration 3

Pop:

```text
2
```

Visit neighbor:

```text
5
```

Queue:

```text
[3,4,5]
```

---

### Final Traversal

```text
0 1 2 3 4 5
```

---

## 14. BFS Traversal Order

BFS always traverses:

```text
Nearest nodes first
```

This is why BFS is ideal for:

```text
Shortest path in unweighted graph
```

---

## 15. BFS Time Complexity

### Complexity

```text
O(V + E)
```

Where:

| Symbol | Meaning |
|---|---|
| V | Number of Vertices |
| E | Number of Edges |

---

### Why?

Each:
- Node visited once
- Edge checked once

---

## 16. BFS Space Complexity

### Complexity

```text
O(V)
```

---

### Reason

Extra memory used for:

- Queue
- Visited map

---

## 17. BFS vs DFS

| BFS | DFS |
|---|---|
| Queue | Stack/Recursion |
| Level-order | Depth-first |
| Shortest path | Not guaranteed |
| More memory | Less memory sometimes |
| Better for shortest distance | Better for exhaustive traversal |

---

## 18. BFS in Directed vs Undirected Graphs

---

### Undirected Graph

```text
0 <-> 1
```

Adjacency list:

```go
0: {1}
1: {0}
```

---

### Directed Graph

```text
0 -> 1
```

Adjacency list:

```go
0: {1}
1: {}
```

BFS works for both.

---

## 19. BFS in Disconnected Graphs

---

### Problem

```text
0 - 1

2 - 3
```

Starting BFS from `0` misses:

```text
2,3
```

---

### Solution

Run BFS from every unvisited node.

---

### Code Pattern

```go
for node := range graph {

    if !visited[node] {
        BFS(graph, node)
    }
}
```

---

## 20. BFS on Matrix/Grid

Very common in interviews.

---

### Grid Example

```text
0 0 0
0 1 0
0 0 0
```

---

### Possible Directions

```text
Up
Down
Left
Right
```

---

### Direction Array Pattern

```go
directions := [][]int{
    {-1, 0},
    {1, 0},
    {0, -1},
    {0, 1},
}
```

---

## 21. Multi-Source BFS

Normal BFS starts from:

```text
1 source node
```

Multi-source BFS starts from:

```text
Multiple nodes simultaneously
```

---

### Use Cases

- Rotten oranges
- Fire spread
- Nearest hospital
- Distance to nearest zero

---

## 22. BFS Shortest Path Concept

BFS guarantees shortest path in:

```text
Unweighted Graphs
```

---

### Why?

Because BFS explores:

```text
Closest distance first
```

---

### Distance Visualization

```text
Level 0 -> Distance 0
Level 1 -> Distance 1
Level 2 -> Distance 2
```

---

## 23. Common BFS Patterns

---

### Standard BFS

```text
Queue + Visited
```

---

### Level BFS

Process level by level.

Useful in:
- Trees
- Minimum steps problems

---

### Multi-Source BFS

Multiple starting points.

---

### BFS with Distance Array

Track shortest distance.

---

### BFS on Grid

Matrix traversal.

---

## 24. Common Mistakes

| Mistake | Problem |
|---|---|
| Forgetting visited map | Infinite loop |
| Using stack instead of queue | Converts to DFS |
| Marking visited too late | Duplicate insertions |
| Incorrect queue pop | Wrong traversal |
| Missing disconnected components | Incomplete traversal |

---

## 25. Senior-Level BFS Understanding

A senior engineer is expected to understand:

---

### 1. Why BFS Guarantees Shortest Path

Because traversal happens level-wise.

---

### 2. Queue Mechanics

How FIFO ordering impacts traversal.

---

### 3. Graph Representation Tradeoffs

#### Adjacency List

```text
Space Efficient
```

#### Adjacency Matrix

```text
Faster edge lookup
```

---

### 4. BFS Variants

- Multi-source BFS
- Bidirectional BFS
- 0-1 BFS

---

### 5. Complexity Analysis

Understanding:

```text
Why BFS is O(V + E)
```

not:

```text
O(V²)
```

---

### 6. BFS Memory Behavior

Queue may grow significantly in wide graphs.

---

### 7. BFS vs DFS Decision Making

Choosing correct traversal strategy.

---

### 8. BFS in Real Systems

Applications:
- Network routing
- Social graph traversal
- Distributed systems
- Service dependency traversal

---

## 26. BFS Mental Model

The best mental model:

```text
Wave propagation
```

Like water spreading outward.

---

### Visualization

```text
Level 0 -> Source
Level 1 -> Immediate neighbors
Level 2 -> Neighbors of neighbors
```

---

## 27. Quick Revision Notes

---

### BFS Core Formula

```text
Queue + Visited = BFS
```

---

### Key Properties

| Topic | Value |
|---|---|
| Data Structure | Queue |
| Traversal Style | Level-order |
| Shortest Path | Yes (Unweighted) |
| Time Complexity | O(V + E) |
| Space Complexity | O(V) |

---

### BFS Template

```go
func BFS(graph map[int][]int, start int) {

    visited := make(map[int]bool)

    queue := []int{start}

    visited[start] = true

    for len(queue) > 0 {

        node := queue[0]
        queue = queue[1:]

        for _, neighbor := range graph[node] {

            if !visited[neighbor] {

                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
}
```

---

## Final Summary

Breadth-First Search (BFS) is:

- A level-order graph traversal algorithm
- Based on queue data structure
- Essential for shortest path problems in unweighted graphs
- Foundational for advanced graph algorithms

Mastering BFS is critical for:

- DSA interviews
- Competitive programming
- Distributed systems understanding
- Real-world graph-based engineering problems
