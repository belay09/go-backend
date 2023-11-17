package actions

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	// "bytes"
	// "encoding/base64"
)

func FileUpload(c *gin.Context) {
	type reqBody struct {
		Input struct {
			Name      string `json:"name"`
			Base64Str string `json:"base64str"`
			Type  string `json:"type"`
		}
	}

	var input reqBody
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON input",
		})
		return
	}
	result, err := uploadImageToCloudinary(input.Input.Base64Str, input.Input.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// fmt.Println(result)  

	c.JSON(http.StatusOK, gin.H{
		"image_url": result.SecureURL,
	})
}

func uploadImageToCloudinary(base64Data, imageName string) (*uploader.UploadResult, error) {
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	cloudinaryService, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	res, err := cloudinaryService.Upload.Upload(ctx, "data:image/jpeg;base64,"+base64Data, uploader.UploadParams{})
	if err != nil {
		fmt.Printf("Failed to upload file, %v\n", err)
		return nil, err
	}

	return res, nil
}