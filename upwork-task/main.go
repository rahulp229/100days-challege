package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	defer sdl.Quit()

	runtime.LockOSThread()

	flag.Parse()

	var err error
	if err = sdl.Init(uint32(sdl.INIT_EVERYTHING)); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	handler := gin.HandlerFunc(func(c *gin.Context) {
		tokenString, err := getToken()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}  
		fmt.Println(tokenString) 
	})
	r.GET("/login",handler, func(c *gin.Context) {
		*gin.HandlerFunc(func(c *gin.Context) {
			tokenString, err := getToken()
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}	
		})
		tokenString, err := getToken()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"token": tokenStringt,
		})
	})

	r.Run(":8080")
}

func getToken() (string, error) {
	const secretKey = "my-secret-key"

	claims := jwt.MapClaims{
		"sub":  "1234567890",
		"name": "John Doe",
		"exp":  time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
