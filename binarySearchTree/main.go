package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Hi there")
//	var i int
	//for i=0;i>=5;i++{
		fmt.Println("Enter the number")
		time.Sleep(5)
		number,_ := fmt.Scanln()
		time.Sleep(5)
		fmt.Println(number)
	os.Exit(0)
	//}
}