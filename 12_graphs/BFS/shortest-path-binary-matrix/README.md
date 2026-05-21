# Shortest Path in Binary Matrix

## Problem Statement

Given an `n x n` binary matrix `grid`, return the length of the **shortest clear path** in the matrix. If there is no clear path, return `-1`.

A **clear path** in a binary matrix is a path from the **top-left** cell (i.e., `(0, 0)`) to the **bottom-right** cell (i.e., `(n - 1, n - 1)`) such that:

- All the visited cells of the path are `0`.
- All the adjacent cells of the path are **8-directionally connected** (i.e., they are different and they share an edge or a corner).

The **length of a clear path** is the number of visited cells of this path.

### Example 1:

```
Input: grid = [[0,1],[1,0]]
Output: 2
```

![Example 1](https://assets.leetcode.com/uploads/2021/02/18/example1_1.png)

### Example 2:

```
Input: grid = [[0,0,0],[1,1,0],[1,1,0]]
Output: 4
```

![Example 2](https://assets.leetcode.com/uploads/2021/02/18/example2_1.png)

### Example 3:

```
Input: grid = [[1,0,0],[1,1,0],[1,1,0]]
Output: -1
```

### Constraints:

- `n == grid.length`
- `n == grid[i].length`
- `1 <= n <= 100`
- `grid[i][j]` is `0` or `1`

---

## Problem Solving Approach

### Algorithm: Breadth-First Search (BFS)

The solution uses **BFS traversal** to find the shortest path from the top-left corner to the bottom-right corner of the grid.

### Why BFS for Shortest Path?

BFS is the optimal choice for finding the shortest path in an **unweighted graph** because:
- It explores nodes level by level (by distance from the source).
- The first time we reach the destination, we are guaranteed to have found the shortest path.
- All edges have equal weight (1 cell = 1 unit distance).

### Step-by-Step Approach:

1. **Handle Edge Cases:**
   - If the starting cell `grid[0][0]` is blocked (`1`), return `-1`.
   - If the ending cell `grid[n-1][n-1]` is blocked (`1`), return `-1`.

2. **Initialize BFS:**
   - Create a queue with the starting position `(0, 0)` and initial distance `1`.
   - Mark the starting cell as visited by setting it to `1`.

3. **Define 8 Directions:**
   - Unlike typical grid problems with 4 directions, this problem allows diagonal movement.
   - Directions: up, down, left, right, and 4 diagonals.

4. **BFS Exploration:**
   - Pop the front element from the queue.
   - If we've reached the destination `(n-1, n-1)`, return the current distance.
   - For each of the 8 directions:
     - Calculate the new position.
     - Check if it's within bounds and is a clear cell (`0`).
     - Mark it as visited and add to the queue with `distance + 1`.

5. **No Path Found:**
   - If the queue is exhausted without reaching the destination, return `-1`.

### Key Observations:

- **8-directional movement**: Unlike standard grid traversal, we can move diagonally.
- **In-place marking**: The grid is modified directly to track visited cells (changing `0` to `1`), eliminating the need for a separate `visited` array.
- **Early termination**: BFS guarantees the first path found to the destination is the shortest.
- **Distance tracking**: Each queue element stores `(row, col, distance)` to track path length.

---

## Complexity Analysis

### Time Complexity: **O(n²)**

- In the worst case, we visit every cell in the `n x n` grid exactly once.
- Each cell is added to the queue at most once and processed once.
- For each cell, we check 8 neighbors (constant time operation).
- Total operations: O(n²)

### Space Complexity: **O(n²)**

- **BFS Queue**: In the worst case (entire grid is clear), the queue can hold up to O(n²) cells.
- **In-place marking**: No additional visited array is needed since we modify the input grid directly.
- The worst case occurs when the grid is entirely filled with `0`s, and we need to explore all cells.

---

## Code Walkthrough

```go
// 8 directions for movement (including diagonals)
directions := [][]int{
    {-1, -1}, // top-left
    {-1, 0},  // top
    {-1, 1},  // top-right
    {0, -1},  // left
    {0, 1},   // right
    {1, -1},  // bottom-left
    {1, 0},   // bottom
    {1, 1},   // bottom-right
}
```

### Algorithm Flow:

1. **Validation**: Check if start or end cells are blocked.
2. **Queue Initialization**: Start with `(0, 0, 1)` - position and distance of 1.
3. **BFS Loop**:
   - Dequeue front element.
   - Check if destination reached → return distance.
   - Explore all 8 neighbors.
   - Enqueue valid unvisited cells with incremented distance.
4. **Return -1**: If queue empties without reaching destination.

### Example Trace (grid = [[0,1],[1,0]]):

```
Step 1: Start at (0,0), distance = 1
        Queue: [(0,0,1)]
        
Step 2: Process (0,0), check 8 neighbors
        Only (1,1) is valid and clear
        Queue: [(1,1,2)]
        
Step 3: Process (1,1) - this is destination!
        Return distance = 2
```

---

## Related Problems

- [1091. Shortest Path in Binary Matrix](https://leetcode.com/problems/shortest-path-in-binary-matrix/) - LeetCode
- [200. Number of Islands](https://leetcode.com/problems/number-of-islands/)
- [994. Rotting Oranges](https://leetcode.com/problems/rotting-oranges/)
- [542. 01 Matrix](https://leetcode.com/problems/01-matrix/)
- [1293. Shortest Path in a Grid with Obstacles Elimination](https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/)

