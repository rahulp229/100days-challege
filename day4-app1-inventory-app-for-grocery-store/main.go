package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	operationsFile = "operations.csv"
	invoiceFile    = "invoice.csv"
	loadItemsFile  = "loadItems.csv"
	dataFile       = "data.csv"
)

var globalInvoiceID int
var operations = []string{
	"AddItem",
	"CreateInvoice",
	"GetInvoice",
	"DeleteItem",
	"UpdateItem",
	"GetItem",
	"GetAllItems",
	"Placeorder",
	"GetAllOrders",
	"GetAllInvoices",
	"GetAllCustomers",
	"Stop",
}

type Invoice struct {
	OrderID              int
	Items                []Item
	CustomerMobileNumber string
	InvoiceID            int
	InvoiceDate          string
	TotalAmount          float64
}

type Customer struct {
	ID                   int
	CustomerName         string
	CustomerMobileNumber string
}

type Item struct {
	ID    int
	Name  string
	Price float64
}

type Order struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func OrderPlaced() *Order {

	return nil
}
func main() {
	printWelcomeMessage()
	printOperations()
	for {
		input := readInput()
		if input >= 10 {
			break
		}
		switch input {
		case 0:
			AddItemToCSV()
		case 1:
			CreateInvoice()
		case 2:
			//DeleteItem()
		case 3:
			//UpdateItem()
		case 4:
			//GetItem()
		case 5:
			//GetAllItems()
		case 6:
			//OrderPlaced()
		case 7:
			//GetAllOrders()
		case 8:
			//GetAllInvoices()
		case 9:
			//GetAllCustomers()
		}
	}
}

func printWelcomeMessage() {
	fmt.Println("Welcome to Bharat Grocery System:")
	time.Sleep(time.Second * 1)
	fmt.Println("Loading system...")
	time.Sleep(time.Second * 1)
	fmt.Println("successfully loaded...")
	time.Sleep(time.Second * 1)
}

func printOperations() {
	for i, op := range operations {
		fmt.Printf("%d. %s\n", i, op)
	}
}

func readInput() int {
	var input int
	fmt.Print("Enter operation number: ")
	fmt.Scanln(&input)
	return input
}

func OpenFile(filename string) (*os.File, [][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	return file, records, nil
}

func AddItemToCSV() {
	lfile, lrecords, err := OpenFile(loadItemsFile)
	if err != nil {
		fmt.Printf("Error loading file: %v\n", err)
		panic(err)
	}
	defer lfile.Close()
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(0)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range lrecords {
		if err := writer.Write(row); err != nil {
			log.Printf("error writing record to csv: %v", err)
			break
		} else {
			time.Sleep(time.Second * 2)
			fmt.Println("Loading operations in progress...")
			for i := 0; i < 10; i++ {
				if i == 9 {
					//fmt.Println("100%")
					time.Sleep(time.Second * 2)
					fmt.Println("Successfully loaded data to csv file")
				}

			}
		}
	}
}

func CreateInvoice() {
	time.Sleep(time.Second * 2)
	fmt.Println("Get Invoice")
	invoiceFile, records, err := OpenFile(invoiceFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		panic(err)
	}
	defer invoiceFile.Close()
	lastRecord := len(records) - 1
	globalInvoiceID, _ = strconv.Atoi(records[lastRecord][0])
	globalInvoiceID++
	fmt.Println("New invoice id - > ", globalInvoiceID)
	//
	var invoice Invoice
	for recordKey, record := range records {
		fmt.Println(record)
		var invoiceInput int
		fmt.Print("Add item in invoice - ")
		fmt.Scanln(&invoiceInput)
		invoice.OrderID = invoiceInput
		invoice.Items = []Item{}
		for {
			var itemInput int
			var itemName string
			var itemPrice float64
			fmt.Print("Item ID - ")
			fmt.Scanln(&itemInput)
			fmt.Print("Item Name - ")
			fmt.Scanln(&itemName)
			fmt.Print("Item Price - ")
			fmt.Scanln(&itemPrice)
			item := Item{
				ID:    itemInput,
				Name:  itemName,
				Price: itemPrice,
			}
			invoice.Items = append(invoice.Items, item)
			fmt.Print("Add more items - 0/1 - ")
			fmt.Scanln(&invoiceInput)
			if invoiceInput == 0 {
				break
			}
		}
		invoice.InvoiceID = globalInvoiceID
		invoice.InvoiceDate = time.Now().Format("2006-01-02 15:04:05")
		invoice.TotalAmount = 0
		for _, item := range invoice.Items {
			invoice.TotalAmount += item.Price
		}
		fmt.Println("Invoice - ", invoice)
		records = append(records, []string{strconv.Itoa(invoice.InvoiceID),
			strconv.Itoa(invoice.OrderID),
			strconv.Itoa(invoice.InvoiceID),
			invoice.InvoiceDate,
			strconv.FormatFloat(invoice.TotalAmount, 'f', -1, 64),
		})
		writer1 := csv.NewWriter(invoiceFile)
		defer writer1.Flush()
		if err := writer1.Write(records[recordKey]); err != nil {
			log.Printf("error writing record to csv: %v", err)
			break
		} else {
			time.Sleep(time.Second * 2)
			fmt.Println("SUCCESS...!")
			fmt.Println("Loading operations...")
			time.Sleep(time.Second * 2)
		}
	}
}

func GetInvoice() {
	fmt.Println("Get Invoice")
	invoiceFile, records, err := OpenFile(invoiceFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		panic(err)
	}
	defer invoiceFile.Close()
	fmt.Println(records)
}
