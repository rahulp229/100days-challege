package main

import (
	"fmt"
	"sync"
)
//Problem statement = trafic signal
func main() {
	fmt.Println("Welcome day21 #100dayscodingchalenge...!")
	ts := make(map[string]string)
	ts["green"] = "go"
	ts["red"] = "stop"

	var wg sync.WaitGroup
	wg.Add(1)

	single := make(chan string, 2)
	single <- "red"

	go func() {
		defer wg.Done()

		for k, v := range ts {
			fmt.Println(k, v)
			//taking up the signal and perform operation accordingly
			select {
			case signal := <-single:
				if signal == v {
					fmt.Println("if single is green the all vehical are performing - go -Operation")
				}
			default:
				fmt.Println("Channel is empty")
			}
		}
	}()

	wg.Wait()
	fmt.Println("Operation perform succesfully...!")
}