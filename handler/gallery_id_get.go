package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.io-backend/database"
)

  
func GalleryGetID(db database.DBInterface) gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		res, err := db.Get(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}