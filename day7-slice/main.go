package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{6, 1, 1, 2, 3, 3, 4, 5}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(s)
	fmt.Println("----------")
	Str := "india is my country"
	result, _ := countOfVowels(Str)
	fmt.Println(result)
	fmt.Println(removeDuplicates(s))
	s1 := []string{"red", "blue", "green", "yellow", "orange", "red"}
	fmt.Println(removeDuplicatesString(s1))
	s3 := []int{30, 10, 60, 20, 50}
	fmt.Println(sortSlice(s3))
}
func countOfVowels(input string) (map[string]int, error) {
	vowelCounts := make(map[string]int)
	for _, char := range input {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			vowelCounts[string(char)]++
		}
	}
	return vowelCounts, nil
}

func removeDuplicates(s []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, item := range s {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

func removeDuplicatesString(s []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, item := range s {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

func sortSlice(s []int) []int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}
