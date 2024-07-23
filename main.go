package main

import (
	"fmt"
	"net/http"
	// "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/", controller.helloHandler)
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on
}
