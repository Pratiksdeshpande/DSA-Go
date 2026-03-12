package main

import (
	"fmt"
	"sort"
)

type Input struct {
	arr    []int
	target int
}

func main() {

	input := Input{arr: []int{1, 4, 3, 6, -3, -1, 0, 0}, target: 0}

	results := threeSum(input.arr, input.target)

	fmt.Print(results)
}

func threeSum(arr []int, target int) [][]int {

	var results [][]int

	//results = bruteForceSolution(arr, target)

	results = optimalSolution(arr, target)

	return results
}

// Brute Force Solution - O(n^3) time complexity and O(1) space complexity
func bruteForceSolution(arr []int, target int) [][]int {

	// Sort to handle duplicates
	sort.Ints(arr)

	var results [][]int
	for i := 0; i < len(arr)-2; i++ {
		// Skip duplicate values for i
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j := i + 1; j < len(arr)-1; j++ {
			// Skip duplicate values for j
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}
			for k := j + 1; k < len(arr); k++ {
				// Skip duplicate values for k
				if k > j+1 && arr[k] == arr[k-1] {
					continue
				}
				if arr[i]+arr[j]+arr[k] == target {
					results = append(results, []int{arr[i], arr[j], arr[k]})
				}
			}
		}
	}
	return results
}

// Optimal Solution - O(n^2) time complexity and O(1) space complexity
func optimalSolution(nums []int, target int) [][]int {

	// Sort the input array to facilitate the two-pointer approach and handle duplicates
	sort.Ints(nums)

	var result [][]int

	for start := 0; start < len(nums); start++ {
		if start > 0 && nums[start] == nums[start-1] {
			continue // Skip duplicate values for start
		}
		pointer, end := start+1, len(nums)-1
		for pointer < end {
			sum := nums[start] + nums[pointer] + nums[end]
			if sum == 0 {
				triplet := []int{nums[start], nums[pointer], nums[end]}
				result = append(result, triplet)
				pointer++
				end--

				// Skip duplicate values for pointer
				for pointer < end && nums[pointer] == nums[pointer-1] {
					pointer++
				}

				// Skip duplicate values for end
				for pointer < end && nums[end] == nums[end+1] {
					end--
				}
			} else if sum < 0 {
				pointer++
			} else {
				end--
			}
		}
	}

	return result
}
