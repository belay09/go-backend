package actions

import (
	"encoding/json"
	"fmt"
	"go_backend/utilities" // Replace with the actual utilities package path
	"net/http"
	"reflect"
	"github.com/gin-gonic/gin"
)

func RiderSignup(c *gin.Context) {
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
	salt := utilities.GenerateSalt()

	hashedPassword, err := utilities.HashPassword(reqBody.Input.Password, salt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	riderResponse, err := utilities.User(reqBody.Input.PhoneNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	type Role struct {
		RoleName string
		RoleID   string
	}
	type Rider struct {
		UserID   string
		PhoneNo  string
		Status   bool
		Password string
		Role     Role
	}

	var response Rider
	// 	if riderResponse == ""{
	// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Account Already Exists by this Phone Number"})
	// 		return
	// 	}
	// 	fmt.Println("from below")

	//    //unmarshal it conditionally if riderResponse exists then unmarshal it otherwise continue
	//   // write a code here
	// 	err = json.Unmarshal([]byte(riderResponse), &response)

	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	// 		return
	// 	}
	// 	fmt.Println("from Here", response)
	// 	if response.UserID != "" {
	// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Rider Already Exists"})
	// 		return
	// 	}

	if riderResponse != "" {
		fmt.Println("hhhhhhhhhhhhhhhhhhhhhhhh")
		err = json.Unmarshal([]byte(riderResponse), &response)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		fmt.Println("from Here", response)

		if response.UserID != "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Rider Already Exists"})
			return
		}
	}

	fmt.Println("from above")

	riderByEmail, err := utilities.FindRider(reqBody.Input.Email)

	fmt.Println("AM great", riderByEmail, err)
	if err != nil {
		fmt.Println(riderByEmail)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if riderByEmail != "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Your Email is Already Registered"})
		return
	}

	riderID, err := utilities.InsertRider(reqBody.Input.Email, reqBody.Input.PhoneNo, reqBody.Input.FullName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	user_id := riderID
	fmt.Println(reflect.TypeOf(riderID))  
	newRider, err := utilities.InsertUserPassword(user_id, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	fmt.Println("Congratulations :", newRider)
	c.JSON(http.StatusOK, gin.H{"success": "Account Created Successfully"})
}
