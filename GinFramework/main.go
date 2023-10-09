package main

import (
	"github.com/gin-gonic/gin"
	
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
			"object": gin.H{"id": 1,
				"title":              "iPhone 9",
				"description":        "An apple mobile which is nothing like apple",
				"price":              549,
				"discountPercentage": 12.96,
				"rating":             4.69,
				"stock":              94},
		})
	})

	router.Run(":8080")
}


