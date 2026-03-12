# 3Sum

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

Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` such that:
- `i != j`, `i != k`, and `j != k`
- `nums[i] + nums[j] + nums[k] == 0`

> **Note:** The solution set must not contain duplicate triplets.

### Examples

<details>
<summary><b>Example 1</b> (Click to expand)</summary>

```
Input:  nums = [-1, 0, 1, 2, -1, -4]
Output: [[-1, -1, 2], [-1, 0, 1]]
```

**Explanation:**
| Triplet | Calculation | Sum |
|---------|-------------|-----|
| `[-1, 0, 1]` | (-1) + 0 + 1 | 0 ✓ |
| `[-1, -1, 2]` | (-1) + (-1) + 2 | 0 ✓ |

</details>

<details>
<summary><b>Example 2</b> (Click to expand)</summary>

```
Input:  nums = [0, 1, 1]
Output: []
```

**Explanation:** No triplet sums to 0.

</details>

<details>
<summary><b>Example 3</b> (Click to expand)</summary>

```
Input:  nums = [0, 0, 0]
Output: [[0, 0, 0]]
```

**Explanation:** The only possible triplet `[0, 0, 0]` sums to 0.

</details>

### Constraints

| Constraint | Range |
|------------|-------|
| Array length | `3 <= nums.length <= 3000` |
| Element value | `-10⁵ <= nums[i] <= 10⁵` |

---

## Intuition

**Why not brute force?**
- A naive O(n³) approach checks every possible triplet
- For `n = 3000`, that's ~27 billion operations!

**The key insight:**
> If we **sort** the array first, we can use the **two-pointer technique** to find pairs that sum to a target in O(n) time.

This transforms the problem:
```
Find triplet where: nums[i] + nums[j] + nums[k] = 0
                    ↓
Fix nums[i], then find: nums[j] + nums[k] = -nums[i]
```

---

## Approach Explanation

### Algorithm: Sort + Two Pointers

```
Step 1: Sort the array
Step 2: For each element nums[i]:
        → Use two pointers to find pairs that sum to -nums[i]
Step 3: Skip duplicates at every level
```

### Detailed Steps

| Step | Action | Purpose |
|------|--------|---------|
| 1 | Sort the array | Enable two-pointer technique & easy duplicate detection |
| 2 | Fix first element `i` | Reduce to two-sum problem |
| 3 | Set `j = i+1`, `k = end` | Two pointers for remaining elements |
| 4 | Calculate sum | Check if triplet sums to 0 |
| 5 | Adjust pointers | Move based on sum comparison |
| 6 | Skip duplicates | Avoid duplicate triplets in result |

### Pointer Movement Logic

```
sum = nums[i] + nums[j] + nums[k]

┌─────────────────┬────────────────────────────┐
│ If sum == 0     │ Found triplet! Move both   │
│                 │ pointers inward            │
├─────────────────┼────────────────────────────┤
│ If sum < 0      │ Need larger sum → move j   │
│                 │ right (→)                  │
├─────────────────┼────────────────────────────┤
│ If sum > 0      │ Need smaller sum → move k  │
│                 │ left (←)                   │
└─────────────────┴────────────────────────────┘
```

---

## Visual Walkthrough

**Input:** `[-1, 2, 1, 2, -1, -3]`

### Step 1: Sort

```
Before: [-1, 2, 1, 2, -1, -3]
After:  [-3, -1, -1, 1, 2, 2]
         ↑
         sorted!
```

### Step 2: Iterate with Two Pointers

**Iteration 1:** `i = 0` (value: -3)

```
Array:  [-3, -1, -1,  1,  2,  2]
          i   j               k
              ↑               ↑
           pointer          end

Target sum for j + k = -(-3) = 3
```

| j | k | nums[j] + nums[k] | Action |
|---|---|-------------------|--------|
| 1 | 5 | -1 + 2 = 1 | < 3, move j → |
| 2 | 5 | -1 + 2 = 1 | < 3, move j → |
| 3 | 5 | 1 + 2 = 3 | = 3, **Found! [-3, 1, 2]** |

**Iteration 2:** `i = 1` (value: -1)

```
Array:  [-3, -1, -1,  1,  2,  2]
              i   j           k

Target sum for j + k = -(-1) = 1
```

| j | k | nums[j] + nums[k] | Action |
|---|---|-------------------|--------|
| 2 | 5 | -1 + 2 = 1 | = 1, **Found! [-1, -1, 2]** |

### Final Result

```
[[-3, 1, 2], [-1, -1, 2]]
```

---

## Complexity Analysis

### Comparison Table

| Approach | Time | Space | Notes |
|----------|------|-------|-------|
| Brute Force | O(n³) | O(1) | Check all triplets |
| **Two Pointers** | **O(n²)** | **O(1)** | Optimal solution |

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
- Output array not counted

---

## Key Insights

| # | Insight | Why It Matters |
|---|---------|----------------|
| 1 | **Sorting is essential** | Enables two-pointer technique and simplifies duplicate detection |
| 2 | **Three-level duplicate skipping** | Skip at `i`, `j`, and `k` to avoid duplicate triplets |
| 3 | **Two-pointer efficiency** | Sorted array means we know which direction to move pointers |

### Common Pitfalls to Avoid

- Forgetting to skip duplicates after finding a valid triplet
- Off-by-one errors when skipping duplicates
- Not handling the case where all elements are the same
