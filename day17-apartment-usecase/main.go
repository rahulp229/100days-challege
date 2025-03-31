package main

import(
	"fmt"
)
func main(){
	fmt.Println("Hello")
	result:=[]int{}
	num := []int{10,30,20,40}
	for i:=0;i<=len(num);i++{
		if num[i] >= num[i+1]{
			result = append(result,num[i+1])
		}
	}
	fmt.Println("result")
	defer panic("exception handle")
}
