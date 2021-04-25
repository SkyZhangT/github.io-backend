package main

import (
	"context"
	"fmt"

	"app/config"
	"app/database"
	"app/handler"

	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConfig()
	fmt.Println(conf)	
	ctx := context.TODO()

	db := database.Initdb(ctx, conf.Mongo)
	token := db.GetToken()
	fmt.Println(token)

	db.Printdb()

	lmt := handler.Limiter(conf.Limiter)

	server := gin.Default()

	// server.Use(Authorization(token))
	// CORS middleware
	server.Use(handler.CORSMiddleware())
	// request limiter middleware
	server.Use(tollbooth_gin.LimitHandler(lmt))

	server.GET("/post", handler.PostGet(db))
	server.POST("/post", handler.PostPost(db, token))
	server.DELETE("/post/:id", handler.PostDeleteID(db, token))
	server.GET("/post/:id", handler.PostGetID(db))
	server.POST("/image", handler.ImagePost(conf.Img_dir, token))
	server.PUT("/like/:id", handler.LikePutID(db))
	
	server.Static("/images", "/images/")


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