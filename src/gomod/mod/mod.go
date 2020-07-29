package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lroolle/hello"
)

// import "gin"

// func main() {
// 	 hello.Hello()
// }

func main() {
	hello.Hello()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
