package main

import "fmt"

func main() {

	//results := fibWithRecursion(10)
	results := fibWithDP(10)

	fmt.Print(results)
}

func fibWithRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return recursion(n)
}

// recursion Solution - O(2^n) time complexity and O(n) space complexity
func recursion(n int) int {
	return fibWithRecursion(n-1) + fibWithRecursion(n-2)
}

func fibWithDP(n int) int {

	cache := make(map[int]int, n)
	answer := dynamicProgramming(n, cache)

	return answer
}

// Dynamic Programming Solution - O(n) time complexity and O(n) space complexity
func dynamicProgramming(n int, cache map[int]int) int {

	if n <= 1 {
		return n
	}

	if val, ok := cache[n]; ok {
		return val
	}

	cache[n] = dynamicProgramming(n-1, cache) + dynamicProgramming(n-2, cache)

	return cache[n]
}
