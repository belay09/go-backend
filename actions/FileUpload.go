package actions

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)


func FileUpload(c *gin.Context) {
	type reqBody struct{
		 Input struct {
		Name      string `json:"name"`
		Type      string `json:"type"`
		Base64Str string `json:"base64str"`
	}
}
var input reqBody
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON input",
		})
		return
	}

	result, err := uploadImageToCloudinary(input.Input.Base64Str)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"image_url": result.SecureURL,
	})
}

func uploadImageToCloudinary(base64str string) (*uploader.UploadResult, error) {
	cloudinaryURL := "cloudinary://your-cloud-name:your-api-key@your-api-secret"
	cloudinary.ConfigFromURL(cloudinaryURL)

	uploadParams := uploader.UploadParams{
		File:      uploader.Base64File(base64str),
		UploadPreset: "your-upload-preset",
	}

	result, err := uploader.Upload(uploadParams)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
