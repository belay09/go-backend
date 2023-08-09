package actions

import (
	"fmt"
	"go_backend/utilities" // Replace with the actual utilities package path
	"net/http"
	"github.com/gin-gonic/gin"
)

func VendorSignup(c *gin.Context) {
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

	vendor, err := utilities.User(input.PhoneNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	if vendor != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account Already Exists by this Phone Number"})
		return
	}

	vendorByEmail, err := utilities.FindVendor(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	if vendorByEmail != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Your Email is Already Registered"})
		return
	}

	vendorID, err := utilities.InsertVendor( input.Email, input.PhoneNo, input.FullName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	newVendor, err := utilities.InsertUserPassword(vendorID, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	fmt.Println("Congratulations :" , newVendor)
	c.JSON(http.StatusOK, gin.H{"success": "Account Created Successfully"})
}


