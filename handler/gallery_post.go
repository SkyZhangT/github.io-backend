package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.io-backend/database"
	"github.io-backend/models"
)

type galleryPostRequest struct {
	Title   string   `json:"title"`
	User    string   `json:"user"`
	Time    string   `json:"time"`
	Text    string   `json:"text"`
	Picture []string `json:"picture"`
}

func GalleryPost(db database.DBInterface) gin.HandlerFunc{
	return func(c *gin.Context){
		item := models.Item{}
		err := c.BindJSON(&item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Insert(item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}