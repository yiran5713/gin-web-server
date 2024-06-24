package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Handle GET request on "/"
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the homepage!",
		})
	})

	// Handle POST request on "/submit"
	r.POST("/submit", func(c *gin.Context) {
		// Example of handling POST data
		type RequestBody struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		var requestBody RequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Serialize the received data
		data, err := json.Marshal(requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize data"})
			return
		}

		// Write to a file
		err = os.WriteFile(os.Getenv("WRITEPATH")+""+requestBody.Name+".json", data, 0644)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write data to file"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Data received successfully",
			"name":    requestBody.Name,
			"age":     requestBody.Age,
		})
	})

	// Handle PUT request on "/update"
	r.PUT("/update", func(c *gin.Context) {
		// Example of handling PUT data
		type UpdateBody struct {
			Status string `json:"status"`
		}

		var updateBody UpdateBody
		if err := c.ShouldBindJSON(&updateBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Status updated",
			"status":  updateBody.Status,
		})
	})

	// Handle DELETE request on "/delete"
	r.DELETE("/delete", func(c *gin.Context) {
		// Example of DELETE request, maybe expecting an ID
		id := c.Query("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing ID"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Record deleted",
			"id":      id,
		})
	})

	// Start the server on port 8080
	r.Run(":8080")
}
