package main

import (
	"fmt"
	"sort"
)	
func main()  {
	s := []int{1, 2, 3, 4, 5}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(s)
}