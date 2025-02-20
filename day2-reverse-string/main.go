package main

import "fmt"

//Write a go code to take array input from user and check if it has duplicates. If it does then return true else return false.

// e.g. Input: [1,2,3,4,4] Output: true
// e.g. Input:[1,2,3,4,5] Output: false
func main() {
	//
	
	
	input := []int{1, 2, 3, 4, 4}
	var result bool
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				result = true
			}
		}
	}
	fmt.Println(result)
}
