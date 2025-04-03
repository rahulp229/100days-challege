package main
 
import "fmt"
import "github.com/gin-gonic/gin"

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Request from", c.ClientIP())
		c.Next()
	}
}

func main() {
	fmt.Println("Hello world...!")
	//create rest api using gin framwork

	route := gin.Default()
	route.Use(LoggerMiddleware())
	route.GET("/ping",func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})

	route.Run(":8080")

}