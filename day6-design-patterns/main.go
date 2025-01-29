// singleton pattern - A pattern that ensures a struct has only one instance, while providing a global point of access to that instance. In Go, you can implement the singleton pattern using a combination of a private constructor and a global variable.
package main

import (
	"fmt"
	"sync"
)

type Singleton struct{} // Define the singleton struct
var instance *Singleton // Define a global variable to hold the singleton instance

func main() {

	s := GetInstance()
	fmt.Println(s)
	//Factory pattern implementation start

}

var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() { // Execute the function only once
		fmt.Println("Creating singleton instance")
		instance = &Singleton{}
	})
	return instance
}

// ***Factory Pattern implementation start***
type Animal interface {
	Speak()
}
type Dog struct{}

func (d *Dog) Speak() {
	fmt.Println("Bark")
}

//***Factory pattern implementation end***
