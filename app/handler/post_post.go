package handler

import (
	"fmt"
	"net/http"

	"app/database"
	"app/models"

	"github.com/gin-gonic/gin"
)

func PostPost(db database.DBInterface) gin.HandlerFunc{
	return func(c *gin.Context){
		item := models.Item{}
		err := c.BindJSON(&item)
		if err != nil {
			fmt.Println("Bind fail")
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Insert(item)
		if err != nil {
			fmt.Println("Insert fail")
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}