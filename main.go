package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.io-backend/config"
	"github.io-backend/database"
	"github.io-backend/handler"
)

type Item struct {
	Title   string   `json:"title"`
	User    string   `json:"user"`
	Time    string   `json:"time"`
	Text    string   `json:"text"`
	Picture []string `json:"picture"`
}

func main() {
	conf := config.GetConfig()
	fmt.Println(conf)	
	ctx := context.TODO()

	db := database.Initdb(ctx, conf.Mongo)

	db.Printdb()

	server := gin.Default()

	// server.Use(Authorization(conf.Token))

	server.GET("/gallery", handler.GalleryGet(db))
	server.POST("/gallery", handler.GalleryPost(db))
	server.DELETE("/gallery/:id", handler.GalleryDeleteID(db))
	server.GET("/gallery/:id", handler.GalleryGetID(db))

	server.Run()
	db.Close()
}

func Authorization(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if token != auth {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid authorization token"})
			return
		}
		c.Next()
	}
}