package actions

import (
	"fmt"
	"go_backend/utilities" // Replace with the actual utilities package path
	"net/http"
	"github.com/gin-gonic/gin"
)

func CustomerSignup(c *gin.Context) {
	var input struct {
		FullName string
		Email    string
		PhoneNo  string
		Password string
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}


	if input.PhoneNo == "" || input.Password == "" || input.FullName == "" || input.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide all the details"})
		return
	}

	salt := utilities.GenerateSalt()

	hashedPassword, err := utilities.HashPassword(input.Password, salt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	customer, err := utilities.User(input.PhoneNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	if customer != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer Already Exists"})
		return
	}

	customerByEmail, err := utilities.FindCustomer(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	if customerByEmail != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Your Email is Already Registered"})
		return
	}

	customerID, err := utilities.InsertCustomer(input.Email, input.PhoneNo, input.FullName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	newCustomer, err := utilities.InsertUserPassword(customerID, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	fmt.Println("Congratulations :" , newCustomer)
	c.JSON(http.StatusOK, gin.H{"success": "Account Created Successfully"})
}


