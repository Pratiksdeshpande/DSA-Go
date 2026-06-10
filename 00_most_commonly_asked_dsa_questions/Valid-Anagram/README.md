# Valid Anagram

## Table of Contents

* [Problem Statement](#problem-statement)
* [Examples](#examples)
* [Constraints](#constraints)
* [Understanding the Problem](#understanding-the-problem)
* [Key Observations](#key-observations)
* [Approach 1: Sorting](#approach-1-sorting)
* [Approach 2: Frequency Counting (Optimal)](#approach-2-frequency-counting-optimal)
* [Dry Run Example 1](#dry-run-example-1)
* [Dry Run Example 2](#dry-run-example-2)
* [Edge Cases](#edge-cases)
* [Complexity Analysis Summary](#complexity-analysis-summary)
* [Interview Discussion Points](#interview-discussion-points)

---

# Problem Statement

Given two strings `s` and `t`, determine whether `t` is an anagram of `s`.

Two strings are considered anagrams if:

* They contain exactly the same characters.
* Every character appears the same number of times in both strings.
* Character ordering does not matter.

Return:

* `true` if `t` is an anagram of `s`
* `false` otherwise

---

# Examples

## Example 1

```text
Input:
s = "anagram"
t = "nagaram"

Output:
true
```

Explanation:

Both strings contain:

```text
a -> 3
n -> 1
g -> 1
r -> 1
m -> 1
```

Since character frequencies are identical, the strings are anagrams.

---

## Example 2

```text
Input:
s = "rat"
t = "car"

Output:
false
```

Explanation:

```text
s:
r -> 1
a -> 1
t -> 1

t:
c -> 1
a -> 1
r -> 1
```

Character frequencies differ.

---

## Example 3

```text
Input:
s = "listen"
t = "silent"

Output:
true
```

Both strings contain exactly the same characters with identical frequencies.

---

# Constraints

Typical interview constraints:

```text
1 <= s.length <= 50,000
1 <= t.length <= 50,000

s and t consist of lowercase English letters.
```

Possible follow-up:

```text
What if strings contain Unicode characters?
```

In that case, a frequency map using runes should be used.

---

# Understanding the Problem

The most important realization is:

```text
Order does not matter.
Frequency matters.
```

Consider:

```text
listen
silent
```

Although the positions of characters differ, the frequency of each character remains the same.

Therefore, instead of comparing positions, we should compare character counts.

---

# Key Observations

## Observation 1

If lengths are different:

```text
s = "abc"
t = "ab"
```

They can never be anagrams.

Therefore:

```text
If len(s) != len(t)
Return false
```

This acts as an immediate optimization.

---

## Observation 2

Two strings are anagrams if and only if:

```text
Frequency(character in s)
=
Frequency(character in t)
```

for every character.

---

## Observation 3

Since ordering is irrelevant:

```text
Sorting is possible
```

but

```text
Frequency counting is more efficient.
```

---

# Approach 1: Sorting

## Idea

Sort both strings alphabetically.

Example:

```text
listen -> eilnst
silent -> eilnst
```

If sorted strings are identical:

```text
Return true
```

Otherwise:

```text
Return false
```

---

## Steps

1. Check length.
2. Convert strings into character arrays.
3. Sort both arrays.
4. Compare them.

---

## Complexity

### Time Complexity

```text
O(n log n)
```

Sorting dominates the runtime.

---

### Space Complexity

```text
O(n)
```

Additional memory may be required for sorting.

---

# Approach 2: Frequency Counting (Optimal)

## Idea

Count occurrences of each character.

Example:

```text
s = "aabbcc"
```

Frequency table:

```text
a -> 2
b -> 2
c -> 2
```

Now process second string:

```text
t = "ccbbaa"
```

Decrease frequencies while scanning.

Final frequencies:

```text
a -> 0
b -> 0
c -> 0
```

Since all counts balance out, the strings are anagrams.

---

## Steps

### Step 1

Check lengths.

If lengths differ:

```text
Return false
```

---

### Step 2

Create a frequency structure.

Store count of every character from first string.

Example:

```text
s = "aab"

Map:

a -> 2
b -> 1
```

---

### Step 3

Traverse second string.

For every character:

```text
Decrease its count.
```

If count becomes negative or character is unavailable:

```text
Return false
```

---

### Step 4

If entire second string is processed successfully:

```text
Return true
```

Because equal lengths guarantee all frequencies balanced perfectly.

---

## Complexity

### Time Complexity

```text
O(n)
```

One pass over each string.

---

### Space Complexity

HashMap solution:

```text
O(k)
```

where:

```text
k = number of unique characters
```

---

For lowercase English letters only:

```text
O(1)
```

because only 26 positions are required.

---

# Dry Run Example 1

## Input

```text
s = "anagram"
t = "nagaram"
```

---

### Build Frequency Table

```text
a -> 3
n -> 1
g -> 1
r -> 1
m -> 1
```

---

### Process "nagaram"

Read n

```text
n -> 0
```

Read a

```text
a -> 2
```

Read g

```text
g -> 0
```

Read a

```text
a -> 1
```

Read r

```text
r -> 0
```

Read a

```text
a -> 0
```

Read m

```text
m -> 0
```

---

### Final State

```text
a -> 0
n -> 0
g -> 0
r -> 0
m -> 0
```

Result:

```text
true
```

---

# Dry Run Example 2

## Input

```text
s = "rat"
t = "car"
```

---

### Build Frequency Table

```text
r -> 1
a -> 1
t -> 1
```

---

### Process "car"

Read c

Character not available.

Result:

```text
false
```

No further processing required.

---

# Edge Cases

## Different Lengths

```text
s = "abc"
t = "ab"
```

Output:

```text
false
```

---

## Empty Strings

```text
s = ""
t = ""
```

Output:

```text
true
```

---

## Single Character

```text
s = "a"
t = "a"
```

Output:

```text
true
```

---

## Same Characters Different Counts

```text
s = "aab"
t = "abb"
```

Output:

```text
false
```

Frequency mismatch.

---

## Case Sensitivity

```text
s = "Rat"
t = "rat"
```

Output:

```text
false
```

Because:

```text
'R' != 'r'
```

unless explicitly stated otherwise.

---

# Complexity Analysis Summary

| Approach                   | Time       | Space |
| -------------------------- | ---------- | ----- |
| Sorting                    | O(n log n) | O(n)  |
| Frequency Map              | O(n)       | O(k)  |
| Frequency Array (26 chars) | O(n)       | O(1)  |

---

# Interview Discussion Points

## Why is sorting not optimal?

Sorting introduces:

```text
O(n log n)
```

time complexity.

Since only frequency information is needed, sorting performs unnecessary work.

---

## Why does frequency counting work?

Anagrams require identical character frequencies.

By counting characters from the first string and removing them using the second string, we directly verify this requirement.

---

## When should a HashMap be used?

Use a HashMap when:

```text
Character set is unknown
```

Examples:

```text
Unicode
Special symbols
Mixed languages
```

---

## When should a Fixed Array be used?

Use a fixed-size array when:

```text
Input contains only lowercase English letters.
```

Benefits:

* Faster lookups
* No hashing overhead
* Constant space usage

---

# Senior-Level Interview Answer

A senior engineer should immediately recognize that ordering is irrelevant for anagrams. Therefore, instead of sorting, the problem can be solved by comparing character frequencies. A frequency-counting solution achieves O(n) time complexity and O(1) space complexity when the input is limited to lowercase English letters, making it the optimal approach.
