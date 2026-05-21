# Number of Islands

## Problem Statement

Given an `m x n` 2D binary grid `grid` which represents a map of `'1'`s (land) and `'0'`s (water), return the **number of islands**.

An **island** is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are surrounded by water.

### Example 1:

```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
```

### Example 2:

```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
```

### Constraints:

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` is `'0'` or `'1'`

---

## Problem Solving Approach

### Algorithm: Breadth-First Search (BFS)

The solution uses **BFS traversal** to explore and count connected components (islands) in the grid.

### Step-by-Step Approach:

1. **Iterate through every cell** in the grid using nested loops.

2. **When a land cell ('1') is found:**
   - Increment the island counter.
   - Start a BFS from this cell to explore all connected land cells.
   - Mark each visited land cell as water ('0') to avoid revisiting.

3. **BFS Exploration:**
   - Use a queue to process cells level by level.
   - For each cell, check all 4 adjacent directions (up, down, left, right).
   - If an adjacent cell is within bounds and is land ('1'):
     - Mark it as visited (change to '0').
     - Add it to the queue for further exploration.

4. **Continue** until all cells in the grid have been processed.

5. **Return** the total island count.

### Why BFS?

- BFS is ideal for exploring all cells of an island systematically.
- By marking cells as visited (changing '1' to '0'), we ensure each cell is processed only once.
- The queue-based approach ensures we explore all neighboring land cells before moving to the next unvisited land cell.

### Key Observations:

- **In-place modification**: The grid is modified directly to track visited cells, eliminating the need for a separate `visited` array.
- **4-directional movement**: Islands are connected only horizontally and vertically (not diagonally).
- **Connected components**: Each BFS call processes one complete island.

---

## Complexity Analysis

### Time Complexity: **O(m × n)**

- We visit each cell in the grid at most once.
- The outer loops iterate through all `m × n` cells.
- Each cell is added to the BFS queue at most once and processed once.
- Total operations: O(m × n)

### Space Complexity: **O(min(m, n))**

- **BFS Queue**: In the worst case, the queue can hold at most `min(m, n)` cells at any time (when the island forms a diagonal pattern).
- **In-place marking**: No additional visited array is needed since we modify the input grid directly.
- In the absolute worst case (entire grid is land), the queue could hold O(m × n) elements, but typically it's bounded by O(min(m, n)).

---

## Code Walkthrough

```go
// Direction vectors for 4-directional movement
var directions = [][]int{
    {-1, 0},  // up
    {1, 0},   // down
    {0, -1},  // left
    {0, 1},   // right
}
```

1. **Grid Traversal**: Loop through each cell `(r, c)`.
2. **Island Detection**: When `grid[r][c] == '1'`, a new island is found.
3. **BFS Initialization**: Create a queue with the starting cell and mark it visited.
4. **BFS Loop**: Process each cell in the queue:
   - Pop the front cell.
   - Explore all 4 neighbors.
   - Add valid unvisited land cells to the queue.
5. **Count Return**: After processing all cells, return the total island count.

---

## Related Problems

- [200. Number of Islands](https://leetcode.com/problems/number-of-islands/) - LeetCode
- [695. Max Area of Island](https://leetcode.com/problems/max-area-of-island/)
- [463. Island Perimeter](https://leetcode.com/problems/island-perimeter/)
- [694. Number of Distinct Islands](https://leetcode.com/problems/number-of-distinct-islands/)

