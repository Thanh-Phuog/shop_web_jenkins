package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api/v1/products", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "get all products",
		})
	})
	r.Run(":3000")
}
