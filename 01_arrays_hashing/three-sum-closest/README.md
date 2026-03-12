# 3Sum Closest

> **Difficulty:** Medium | **Category:** Arrays & Hashing, Two Pointers

## Table of Contents
- [Problem Description](#problem-description)
- [Intuition](#intuition)
- [Approach Explanation](#approach-explanation)
- [Visual Walkthrough](#visual-walkthrough)
- [Complexity Analysis](#complexity-analysis)
- [Key Insights](#key-insights)

---

## Problem Description

Given an integer array `nums` of length `n` and an integer `target`, find three integers in `nums` such that the sum is **closest** to `target`.

Return the sum of the three integers.

> **Note:** You may assume that each input would have exactly one solution.

### Examples

<details>
<summary><b>Example 1</b> (Click to expand)</summary>

```
Input:  nums = [-1, 2, 1, -4], target = 1
Output: 2
```

**Explanation:**
| Triplet | Sum | Distance from target |
|---------|-----|---------------------|
| `[-1, 2, 1]` | 2 | \|2 - 1\| = 1 ✓ |
| `[-1, 2, -4]` | -3 | \|-3 - 1\| = 4 |
| `[-1, 1, -4]` | -4 | \|-4 - 1\| = 5 |
| `[2, 1, -4]` | -1 | \|-1 - 1\| = 2 |

The sum that is closest to the target is `2` with distance `1`.

</details>

<details>
<summary><b>Example 2</b> (Click to expand)</summary>

```
Input:  nums = [0, 0, 0], target = 1
Output: 0
```

**Explanation:** The only possible sum is `0 + 0 + 0 = 0`, which is the closest to target `1`.

</details>

### Constraints

| Constraint    | Range                      |
|---------------|----------------------------|
| Array length  | `3 <= nums.length <= 500`  |
| Element value | `-1000 <= nums[i] <= 1000` |
| Target value  | `-10⁴ <= target <= 10⁴`    |

---

## Intuition

**How is this different from 3Sum?**
- In 3Sum, we find triplets that sum to exactly `0`
- Here, we find the triplet whose sum is **closest** to a given target

**The key insight:**
> We can still use the **sort + two-pointer technique**, but instead of looking for exact matches, we track the **minimum distance** from the target.

```
For each triplet sum:
    distance = |sum - target|
    
    Keep the sum with the smallest distance
```

If we find `sum == target`, we can return immediately (distance = 0, can't do better).

---

## Approach Explanation

### Algorithm: Sort + Two Pointers

```
Step 1: Sort the array
Step 2: Initialize answer with first three elements
Step 3: For each element nums[i]:
        → Use two pointers to find the closest sum
        → Update answer if current sum is closer to target
Step 4: Return the closest sum found
```

### Detailed Steps

| Step  | Action                                            | Purpose                                           |
|-------|---------------------------------------------------|---------------------------------------------------|
| 1     | Sort the array                                    | Enable two-pointer technique                      |
| 2     | Initialize `answer = nums[0] + nums[1] + nums[2]` | Start with a valid triplet sum                    |
| 3     | Fix first element `i`                             | Reduce to two-pointer problem                     |
| 4     | Set `j = i+1`, `k = end`                          | Two pointers for remaining elements               |
| 5     | Calculate sum & distance                          | Compare `(sum - target)` with `(answer - target)` |
| 6     | Update answer if closer                           | Keep track of the best sum                        |
| 7     | Adjust pointers                                   | Move based on sum vs target comparison            |

### Pointer Movement Logic

```
sum = nums[i] + nums[j] + nums[k]

┌─────────────────┬────────────────────────────┐
│ If sum == target│ Perfect match! Return sum  │
│                 │ (can't get closer)         │
├─────────────────┼────────────────────────────┤
│ If sum < target │ Need larger sum → move j   │
│                 │ right (→)                  │
├─────────────────┼────────────────────────────┤
│ If sum > target │ Need smaller sum → move k  │
│                 │ left (←)                   │
└─────────────────┴────────────────────────────┘

At each step: if |sum - target| < |answer - target|
              then answer = sum
```

---

## Visual Walkthrough

**Input:** `nums = [-1, 2, 1, -4]`, `target = 1`

### Step 1: Sort

```
Before: [-1, 2, 1, -4]
After:  [-4, -1, 1, 2]
         ↑
         sorted!
```

### Step 2: Initialize Answer

```
answer = nums[0] + nums[1] + nums[2]
answer = -4 + (-1) + 1 = -4
```

### Step 3: Iterate with Two Pointers

**Iteration 1:** `i = 0` (value: -4)

```
Array:  [-4, -1,  1,  2]
          i   j       k

Current answer = -4 (distance from 1 = 5)
```

| j | k | sum                 | \|sum - target\| | \|answer - target\| | Action                                     |
|---|---|---------------------|------------------|---------------------|--------------------------------------------|
| 1 | 3 | -4 + (-1) + 2 = -3  | 4                | 5                   | -3 closer! answer = -3, sum < target → j++ |
| 2 | 3 | -4 + 1 + 2 = -1     | 2                | 4                   | -1 closer! answer = -1, sum < target → j++ |

**Iteration 2:** `i = 1` (value: -1)

```
Array:  [-4, -1,  1,  2]
              i   j   k

Current answer = -1 (distance from 1 = 2)
```

| j | k | sum            | \|sum - target\| | \|answer - target\| | Action                                   |
|---|---|----------------|------------------|---------------------|------------------------------------------|
| 2 | 3 | -1 + 1 + 2 = 2 | 1                | 2                   | 2 closer! answer = 2, sum > target → k-- |

### Final Result

```
answer = 2 (closest to target 1, with distance 1)
```

---

## Complexity Analysis

### Comparison Table

| Approach         | Time      | Space    | Notes              |
|------------------|-----------|----------|--------------------|
| Brute Force      | O(n³)     | O(1)     | Check all triplets |
| **Two Pointers** | **O(n²)** | **O(1)** | Optimal solution   |

### Breakdown

**Time Complexity: O(n²)**
```
Sorting:           O(n log n)
Outer loop:        O(n)
  └─ Two pointers: O(n)
                   ─────────
Total:             O(n log n) + O(n²) = O(n²)
```

**Space Complexity: O(1)**
- Only constant extra variables used
- Sorting is done in-place

---

## Key Insights

| # | Insight                        | Why It Matters                                         |
|---|--------------------------------|--------------------------------------------------------|
| 1 | **Same technique as 3Sum**     | Sort + two pointers works for "closest" variant too    |
| 2 | **Track minimum distance**     | Compare `(sum - target)` instead of checking equality  |
| 3 | **Early termination**          | If `sum == target`, return immediately (optimal)       |
| 4 | **Pointer movement unchanged** | Move `j` right if sum too small, `k` left if too large |

### Key Differences from 3Sum

| Aspect      | 3Sum                                  | 3Sum Closest                         |
|-------------|---------------------------------------|--------------------------------------|
| Goal        | Find sums equal to 0                  | Find sum closest to target           |
| Return      | List of triplets                      | Single sum value                     |
| Duplicates  | Must skip to avoid duplicate triplets | Can skip for optimization (optional) |
| Exact match | Required                              | Not required (but optimal if found)  |

### Common Pitfalls to Avoid

- Forgetting to initialize `answer` with a valid triplet sum
- Using integer overflow when calculating distances (use `math.Abs` carefully)
- Not returning early when exact match is found
