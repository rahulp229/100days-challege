package main

import (
	"fmt"
)

func main() {

	balance := 10000.0
	var withdraw float64
	var deposit float64
	exit := false
	var Choice int
	operations()
	fmt.Scanln(&Choice)
	for Choice != 4 {
		switch Choice {
		case 1: //balance
			fmt.Println("Your balance is ", balance)
			operations()
			fmt.Scanln(&Choice)

		case 2: //withdraw
			fmt.Println("Enter amount to withdraw")
			fmt.Scanln(&withdraw)
			balance = balance - withdraw
			fmt.Println("Your balance is ", balance)
			operations()
			fmt.Scanln(&Choice)

		case 3: //deposit
			fmt.Println("Enter amount to deposit")
			fmt.Scanln(&deposit)
			balance = balance + deposit
			fmt.Println("Your balance is ", balance)
			operations()
			fmt.Scanln(&Choice)

		case 4: //exit
			if !exit {
				fmt.Println("Thank you for using the ATM")
				break
			}
			exit = true

		}
	}
}
func operations() {
	fmt.Println("Welcome to the ATM")
	fmt.Println("1. Balance")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Deposit")
	fmt.Println("4. Exit")
}

type ATM interface {
	Balance() float64
	Withdraw(amount float64) error
	Deposit(amount float64) error
	Exit() error
}
