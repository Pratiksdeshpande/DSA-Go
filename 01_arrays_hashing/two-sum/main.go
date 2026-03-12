package main

import "fmt"

type Input struct {
	arr    []int
	target int
}

func main() {

	input := Input{arr: []int{1, 4, 3, 6, 3, 1, 3}, target: 5}

	result := twoSum(input.arr, input.target)

	fmt.Print(result)
}

func twoSum(arr []int, target int) []int {

	var answer []int

	//answer = bruteForceSolution(arr, target)

	answer = optimalSolution(arr, target)

	return answer
}

// Brute Force Solution - O(n^2) time complexity and O(1) space complexity
func bruteForceSolution(arr []int, target int) []int {

	var result []int

	for i, num := range arr {
		complement := target - num

		for j := i + 1; j < len(arr); j++ {
			if arr[j] == complement {
				result = append(result, i, j)
				return result
			}
		}
	}

	return result
}

// Optimal Solution - O(n) time complexity and O(n) space complexity
func optimalSolution(arr []int, target int) []int {

	var result []int

	// Create a hash map to store the indices of the numbers we have seen so far. map[num] = index
	hashMap := make(map[int]int)

	for i, num := range arr {
		complement := target - num
		if index, found := hashMap[complement]; found {
			result = append(result, index, i)
			return result
		}
		hashMap[num] = i
	}
	return result
}
