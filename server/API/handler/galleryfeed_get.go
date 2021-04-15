package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.io-backend/server/API/dummyData"
)

  
func GalleryGet(data *dummyData.Posts) gin.HandlerFunc{
	return func(c *gin.Context){

		c.JSON(http.StatusOK, data.Get())
	}
}