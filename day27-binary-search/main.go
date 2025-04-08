package main

import (
	"fmt"
)

func main() {
	fmt.Println("Binary Search")
	var pivotpoint int = 0
	num := []int{60, 20, 30, 40, 50, 10, 70, 80, 90}
	var result []int
	fmt.Println("len=", len(num))
	//finding pivot point
	for i := 0; i < len(num)/2; i++ {
		pivotpoint = num[i]
		if pivotpoint < num[i+1] {
			//go-staticcheck
			result = append(result, num[i])
			pivotpoint = num[i+1]
		} else if pivotpoint > num[i+1] {
			result = append(result, num[i])
		}
	}
	fmt.Println("pp=", pivotpoint)
	left := num[:4]
	right := num[4:]
	fmt.Println("left=", left, "len of left = ", len(left))
	fmt.Println("right=", right)
	fmt.Println("=========================")
	for j := 0; j < len(left); j++ {
		//left
		if left[j] < pivotpoint {
			result = append(result, left[j])
		}
		if left[j] > pivotpoint {
			result = append(result, right[j])
		}
	}
	fmt.Println("right result = ", right)
	for k := 0; k < len(right); k++ {
		//right
		if right[k] < pivotpoint {
			result = append(result, left[k])
		}
		if right[k] > pivotpoint {
			result = append(result, right[k])
		}
	}
	fmt.Println("Final Result = ", result)
}
