package main

import (
	"github.com/gin-gonic/gin"
	"github.io-backend/server/API/db-handler"
	"github.io-backend/server/API/dummyData"
	"github.io-backend/server/API/handler"
)

type Item struct {
	Title   string   `json:"title"`
	User    string   `json:"user"`
	Time    string   `json:"time"`
	Text    string   `json:"text"`
	Picture []string `json:"picture"`
}


func main() {
	client := db.Initdb()


	posts := dummyData.New()
	server := gin.Default()

	server.GET("/gallery/feed", handler.GalleryGet(posts))
	server.POST("/gallery/feed", handler.GalleryPost(posts))

	server.Run()
}