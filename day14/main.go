package main

import "fmt"

// binary search
// input = 10,20,30,40
// output =
func main() {
	fmt.Println("Hello")
	//
	input := []int{10, 20, 30, 40}
	pivotpoint := 0
	var result []int
	for _, val := range input {
		pivotpoint = val
		if pivotpoint < val {
			result = append(result, val)
		} else if pivotpoint > val {
			result = append(result, val)
		}

	}
	fmt.Println(result)
}
