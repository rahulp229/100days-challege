// write program to sort array/slice using built in package sort and without sort
package main

import (
	"fmt"
	"math/rand"
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
	colors := []string{"red", "blue", "green", "yellow", "orange"}
	sort.Strings(colors)
	fmt.Println("using sort.Strings ", colors)

	// sort slice using sort.SliceStable
	sort.SliceStable(colors, func(i, j int) bool { return colors[i] > colors[j] })
	fmt.Println("using sort.SliceStable ", colors)
	rn := generateRandomNumber(1, 6)
	fmt.Println("random number is ->", rn)
	input := "rahul"
	reversed := reverseString(input)
	fmt.Println("reverse string is ->", reversed)

}

func generateRandomNumber(min, max int) int {
	seed := 123456
	rand.New(rand.NewSource(int64(seed)))
	return rand.Intn(max-min) + min
}

func reverseString(s string) string {
	runes := []rune(s) // rune is alias for int32
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
