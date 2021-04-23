package handler

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"image/jpeg"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

  
func ImagePost(path string, token string) gin.HandlerFunc{
	return func(c *gin.Context){
		auth := c.GetHeader("Authorization")
		if auth != token {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid authentication"})
			return
		}

		c.Request.ParseMultipartForm(30)

		file, handler, err := c.Request.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		defer file.Close()

		fileBytes, err := resizeImage(file)
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

		fileSuffix := ".ukn"
		switch c_type := handler.Header.Get("Content-Type"); c_type {
			case "image/jpeg":
				fileSuffix = ".jpg"
			// case "image/png":
			// 	fileSuffix = ".png"
			// case "image/gif":
			// 	fileSuffix = ".gif"
			default:
				// not supported file type, return bad request
				c.JSON(http.StatusAlreadyReported, gin.H{"message": c_type + " file type is not supported."})
				return
		}

		newpath = filepath.Join(newpath, md5String + fileSuffix)
		if _, err := os.Stat(newpath); err == nil {
			c.JSON(http.StatusAlreadyReported, newpath)
			return
		}

		outputFile, err := os.Create(newpath)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		defer outputFile.Close()

		outputFile.Write(fileBytes)

		c.JSON(http.StatusOK, newpath)
	}
}

func resizeImage(file multipart.File) ([]byte, error) {
	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}
	resized := resize.Thumbnail(1920, 1920, img, resize.Lanczos3)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, resized, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}