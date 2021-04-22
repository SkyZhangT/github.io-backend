package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

  
func ImageIDGet(path string) gin.HandlerFunc{
	return func(c *gin.Context){

		c.JSON(http.StatusOK, gin.H{"message": "not implemented"})
	}
}