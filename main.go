package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	if os.Getenv("PORT") == "" {
		log.Printf("Settings PORT: 8001")
		os.Setenv("PORT", "8001")
	}
}

func main() {
	port := ":" + os.Getenv("PORT")
	log.Printf("Starting server on %s...", port)

	fmt.Println("Hello, world.")

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/events/:sport/*league", func(c *gin.Context) {
		sport := c.Param("sport")
		league := c.Param("league")
		log.Printf("Sport=%s, League=%s", sport, league)
		c.JSON(200, gin.H{
			"sport":  sport,
			"league": league,
		})
	})

	router.Run()
}
