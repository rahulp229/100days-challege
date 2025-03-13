package design
import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// ...
 


type Calculator interface {
	Add(x, y int) int
	Subtract(x, y int) int
	Multiply(x, y int) int
	Divide(x, y int) int
}

 
 type calculator struct{}

 func (c *calculator) Add(x, y int) int {
	return x + y
 }

 func (c *calculator) Subtract(x, y int) int {
	return x - y
 }

 func (c *calculator) Multiply(x, y int) int {
	return x * y
 }

 func (c *calculator) Divide(x, y int) int {
	return x / y
}

func main() {
	calcSvc := Service("calculator", func() { ... })
addMethod := Method("add", func() { ... })
calcSvc.Method("add", addMethod)
	// ...
	var _ = Service("calculator", func() {             // A service groups related methods
		Method("add", func() {                          // A method (endpoint)
			 Payload(func() {                           // The method's payload (request body)
				 Attribute("a", Int, "First number")    
				 Attribute("b", Int, "Second number")
				 Required("a", "b")                     // Payload validation
			 })
			 Result(Int)                                // The method's result (response body)
	 
			 HTTP(func() {                              // HTTP transport details
				 GET("/add/{a}/{b}")                    // HTTP request verb and path
				 Response(StatusOK)                     // HTTP response status
			 })
		 })
	 })
	// ...}

}	