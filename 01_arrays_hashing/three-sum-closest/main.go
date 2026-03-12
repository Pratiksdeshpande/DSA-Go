package main

import (
	"fmt"
	"math"
	"sort"
)

type Input struct {
	arr    []int
	target int
}

func main() {

	input := Input{arr: []int{-1, 2, 1, -4, -8, -8, -8}, target: 1}

	results := threeSum(input.arr, input.target)

	fmt.Print(results)
}

func threeSum(arr []int, target int) int {

	var results int

	results = optimalSolution(arr, target)

	return results
}

func optimalSolution(nums []int, target int) int {

	// Sort the input array to facilitate the two-pointer approach and handle duplicates
	sort.Ints(nums)

	answer := nums[0] + nums[1] + nums[2]

	for start := 0; start < len(nums)-1; start++ {
		if start > 0 && nums[start] == nums[start-1] {
			continue
		}
		pointer := start + 1
		end := len(nums) - 1

		for pointer < end {
			sum := nums[start] + nums[pointer] + nums[end]
			if math.Abs(float64(sum-target)) < math.Abs(float64(answer-target)) {
				answer = sum
			}
			if sum < target {
				pointer++
			} else if sum > target {
				end--
			} else {
				return sum
			}
		}
	}

	return answer
}
