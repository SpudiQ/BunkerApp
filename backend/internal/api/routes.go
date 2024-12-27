// internal/api/routes.go
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!"})
	})

	// r.POST("/rooms", createRoom)
	// r.GET("/rooms/:id", getRoom)
	// ...
}
