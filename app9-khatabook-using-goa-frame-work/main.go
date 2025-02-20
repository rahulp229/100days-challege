package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func OperationList() {
	fmt.Println("1. Add an entry")
	fmt.Println("2. List all entries")
	fmt.Println("3. Exit")
}
func main() {
	fmt.Println("Welcome to the Khatabook")
	fmt.Println("1. Add an entry")
	fmt.Println("2. List all entries")
	fmt.Println("3. Exit")

	var Choice int
	fmt.Scanln(&Choice)
	exit := false

	dailyEntries := make(map[string]interface{})
	for Choice != 4 {
		switch Choice {
		case 1:
			fmt.Println("**Add an entry**")
			fmt.Println(" ")
			fmt.Println("Enter name")
			var name string
			fmt.Scanln(&name)
			dailyEntries["name"] = name
			fmt.Println("Enter amount")
			var amount float64
			fmt.Scanln(&amount)
			dailyEntries["amount"] = amount
			ds, err := addDataToCSV(dailyEntries)
			if err != nil {
				fmt.Println(err)
			} else {
				if ds != "" {
					fmt.Println("Data saved to", ds)
				}
			}
			OperationList()
			fmt.Scanln(&Choice)
		case 2:
			fmt.Println("List all entries")
			listAllEntries()
			OperationList()
			fmt.Scanln(&Choice)
		case 3:
			fmt.Println("Exit")
			if !exit {
				fmt.Println("Thank you for using the ATM")
				break
			}
			exit = true
		}
	}
}
func addDataToCSV(dailyEntries map[string]interface{}) (string, error) {
	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	for key, value := range dailyEntries {
		record := []string{key, fmt.Sprint(value)}
		if err := csvWriter.Write(record); err != nil {
			return "", err
		}
	}

	return "data.csv", nil
}

func listAllEntries() {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(records)
}
