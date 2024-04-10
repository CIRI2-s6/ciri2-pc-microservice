// Package utils holds utility functions that can be used across the application
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BindJSONAndValidate binds the JSON request body to the input interface and validates it
func BindJSONAndValidate(c *gin.Context, input interface{}) bool {
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	// Additional validation logic can be added here
	return true
}
