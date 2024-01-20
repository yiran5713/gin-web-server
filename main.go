package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Create a Gin router
    r := gin.Default()

    // Define a handler for the GET request
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, World!",
        })
    })

    // Run the server
    r.Run(":8081") // By default, it will listen on port 8080
}
