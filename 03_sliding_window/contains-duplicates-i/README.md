# Contains Duplicate

> **Difficulty:** Easy | **Category:** Arrays & Hashing

## Table of Contents
- [Problem Description](#problem-description)
- [Intuition](#intuition)
- [Approach Explanation](#approach-explanation)
- [Visual Walkthrough](#visual-walkthrough)
- [Complexity Analysis](#complexity-analysis)
- [Key Insights](#key-insights)

---

## Problem Description

Given an integer array `nums`, return `true` if any value appears **at least twice** in the array, and return `false` if every element is distinct.

### Examples

<details>
<summary><b>Example 1</b> (Click to expand)</summary>

```
Input:  nums = [1, 2, 3, 1]
Output: true
```

**Explanation:** The element `1` appears at indices 0 and 3.

</details>

<details>
<summary><b>Example 2</b> (Click to expand)</summary>

```
Input:  nums = [1, 2, 3, 4]
Output: false
```

**Explanation:** All elements are distinct.

</details>

<details>
<summary><b>Example 3</b> (Click to expand)</summary>

```
Input:  nums = [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]
Output: true
```

**Explanation:** Multiple duplicates exist.

</details>

### Constraints

| Constraint | Range |
|------------|-------|
| Array length | `1 <= nums.length <= 10⁵` |
| Element value | `-10⁹ <= nums[i] <= 10⁹` |

---

## Intuition

**The brute force approach:**
- Compare every pair of elements → O(n²)

**The key insight:**
> We just need to know if we've seen a number before!

A **hash set** can answer "have I seen this?" in O(1) time:

```
For each number:
    Seen before? → Duplicate found!
    Not seen?    → Add to set, continue
```

**Alternative approach:**
- Sort the array, then check adjacent elements
- Duplicates will be next to each other after sorting

---

## Approach Explanation

### Approach 1: Brute Force

```
For each element i:
    For each element j > i:
        If nums[i] == nums[j]: return true
Return false
```

**Time:** O(n²) | **Space:** O(1)

### Approach 2: Sorting

```
Step 1: Sort the array
Step 2: Check adjacent elements
        If nums[i] == nums[i+1]: return true
Step 3: Return false
```

**Time:** O(n log n) | **Space:** O(1) or O(n) depending on sort

### Approach 3: Hash Set (Optimal)

```
Step 1: Create empty hash set
Step 2: For each element:
        → If element in set: return true
        → Add element to set
Step 3: Return false
```

### Detailed Steps (Hash Set)

| Step | Action | Purpose |
|------|--------|---------|
| 1 | Create hash set | Track seen numbers |
| 2 | Iterate through array | Check each number once |
| 3 | Check if exists in set | O(1) lookup |
| 4 | Add to set | Remember for future checks |

### Why Hash Set Works

```
┌─────────────────────────────────────────────────┐
│ Hash Set: {seen numbers}                        │
├─────────────────────────────────────────────────┤
│ For each nums[i]:                               │
│                                                 │
│    Is nums[i] in set?                           │
│    YES → Duplicate found! Return true           │
│    NO  → Add to set, continue                   │
│                                                 │
│ Lookup and insert are both O(1)                 │
└─────────────────────────────────────────────────┘
```

---

## Visual Walkthrough

**Input:** `nums = [1, 2, 3, 2, 5]`

### Step-by-Step Execution

```
Initial: hashSet = {}
```

**Iteration 1:** `nums[0] = 1`

```
Is 1 in hashSet? NO
Add: hashSet = {1}
```

**Iteration 2:** `nums[1] = 2`

```
Is 2 in hashSet? NO
Add: hashSet = {1, 2}
```

**Iteration 3:** `nums[2] = 3`

```
Is 3 in hashSet? NO
Add: hashSet = {1, 2, 3}
```

**Iteration 4:** `nums[3] = 2`

```
Is 2 in hashSet? YES!
```

### Found! Return `true`

### Visual Table

| i | nums[i] | In Set? | Action | hashSet |
|---|---------|---------|--------|---------|
| 0 | 1 | No | Add | `{1}` |
| 1 | 2 | No | Add | `{1, 2}` |
| 2 | 3 | No | Add | `{1, 2, 3}` |
| 3 | 2 | **Yes!** | Return `true` | - |

---

### Sorting Approach Visualization

**Input:** `nums = [1, 2, 3, 2, 5]`

```
Before sorting: [1, 2, 3, 2, 5]
After sorting:  [1, 2, 2, 3, 5]
                    ↑  ↑
                  Adjacent duplicates!
```

| i | nums[i] | nums[i+1] | Equal? |
|---|---------|-----------|--------|
| 0 | 1 | 2 | No |
| 1 | 2 | 2 | **Yes!** |

---

## Complexity Analysis

### Comparison Table

| Approach | Time | Space | Notes |
|----------|------|-------|-------|
| Brute Force | O(n²) | O(1) | Compare all pairs |
| Sorting | O(n log n) | O(1)* | Check adjacent after sort |
| **Hash Set** | **O(n)** | **O(n)** | Single pass with lookup |

*Sorting space depends on algorithm used

### Breakdown

**Hash Set Approach**

```
Time Complexity: O(n)
├─ Single pass:     O(n)
└─ Set operations:  O(1) each

Space Complexity: O(n)
└─ Set stores at most n elements
```

**Sorting Approach**

```
Time Complexity: O(n log n)
├─ Sorting:         O(n log n)
└─ Single pass:     O(n)

Space Complexity: O(1) to O(n)
└─ Depends on sorting algorithm
```

### Trade-off

```
┌────────────────────────────────────────────────┐
│ Hash Set:  Fast (O(n)) but uses O(n) space     │
│ Sorting:   Slower but can be O(1) space        │
│                                                │
│ Choose based on constraints!                   │
└────────────────────────────────────────────────┘
```

---

## Key Insights

| # | Insight | Why It Matters |
|---|---------|----------------|
| 1 | **Hash set = O(1) lookup** | Perfect for "have I seen this?" questions |
| 2 | **Sorting groups duplicates** | Alternative when space is limited |
| 3 | **Early termination** | Return as soon as duplicate found |
| 4 | **Space-time trade-off** | Hash set trades space for speed |

### Relationship to Other Problems

| Problem | Technique |
|---------|-----------|
| **Contains Duplicate I** | Simple hash set (any duplicate) |
| Contains Duplicate II | Hash map/sliding window (distance constraint) |
| Contains Duplicate III | Bucket sort / TreeSet (value + distance constraint) |

### Common Pitfalls to Avoid

- Using a map when a set is sufficient (set is simpler)
- Forgetting that sorting modifies the original array
- Not considering the space-time trade-off based on constraints
