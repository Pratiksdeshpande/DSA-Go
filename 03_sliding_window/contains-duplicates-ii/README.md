# Contains Duplicate II

> **Difficulty:** Easy | **Category:** Sliding Window, Hash Map

## Table of Contents
- [Problem Description](#problem-description)
- [Intuition](#intuition)
- [Approach Explanation](#approach-explanation)
- [Visual Walkthrough](#visual-walkthrough)
- [Complexity Analysis](#complexity-analysis)
- [Key Insights](#key-insights)

---

## Problem Description

Given an integer array `nums` and an integer `k`, return `true` if there are **two distinct indices** `i` and `j` in the array such that:
- `nums[i] == nums[j]`
- `abs(i - j) <= k`

### Examples

<details>
<summary><b>Example 1</b> (Click to expand)</summary>

```
Input:  nums = [1, 2, 3, 1], k = 3
Output: true
```

**Explanation:**
| Index i | Index j | Values | Distance |
|---------|---------|--------|----------|
| 0 | 3 | nums[0] = nums[3] = 1 | \|0 - 3\| = 3 ≤ 3 ✓ |

</details>

<details>
<summary><b>Example 2</b> (Click to expand)</summary>

```
Input:  nums = [1, 0, 1, 1], k = 1
Output: true
```

**Explanation:**
| Index i | Index j | Values | Distance |
|---------|---------|--------|----------|
| 2 | 3 | nums[2] = nums[3] = 1 | \|2 - 3\| = 1 ≤ 1 ✓ |

</details>

<details>
<summary><b>Example 3</b> (Click to expand)</summary>

```
Input:  nums = [1, 2, 3, 1, 2, 3], k = 2
Output: false
```

**Explanation:** The closest duplicates are at distance 3, which is > k.

</details>

### Constraints

| Constraint | Range |
|------------|-------|
| Array length | `1 <= nums.length <= 10⁵` |
| Element value | `-10⁹ <= nums[i] <= 10⁹` |
| Window size | `0 <= k <= 10⁵` |

---

## Intuition

**The brute force approach:**
- For each element, check all elements within distance `k` → O(n × k)

**The key insight:**
> We only care about duplicates within a **window of size k**!

This is a classic **sliding window** problem:

```
Window of size k
┌─────────────────┐
[1, 2, 3, 1, 2, 3]
 └───────────────┘
 
Check: Are there any duplicates within this window?
```

**Two optimal approaches:**

1. **Hash Map (track last index):** For each number, remember where we last saw it
2. **Sliding Window (hash set):** Maintain a set of at most `k` recent elements

---

## Approach Explanation

### Approach 1: Brute Force

```
For each element i:
    Check elements from i+1 to min(n, i+k):
        If nums[i] == nums[j]: return true
Return false
```

**Time:** O(n × k) | **Space:** O(1)

### Approach 2: Hash Map (Track Last Index)

```
Step 1: Create hash map {number → last_seen_index}
Step 2: For each element nums[i]:
        → If number exists in map AND (i - last_index) ≤ k:
            Return true
        → Update map[nums[i]] = i
Step 3: Return false
```

### Approach 3: Sliding Window with Hash Set

```
Step 1: Maintain a hash set (window) of at most k elements
Step 2: For each element:
        → If window size > k: remove leftmost element
        → If current element exists in window: return true
        → Add current element to window
Step 3: Return false
```

### Detailed Steps (Hash Map Approach)

| Step | Action | Purpose |
|------|--------|---------|
| 1 | Create hash map | Store last seen index for each number |
| 2 | Iterate through array | Check each number once |
| 3 | Check if duplicate within range | `i - hashMap[num] <= k` |
| 4 | Update last seen index | `hashMap[num] = i` |

### Why Hash Map Works

```
For each nums[i]:
┌─────────────────────────────────────────────────┐
│ If nums[i] was seen before at index j:          │
│                                                 │
│    Check: i - j <= k ?                          │
│                                                 │
│    YES → Found duplicate within range!          │
│    NO  → Update to new index (closer for next)  │
└─────────────────────────────────────────────────┘
```

### Sliding Window Logic

```
Window maintains last k elements

┌─────────────────────────────────────────────────┐
│ When R - L > k:                                 │
│     Remove arr[L] from window                   │
│     L++                                         │
│                                                 │
│ If arr[R] already in window:                    │
│     → Duplicate found within distance k!        │
│                                                 │
│ Add arr[R] to window                            │
└─────────────────────────────────────────────────┘
```

---

## Visual Walkthrough

**Input:** `nums = [1, 2, 3, 1]`, `k = 3`

### Hash Map Approach

```
Initial: hashMap = {}
```

**Iteration 1:** `i = 0`, `nums[i] = 1`

```
Is 1 in hashMap? NO
Store: hashMap = {1: 0}
```

**Iteration 2:** `i = 1`, `nums[i] = 2`

```
Is 2 in hashMap? NO
Store: hashMap = {1: 0, 2: 1}
```

**Iteration 3:** `i = 2`, `nums[i] = 3`

```
Is 3 in hashMap? NO
Store: hashMap = {1: 0, 2: 1, 3: 2}
```

**Iteration 4:** `i = 3`, `nums[i] = 1`

```
Is 1 in hashMap? YES, at index 0
Distance = 3 - 0 = 3
Is 3 <= k (3)? YES!
```

### Found! Return `true`

### Visual Table

| i | nums[i] | In Map? | Last Index | Distance | ≤ k? | hashMap |
|---|---------|---------|------------|----------|------|---------|
| 0 | 1 | No | - | - | - | `{1: 0}` |
| 1 | 2 | No | - | - | - | `{1: 0, 2: 1}` |
| 2 | 3 | No | - | - | - | `{1: 0, 2: 1, 3: 2}` |
| 3 | 1 | **Yes** | 0 | 3 | **Yes!** | Return `true` |

---

### Sliding Window Approach

**Input:** `nums = [1, 2, 3, 1]`, `k = 3`

```
Window size = k = 3

Step 1: [1]           window = {1}
Step 2: [1, 2]        window = {1, 2}
Step 3: [1, 2, 3]     window = {1, 2, 3}
Step 4: [1, 2, 3, 1]  → 1 already in window!
                        
        ✓ Found duplicate!
```

---

## Complexity Analysis

### Comparison Table

| Approach | Time | Space | Notes |
|----------|------|-------|-------|
| Brute Force | O(n × k) | O(1) | Check pairs within range |
| **Hash Map** | **O(n)** | **O(n)** | Track last index |
| **Sliding Window** | **O(n)** | **O(min(n, k))** | Best space efficiency |

### Breakdown

**Hash Map Approach**

```
Time Complexity: O(n)
├─ Single pass:        O(n)
└─ Hash map ops:       O(1) each

Space Complexity: O(n)
└─ Hash map stores at most n elements
```

**Sliding Window Approach**

```
Time Complexity: O(n)
├─ Single pass:        O(n)
└─ Hash set ops:       O(1) each

Space Complexity: O(min(n, k))
└─ Window stores at most k elements
```

### Trade-off

```
┌────────────────────────────────────────────────┐
│ Hash Map:       Simpler but O(n) space         │
│ Sliding Window: Better space O(k) when k << n  │
└────────────────────────────────────────────────┘
```

---

## Key Insights

| # | Insight | Why It Matters |
|---|---------|----------------|
| 1 | **Think in windows** | We only care about elements within distance k |
| 2 | **Track last index** | For duplicates, the most recent position matters |
| 3 | **Sliding window = bounded set** | Automatically maintains k-sized constraint |
| 4 | **Hash operations are O(1)** | Makes both approaches linear time |

### Relationship to Other Problems

| Problem | Technique |
|---------|-----------|
| Contains Duplicate I | Simple hash set (any duplicate) |
| **Contains Duplicate II** | Hash map/sliding window (distance constraint) |
| Contains Duplicate III | Bucket sort / TreeSet (value + distance constraint) |

### Common Pitfalls to Avoid

- Not updating the index when a duplicate is found but distance > k
- Off-by-one error with window size (should be ≤ k, not < k)
- Using `abs(i - j)` when iterating forward (j is always > i)
