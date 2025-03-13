package main

import "fmt"

func main() {
	patternResults := []int{}
	for row := 0; row < 5; row++ {
		if row == 0 {
			fmt.Println("*")
		} else {
			patternResults = append(patternResults, row+1)
			for _, count := range patternResults {
				for star := 0; star < count; star++ {
					fmt.Print("*")
				}
				fmt.Println()
			}
		}
	}
}

