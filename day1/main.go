// write program to sort array/slice using built in package sort and without sort
package main

import (
	"fmt"
	"sort"
)

func main() {
	//array sort with out built in package sort
	numbers := []int{40, 60, 30, 10, 100}
	result := []int{}
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
		result = append(result, numbers[i])
	}
	fmt.Println("sorted slice ", result)

	// sort slice using sort.Slice
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] > numbers[j] })

	fmt.Println("using sort.Slice ", numbers)

}
