package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func main() {
	r := gin.Default()

	r.GET("/user", userInfo)

	r.Run(":8080")
}

func userInfo(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "user created at this time",
	})
}
