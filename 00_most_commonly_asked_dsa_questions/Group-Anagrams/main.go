package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(validGroupAnagram(s))
}

func validGroupAnagram(a []string) [][]string {

	hashMap := map[string][]string{}

	for _, word := range a {

		chars := []rune(word)

		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		key := string(chars)
		hashMap[key] = append(hashMap[key], word)
	}

	result := make([][]string, 0, len(hashMap))

	for _, group := range hashMap {
		result = append(result, group)
	}

	return result
}
