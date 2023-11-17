package router
import (
	"github.com/gin-gonic/gin"
	"go_backend/actions"
)
func InitializeRoutes(router *gin.RouterGroup) {

	router.POST("/login", actions.Login)
	router.POST("/customer_signup", actions.CustomerSignup)
	router.POST("/file_upload", actions.FileUpload)
}