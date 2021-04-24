package handler

import (
	"fmt"
	"net/http"

	"app/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type like struct {
	Inc		bool	`json:"inc"`
}

  
func LikePutID(db database.DBInterface) gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")

		payload := like{}
		err := c.BindJSON(&payload)
		if err != nil {
			fmt.Println("Bind fail")
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Update(id, bson.D{{Key: "$inc", Value: bson.D{{Key: "likes", Value: 1}}}})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
	