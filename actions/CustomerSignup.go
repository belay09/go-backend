package actions

import (
	"encoding/json"
	"fmt"
	"go_backend/utilities" // Replace with the actual utilities package path
	"net/http"
	// "reflect"
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

	if err := c.ShouldBind(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	fmt.Println("Congratulations", reqBody.Input)

	if reqBody.Input.FName == "" || reqBody.Input.LName == "" || reqBody.Input.Username == "" || reqBody.Input.Password == "" || reqBody.Input.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide all the data"})
		return
	}

	hashedPassword, err := utilities.HashPassword(reqBody.Input.Password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	customerResponse, err := utilities.User(reqBody.Input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}
	fmt.Println("Hashed Password  Stringified",hashedPassword)

	type Customer struct {
		username   string
	}

	var response Customer
	if customerResponse != "" {
		fmt.Println("hhhhhhhhhhhhhhhhhhhhhhhh")
		fmt.Println(customerResponse)

		err = json.Unmarshal([]byte(customerResponse), &response)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		fmt.Println("from Here", response)

		if response.username != "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Customer Already Exists"})
			return
		}
	}

	fmt.Println("from above")

	Customername, err := utilities.InsertCustomer(reqBody.Input.Email, reqBody.Input.FName, reqBody.Input.LName, reqBody.Input.Password, reqBody.Input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	// user_id := CustomerID
	fmt.Println(Customername)  
	user_name := Customername
	fmt.Println("Hashed Password not  Stringified",hashedPassword)
	fmt.Println("Hashed Password  Stringified",string(hashedPassword))
	newCustomer, err := utilities.InsertUserPassword(user_name, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	fmt.Println("Congratulations :", newCustomer)
	c.JSON(http.StatusOK, gin.H{"success": "Account Created Successfully"})
}
