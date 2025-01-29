package main

import (
	"fmt"
)

// Note that this code uses the rune type to represent the characters in the string, rather than the byte type. This is because the byte type is not sufficient to represent all Unicode characters, whereas the rune type is.

// Also, this code uses the len function to get the length of the string, and the string function to convert the slice of runes back to a string.
func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	s := "hello"
	fmt.Println("Original string:", s)
	fmt.Println("Reversed string:", reverse(s))

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	s2 := append(s1, 6, 7, 8)
	fmt.Println(s2)
	
	fmt.Println("length of s1 | capacity of s1", len(s1), cap(s1))
	fmt.Println("length of s2 | capacity of s2", len(s2), cap(s2))
	s3 := append(s2, 9, 10)
	fmt.Println(s3)
	fmt.Println("length of s3 | capacity of s3", len(s3), cap(s3))
	generateRandomNumber(1, 6)
	
}
func generateRandomNumber(min, max int) int {
	seed := time.Now().UnixNano() / 123456789
	rand.New(rand.NewSource(int64(seed)))
	fmt.Println("random string - ", rand.Intn(max-min)+min)
	return rand.Intn(max-min) + min
}
