package main

import "fmt"

func main(){
	fmt.Println("Good evening ...")
	//create slice 
	//reverse slice
	input :=[]string{"a","b","c","d"}
	// a -> b -> c -> d none
	// d -> c -> b -> a none
	for i:=len(input);i >=0;i--{
		fmt.Println(i)
	}
	
}