package main

import (
	"fmt"
)

func main() {

	input := []int{1, 2, 3, 2, 5}

	result := containsDuplicate(input)

	fmt.Print(result)
}

func containsDuplicate(arr []int) bool {

	var answer bool

	//answer = bruteForceSolution(arr)

	answer = optimalSolution(arr)
	return answer
}

// Brute Force Solution - O(n^2) time complexity and O(1) space complexity
func bruteForceSolution(arr []int) bool {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

// optimalSolution - O(n) time complexity and O(n) space complexity
func optimalSolution(arr []int) bool {

	// Create a hash set to store the unique numbers we have seen so far.
	hashSet := make(map[int]struct{})

	for _, num := range arr {
		if _, found := hashSet[num]; found {
			return true
		}
		hashSet[num] = struct{}{}
	}

	return false
}
