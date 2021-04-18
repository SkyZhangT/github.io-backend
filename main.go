package main

import (
	"fmt"

	"github.io-backend/config"
	"github.io-backend/database"
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
	fmt.Print(conf)

	db := database.Initdb(conf.Mongo)

	db.Printdb()

	// posts := dummyData.New()
	// server := gin.Default()

	// server.GET("/gallery/feed", handler.GalleryGet(posts))
	// server.POST("/gallery/feed", handler.GalleryPost(posts))

	// server.Run()
}