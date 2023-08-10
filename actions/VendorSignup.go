package actions

import (
	"encoding/json"
	"fmt"
	"go_backend/utilities" // Replace with the actual utilities package path
	"net/http"
	"reflect"
	"github.com/gin-gonic/gin"
)

func VendorSignup(c *gin.Context) {
	var reqBody struct {
		Input struct {
			FullName string `json:"full_name"`
			Email    string `json:"email"`
			PhoneNo  string `json:"phone_no"`
			Password string `json:"password"`
		} `graphql:"input"`
	}

	if err := c.ShouldBind(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	fmt.Println("Congratulations", reqBody.Input)

	if reqBody.Input.PhoneNo == "" || reqBody.Input.Password == "" || reqBody.Input.FullName == "" || reqBody.Input.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide all the data"})
		return
	}

	hashedPassword, err := utilities.HashPassword(reqBody.Input.Password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	vendorResponse, err := utilities.User(reqBody.Input.PhoneNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	type Role struct {
		RoleName string
		RoleID   string
	}
	type Vendor struct {
		UserID   string
		PhoneNo  string
		Status   bool
		Password string
		Role     Role
	}

	var response Vendor

	if vendorResponse != "" {
		fmt.Println("hhhhhhhhhhhhhhhhhhhhhhhh")
		err = json.Unmarshal([]byte(vendorResponse), &response)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		fmt.Println("from Here", response)

		if response.UserID != "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "vendor Already Exists"})
			return
		}
	}

	fmt.Println("from above")

	vendorByEmail, err := utilities.FindVendor(reqBody.Input.Email)

	fmt.Println("AM great", vendorByEmail, err)
	if err != nil {
		fmt.Println(vendorByEmail)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if vendorByEmail != "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Your Email is Already Registered"})
		return
	}

	vendorID, err := utilities.InsertVendor(reqBody.Input.Email, reqBody.Input.PhoneNo, reqBody.Input.FullName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	user_id := vendorID
	fmt.Println(reflect.TypeOf(vendorID))  
	newvendor, err := utilities.InsertUserPassword(user_id, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	fmt.Println("Congratulations :", newvendor)
	c.JSON(http.StatusOK, gin.H{"success": "Account Created Successfully"})
}
