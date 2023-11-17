package actions

import (
	"encoding/json"
	"fmt"
	"go_backend/utilities" // Replace with the actual utilities package path
	"net/http"
	"github.com/gin-gonic/gin"
)
func CustomerSignup(c *gin.Context) {
	var reqBody struct {
		Input struct {
			FName string `json:"fname"`
			LName string `json:"lname"`
			Username  string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
		} `graphql:"input"`
	}
}