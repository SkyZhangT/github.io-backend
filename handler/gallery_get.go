package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.io-backend/database"
)

  
func GalleryGet(db database.DBInterface) gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println(c.ClientIP())
		fmt.Println(c.RemoteIP())
		offset := c.Param("skip")
		
		if offset == "" {
			offset = "1"
		}
		n, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			fmt.Printf("Int64 conversion failed. value: %d", n)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		res, err := db.NextTen(n)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}