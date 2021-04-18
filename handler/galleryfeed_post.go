package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.io-backend/dummyData"
)

type galleryPostRequest struct {
	Title   string   `json:"title"`
	User    string   `json:"user"`
	Time    string   `json:"time"`
	Text    string   `json:"text"`
	Picture []string `json:"picture"`
}

func GalleryPost(data *dummyData.Posts) gin.HandlerFunc{
	return func(c *gin.Context){
		requestBody := galleryPostRequest{}
		c.Bind(&requestBody)

		item := dummyData.Item{
			Title 	: requestBody.Title,
			User   	: requestBody.User,
			Time    : requestBody.Time,
			Text    : requestBody.Text,
			Picture : requestBody.Picture,
		}
		data.Add(item)
		
		c.Status(http.StatusAccepted)
	}
}