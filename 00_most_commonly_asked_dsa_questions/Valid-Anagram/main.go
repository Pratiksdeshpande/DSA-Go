package main

import (
	"fmt"
)

func main() {
	a, b := "Listen&%", "%siL&ent"

	fmt.Println(validAnagram(a, b))
}

func validAnagram(a, b string) bool {

	if len(a) != len(b) {
		return false
	}

	hashMap := map[rune]int{}

	for _, c := range a {
		hashMap[c]++
	}

	for _, c := range b {
		if hashMap[c] == 0 {
			return false
		}
		hashMap[c]--
	}

	return true
}
