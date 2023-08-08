package actions

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go_backend/config"
)

type Vendor struct {
	Email      string `json:"email"`
	Full_name  string `string:"full_name"`
	Phone_no   string `json:"phone_no"`
	Status     bool   `json:"status"`
	Updated_at string `json:"updated_at"`
	Vendor_id  string `json:"vendor_id"`
	Created_at string `json:"created_at"`
}

func Test(c *gin.Context) {
	client := config.GraphqlClient()

	var data struct {
		Vendors []Vendor `json:"vendors"`
	}

	err := client.Query(context.Background(), &data, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, data)
}
