package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello World")
	wg := sync.WaitGroup{}
	ch := make(chan bool)
	wg.Add(10)
	for i := 1f; i <= 10; i++ {
		go task(i, ch, &wg)
	}
	go func() {
		wg.Wait()

		close(ch)
	}()
	go func() {
		for signal := range ch {
			fmt.Println(signal)
			if signal {
				// Handle the signal here
				ch <- false
			}
		}
	}()

	time.Sleep(time.Second * 2)
}

func task(i int, ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
	ch <- true
}
