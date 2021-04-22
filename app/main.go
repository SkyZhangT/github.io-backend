package main

import (
	"context"
	"fmt"

	"app/config"
	"app/database"
	"app/handler"

	"github.com/gin-gonic/gin"
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
	server.Use(handler.CORSMiddleware())

	server.GET("/post", handler.PostGet(db))
	server.POST("/post", handler.PostPost(db))
	server.DELETE("/post/:id", handler.PostDeleteID(db))
	server.GET("/post/:id", handler.PostGetID(db))
	server.GET("/image/:id", handler.ImageIDGet(conf.Img_dir))
	server.POST("/image", handler.ImagePost(conf.Img_dir))


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