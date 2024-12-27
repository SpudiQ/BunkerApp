package main

import (
	"backend/internal/api"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
