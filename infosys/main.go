package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

// Given an array of integers, return the majority element.
// The majority element is the element that appears more than n / 2 times.
// You may assume that the majority element always exists in the array.
func majorityElement(nums []int) int {
	majority := nums[0]
	count := 1

	for _, num := range nums[1:] {
		if num == majority {
			count++
		} else {
			count--
			if count == 0 {
				majority = num
				count = 1
			}
		}
	}

	return majority
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2, 1,3}
	fmt.Println(majorityElement(nums))
}