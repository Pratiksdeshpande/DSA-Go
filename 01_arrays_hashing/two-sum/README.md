# Two Sum

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

Given an array of integers `nums` and an integer `target`, return the **indices of the two numbers** such that they add up to `target`.

> **Note:** 
> - Each input has **exactly one solution**
> - You **cannot use the same element twice**
> - The answer can be returned in **any order**

### Examples

<details>
<summary><b>Example 1</b> (Click to expand)</summary>

```
Input:  nums = [2, 7, 11, 15], target = 9
Output: [0, 1]
```

**Explanation:**
| Index i | Index j | Calculation | Sum |
|---------|---------|-------------|-----|
| 0 | 1 | 2 + 7 | 9 ✓ |

`nums[0] + nums[1] = 2 + 7 = 9`, so we return `[0, 1]`.

</details>

<details>
<summary><b>Example 2</b> (Click to expand)</summary>

```
Input:  nums = [3, 2, 4], target = 6
Output: [1, 2]
```

**Explanation:** `nums[1] + nums[2] = 2 + 4 = 6`

</details>

<details>
<summary><b>Example 3</b> (Click to expand)</summary>

```
Input:  nums = [3, 3], target = 6
Output: [0, 1]
```

**Explanation:** `nums[0] + nums[1] = 3 + 3 = 6`

</details>

### Constraints

| Constraint    | Range                     |
|---------------|---------------------------|
| Array length  | `2 <= nums.length <= 10⁴` |
| Element value | `-10⁹ <= nums[i] <= 10⁹`  |
| Target value  | `-10⁹ <= target <= 10⁹`   |

---

## Intuition

**The brute force approach:**
- Check every pair of numbers → O(n²)
- For each element, scan the rest of the array

**The key insight:**
> Instead of searching for a pair, search for the **complement**!

```
If a + b = target
Then b = target - a
```

For each number, we ask: *"Have I seen its complement before?"*

Using a **hash map**, we can answer this in O(1) time!

```
For each nums[i]:
    complement = target - nums[i]
    
    If complement in hashMap → Found the pair!
    Else → Store nums[i] in hashMap for future lookups
```

---

## Approach Explanation

### Approach 1: Brute Force

```
For each element i:
    For each element j > i:
        If nums[i] + nums[j] == target:
            Return [i, j]
```

**Time:** O(n²) | **Space:** O(1)

### Approach 2: Hash Map (Optimal)

```
Step 1: Create an empty hash map
Step 2: For each element nums[i]:
        → Calculate complement = target - nums[i]
        → If complement exists in map, return [map[complement], i]
        → Otherwise, store nums[i] with its index in map
```

### Detailed Steps

| Step | Action                     | Purpose                               |
|------|----------------------------|---------------------------------------|
| 1    | Create hash map            | Store seen numbers with their indices |
| 2    | Iterate through array      | Check each number once                |
| 3    | Calculate complement       | `complement = target - nums[i]`       |
| 4    | Check if complement exists | O(1) lookup in hash map               |
| 5    | Store current number       | For future complement lookups         |

### Why Hash Map Works

```
┌─────────────────────────────────────────────────┐
│ Hash Map: { number → index }                    │
├─────────────────────────────────────────────────┤
│ For nums[i], we need to find nums[j] where:     │
│                                                 │
│    nums[i] + nums[j] = target                   │
│    nums[j] = target - nums[i]  (complement)     │
│                                                 │
│ Check if complement exists in O(1) time!        │
└─────────────────────────────────────────────────┘
```

---

## Visual Walkthrough

**Input:** `nums = [2, 7, 11, 15]`, `target = 9`

### Step-by-Step Execution

```
Initial: hashMap = {}
```

**Iteration 1:** `i = 0`, `nums[i] = 2`

```
complement = 9 - 2 = 7
Is 7 in hashMap? NO
Store: hashMap = {2: 0}
```

**Iteration 2:** `i = 1`, `nums[i] = 7`

```
complement = 9 - 7 = 2
Is 2 in hashMap? YES! At index 0
```

### Found! Return `[0, 1]`

### Visual Table

| i | nums[i] | complement | In Map?  | Action        | hashMap  |
|---|---------|------------|----------|---------------|----------|
| 0 | 2       | 7          | No       | Store 2       | `{2: 0}` |
| 1 | 7       | 2          | **Yes!** | Return [0, 1] | -        |

---

## Complexity Analysis

### Comparison Table

| Approach     | Time     | Space    | Notes                   |
|--------------|----------|----------|-------------------------|
| Brute Force  | O(n²)    | O(1)     | Check all pairs         |
| **Hash Map** | **O(n)** | **O(n)** | Single pass with lookup |

### Breakdown

**Time Complexity: O(n)**
```
Single pass:       O(n)
Hash map lookup:   O(1) per lookup
                   ─────────
Total:             O(n)
```

**Space Complexity: O(n)**
- Hash map stores at most `n` elements
- Each element stored with its index

### Trade-off

```
┌────────────────────────────────────────┐
│ Brute Force: Slow but no extra memory  │
│ Hash Map:    Fast but uses O(n) space  │
│                                        │
│ We trade SPACE for TIME!               │
└────────────────────────────────────────┘
```

---

## Key Insights

| # | Insight                    | Why It Matters                                |
|---|----------------------------|-----------------------------------------------|
| 1 | **Think in complements**   | Instead of finding `a + b`, find `target - a` |
| 2 | **Hash map = O(1) lookup** | Eliminates nested loop                        |
| 3 | **Single pass is enough**  | Store and check simultaneously                |
| 4 | **Order doesn't matter**   | We can return indices in any order            |

### Why This Pattern is Important

Two Sum is the foundation for many other problems:

| Problem             | Technique         |
|---------------------|-------------------|
| Two Sum             | Hash map lookup   |
| Three Sum           | Fix one + Two Sum |
| Four Sum            | Fix two + Two Sum |
| Two Sum II (Sorted) | Two pointers      |

### Common Pitfalls to Avoid

- Returning the same index twice (can't use same element twice)
- Forgetting to store the current element after checking
- Using the wrong order for indices (usually doesn't matter, but be consistent)

-----------------
