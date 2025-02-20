package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// Producer consumer problem: The producer is streaming tweets,
// the consumer needs to filter out tweets which contains the keyword "golang".
// Constraints: The tweets need to be processed in concurrently,
// however the concurrency needs to be controlled through a worker pool, that is, at any point there should be n workers running.
func main() {
	fmt.Println("hello")
	tweets := make(chan string)
	var wg = sync.WaitGroup{}
	wg.Add(50)
	var result []string
	var str string

	//worker pool
	for i := 1; i <= 50; i++ {
		go func() {

			//	tweets <- "hello golang" + " " + strconv.Itoa(i)
			tweets <- getTweets()
			wg.Wait()
		}()
		str = <-tweets
		if !strings.Contains(str, "golang") {
			result = append(result, str)
		}

	}
	if len(result) != 0 {
		os.Exit(1)
	}


}
func getTweets() string {
	tweet := time.Now().String() + " " + "golang"
	//strings.Ra
	return tweet
}
