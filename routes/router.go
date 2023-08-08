package router

import (
	"github.com/gin-gonic/gin"
	"go_backend/actions"
	"net/http"
)

func InitializeRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
	})
	router.POST("/login", actions.Login)
	router.GET("/test", actions.Test)
	router.POST("/customer_signup", actions.CustomerSignup)
	router.POST("/rider_signup", actions.RiderSignup)
	router.POST("/vendor_signup", actions.VendorSignup)
	router.POST("/file_upload", actions.FileUpload)
}