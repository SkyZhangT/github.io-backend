package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"app/database"

	"github.com/gin-gonic/gin"
)

  
func PostGet(db database.DBInterface) gin.HandlerFunc{
	return func(c *gin.Context){
		page := c.Param("p")
		if page == "" {
			page = "0"
		}

		n, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			fmt.Printf("Int64 conversion failed. value: %d", n)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.GetPage(n)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}