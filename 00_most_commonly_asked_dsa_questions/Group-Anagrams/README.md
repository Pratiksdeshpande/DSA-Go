# Group Anagrams

## Table of Contents

* [Problem Statement](#problem-statement)
* [Examples](#examples)
* [Constraints](#constraints)
* [Understanding the Problem](#understanding-the-problem)
* [Key Observation](#key-observation)
* [Brute Force Approach](#brute-force-approach)
* [Optimal Approach](#optimal-approach)
* [Algorithm](#algorithm)
* [Dry Run](#dry-run)
* [Complexity Analysis](#complexity-analysis)
* [Edge Cases](#edge-cases)
* [Interview Discussion Points](#interview-discussion-points)
* [Common Follow-Up Questions](#common-follow-up-questions)

---

# Problem Statement

Given an array of strings `strs`, group all strings that are anagrams of each other.

Two strings are considered anagrams if:

* They contain exactly the same characters.
* Every character appears the same number of times.
* Character order does not matter.

Return all groups of anagrams in any order.

---

# Examples

## Example 1

### Input

```go
strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
```

### Output

```go
[
    ["eat", "tea", "ate"],
    ["tan", "nat"],
    ["bat"]
]
```

### Explanation

```text
eat -> aet
tea -> aet
ate -> aet
```

All three produce the same character composition and belong to the same group.

```text
tan -> ant
nat -> ant
```

These belong to another group.

```text
bat -> abt
```

Forms its own group.

---

## Example 2

### Input

```go
strs := []string{""}
```

### Output

```go
[
    [""]
]
```

---

## Example 3

### Input

```go
strs := []string{"a"}
```

### Output

```go
[
    ["a"]
]
```

---

# Constraints

```text
1 <= strs.length <= 10^4

0 <= strs[i].length <= 100

strs[i] consists of lowercase English letters.
```

---

# Understanding the Problem

The goal is to group words that contain the same characters with the same frequencies.

For example:

```text
eat
tea
ate
```

All contain:

a -> 1
e -> 1
t -> 1

Therefore they belong to the same group.

Similarly:

```text
tan
nat
```

Both contain:

a -> 1
n -> 1
t -> 1

Hence they belong to another group.

The challenge is efficiently identifying which words are anagrams without comparing every word against every other word.

---

# Key Observation

All anagrams share the same canonical representation.

Consider:

```text
eat
tea
ate
```

If we sort the characters of every word:

```text
eat -> aet
tea -> aet
ate -> aet
```

All anagrams produce the same sorted string.

This sorted string can be used as a unique key.

Instead of comparing words with one another, we simply:

1. Generate a key.
2. Group words with the same key.

This transforms the problem into a hashmap grouping problem.

---

# Brute Force Approach

## Idea

For every string:

* Compare it with every other string.
* Check if both are anagrams.

### Example

```text
eat vs tea
eat vs tan
eat vs ate
eat vs nat
...
```

To check if two strings are anagrams:

* Count character frequencies OR
* Sort both strings and compare

---

## Time Complexity

Let:

```text
n = number of strings
k = average string length
```

Comparing every pair:

```text
O(n²)
```

Checking if two strings are anagrams:

```text
O(k)
```

Overall:

```text
O(n² × k)
```

This becomes inefficient for large inputs.

---

# Optimal Approach

## Idea

Instead of comparing strings with one another:

Generate a unique key for every word.

For the sorting-based solution:

```text
eat -> aet
tea -> aet
ate -> aet
```

Store them in a hashmap:

```go
map[string][]string
```

Where:

```text
key = sorted string
value = list of words
```

Words with identical keys automatically belong to the same group.

---

# Algorithm

1. Create a hashmap.

```go
groups := make(map[string][]string)
```

2. Iterate through every word.

3. Sort the characters of the word.

4. Convert sorted characters back to a string.

5. Use the sorted string as a key.

6. Append the original word into the corresponding group.

7. After processing all words, collect all hashmap values into the result.

8. Return the result.

---

# Dry Run

## Input

```go
[]string{"eat", "tea", "tan", "ate", "nat", "bat"}
```

---

### Step 1

Word:

```text
eat
```

Sorted:

```text
aet
```

Map:

```go
{
    "aet": ["eat"]
}
```

---

### Step 2

Word:

```text
tea
```

Sorted:

```text
aet
```

Map:

```go
{
    "aet": ["eat", "tea"]
}
```

---

### Step 3

Word:

```text
tan
```

Sorted:

```text
ant
```

Map:

```go
{
    "aet": ["eat", "tea"],
    "ant": ["tan"]
}
```

---

### Step 4

Word:

```text
ate
```

Sorted:

```text
aet
```

Map:

```go
{
    "aet": ["eat", "tea", "ate"],
    "ant": ["tan"]
}
```

---

### Step 5

Word:

```text
nat
```

Sorted:

```text
ant
```

Map:

```go
{
    "aet": ["eat", "tea", "ate"],
    "ant": ["tan", "nat"]
}
```

---

### Step 6

Word:

```text
bat
```

Sorted:

```text
abt
```

Map:

```go
{
    "aet": ["eat", "tea", "ate"],
    "ant": ["tan", "nat"],
    "abt": ["bat"]
}
```

---

### Final Result

```go
[
    ["eat", "tea", "ate"],
    ["tan", "nat"],
    ["bat"]
]
```

---

# Complexity Analysis

Assume:

```text
n = number of strings
k = average length of each string
```

---

## Time Complexity

For every string:

### Sorting

```text
O(k log k)
```

For all strings:

```text
O(n × k log k)
```

---

## Space Complexity

HashMap stores all strings:

```text
O(n × k)
```

---

# Edge Cases

## Empty Input

### Input

```go
[]
```

### Output

```go
[]
```

---

## Single String

### Input

```go
["abc"]
```

### Output

```go
[
    ["abc"]
]
```

---

## All Strings Are Same

### Input

```go
["abc","abc","abc"]
```

### Output

```go
[
    ["abc","abc","abc"]
]
```

---

## No Anagrams

### Input

```go
["abc","def","ghi"]
```

### Output

```go
[
    ["abc"],
    ["def"],
    ["ghi"]
]
```

---

## Empty Strings

### Input

```go
["","",""]
```

### Output

```go
[
    ["","",""]
]
```

---

# Interview Discussion Points

A senior engineer should discuss:

## Solution Progression

1. Brute Force
2. HashMap Grouping
3. Sorting-Based Key
4. Frequency-Based Optimization

---

## Why HashMap?

HashMap allows constant-time insertion and lookup.

```text
Average Complexity = O(1)
```

This makes grouping efficient.

---

## Why Sorting Works?

Anagrams contain identical characters.

After sorting:

```text
eat -> aet
tea -> aet
ate -> aet
```

All produce the same representation.

---

## Why Frequency Counting Can Be Faster?

Instead of sorting:

```text
eat
```

Create frequency count:

```text
a=1
e=1
t=1
```

Use the frequency array as the key.

This avoids sorting completely.

---

# Common Follow-Up Questions

## Can We Do Better Than O(n × k log k)?

Yes.

Use character frequency counting instead of sorting.

---

## Frequency Count Complexity

For each word:

```text
O(k)
```

Total:

```text
O(n × k)
```

which is optimal for lowercase English letters.

---

## Why Is Frequency Counting Faster?

Sorting requires:

```text
O(k log k)
```

Frequency counting requires:

```text
O(k)
```

because each character is visited exactly once.

---

## What If Strings Contain Unicode Characters?

The fixed-size array approach no longer works.

Possible solutions:

* Sort runes
* Use `map[rune]int`

---

## Which Solution Should Be Given First In Interviews?

Start with:

```text
Sorting + HashMap
```

because it is simpler and easier to explain.

Then mention:

```text
Frequency Count + HashMap
```

as an optimization.

This demonstrates problem-solving progression and optimization skills expected from a Senior Golang Engineer.

---

# Summary

The key insight is that all anagrams share the same canonical representation.

Sorting-based solution:

```text
Time  : O(n × k log k)
Space : O(n × k)
```

Frequency-count solution:

```text
Time  : O(n × k)
Space : O(n × k)
```

For senior-level interviews, always explain both approaches and discuss the trade-offs between simplicity and optimal performance.
