package main

import (
	"github.com/gin-gonic/gin"
	"github.io-backend/server/API/dummyData"
	"github.io-backend/server/API/handler"
)

func main() {
	posts := dummyData.New()
	server := gin.Default()

	server.GET("/gallery/feed", handler.GalleryGet(posts))
	server.POST("/gallery/feed", handler.GalleryPost(posts))

	server.Run()
}