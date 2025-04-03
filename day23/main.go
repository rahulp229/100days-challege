package main

import "fmt"

func main(){
	fmt.Println("Hellow World...!")
	//input = [1,1,1,2,2,3,4,4,4,4,5,5,5,5,5]
	//outpu = count of each number[1=3,2=2,3=1,4=4,5=5]
	num := []int{1,1,1,2,2,3,4,4,4,4,5,5,5,5}
	result := make(map[int]int)
	var(
		 countJ1 int
	 	 countJ2 int
	     countJ3 int
	     countJ4 int
	     countJ5 int
	     countJ6 int
	     countJ7 int
	     countJ8 int
	     countJ9 int
	)
	


	for i:=0;len(num)<=1000;i++{
		 if len(i) == 1{
		 	for j:=1;j<10;j++{
		 		if j == 1{
		 			fmt.Println(countJ1)
		 			result[i]=counJ1+1

		 		}
		 		if j == 2{
		 			fmt.Println(countJ2)
		 			result[i]=countJ2+1

		 		}
		 		if j == 3 {
		 		fmt.Println(countJ3)

		 		result[i]=countJ3+1

		 		}
		 		if j == 4 {
		 		fmt.Println(countJ4)

		 		result[i]=countJ4+1

		 		}
		 		if j == 5 {
		 		fmt.Println(countJ5)

		 			result[i]=countJ5+1

		 		}
		 		if j == 6 {
		 		 fmt.Println(countJ6)

		 			result[i]=countJ6+1

		 		}
		 		if j == 7 {
		 		fmt.Println(countJ7)

		 			result[i]=countJ7+1

		 		}
		 		if j == 8 {
		 		fmt.Println(countJ8)

		 			result[i]=countJ8+1

		 		}
		 		if j == 9 {
				fmt.Println(countJ9)

		 		result[i]=countJ9+1
	
		 		}

		 	}	
		 }
	}
	fmt.Println("result=",result)
}