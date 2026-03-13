package main

import (
	"fmt"
	"math"
)

type Input struct {
	arr    []int
	target int
}

func main() {

	input := Input{arr: []int{1, 2, 3, 1}, target: 3}

	result := containsDuplicate(input.arr, input.target)

	fmt.Print(result)
}

func containsDuplicate(arr []int, target int) bool {
	var answer bool

	//answer = bruteForceSolution(arr, target)
	answer = optimalSolution(arr, target)
  
	return answer
}

// Brute Force Solution - O(n^2) time complexity and O(1) space complexity
func bruteForceSolution(arr []int, target int) bool {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < int(math.Min(float64(arrLen), float64(i+1+target))); j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

// optimalSolution uses a hash map to track the most recent index of each number.
// For each element, it checks if the same number was seen within the target distance.
// Time Complexity: O(n) - single pass through the array with O(1) hash map operations.
// Space Complexity: O(n) - hash map stores at most n elements.
func optimalSolution(arr []int, target int) bool {

	// Create a hash map to store the unique numbers we have seen so far.
	hashMap := make(map[int]int)

	for i, num := range arr {
		if _, found := hashMap[num]; found && i-hashMap[num] <= target {
			return true
		}
		hashMap[num] = i
	}

	return false
}

// optimalSolution2 uses a sliding window approach with a hash set.
// Maintains a window of at most 'target' elements and checks for duplicates within it.
// Time Complexity: O(n) - single pass through the array with O(1) hash set operations.
// Space Complexity: O(min(n, target)) - hash set stores at most 'target' elements.
func optimalSolution2(arr []int, target int) bool {
	window := make(map[int]bool)
	L := 0

	for R := 0; R < len(arr); R++ {
		if R-L > target {
			delete(window, arr[L])
			L++
		}
		if window[arr[R]] {
			return true
		}
		window[arr[R]] = true
	}

	return false
}
