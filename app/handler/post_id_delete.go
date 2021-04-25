package handler

import (
	"net/http"

	"app/database"

	"github.com/gin-gonic/gin"
)

  
func PostDeleteID(db database.DBInterface, token string) gin.HandlerFunc{
	return func(c *gin.Context){
		auth := c.GetHeader("Authorization")
		if auth != token {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authentication"})
			return
		}

		id := c.Param("id")
		res, err := db.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}