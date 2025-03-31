package main

import "fmt"


type escalator interface {
	Up()
	Down()
}

type Escalator struc{}

func main(){
	fmt.Println("Good after noon")
	//create slice 
	//reverse slice
	input :=[]string{"a","b","c","d"}
	// a -> b -> c -> d none
	// d -> c -> b -> a none
	for i:=len(input);i>=0;i--{
		fmt.Println(i)
	}
}