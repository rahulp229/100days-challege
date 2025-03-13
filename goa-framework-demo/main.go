// goa framework demo
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	startServer()
}

func startServer() {
	// Start the server
	http.ListenAndServe(":8080", nil)
}
