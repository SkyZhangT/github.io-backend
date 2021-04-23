package handler

import (
	"net/http"

	"app/database"

	"github.com/gin-gonic/gin"
)

  
func PostGetID(db database.DBInterface) gin.HandlerFunc{
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