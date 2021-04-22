package handler

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

  
func ImagePost(path string) gin.HandlerFunc{
	return func(c *gin.Context){
		c.Request.ParseMultipartForm(30)

		file, _, err := c.Request.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		hash := md5.New()
		hash.Write(fileBytes)
		md5String := hex.EncodeToString(hash.Sum(nil))

		newpath := path
		for i := 0; i < 3; i++ {
			newpath = filepath.Join(newpath, md5String[2*i:2*(i+1)])

			if _, err := os.Stat(newpath); os.IsNotExist(err) {
				err = os.Mkdir(newpath, 0700)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
					return
				}
			}
		}

		newpath = filepath.Join(newpath, md5String + ".jpg")
		if _, err := os.Stat(newpath); err == nil {
			c.JSON(http.StatusAlreadyReported, md5String + ".jpg")
			return
		}

		outputFile, err := os.Create(newpath)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		defer outputFile.Close()

		outputFile.Write(fileBytes)


		c.JSON(http.StatusOK, md5String + ".jpg")
	}
}