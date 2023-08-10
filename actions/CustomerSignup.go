package actions

import (
	"encoding/json"
	"fmt"
	"go_backend/utilities" // Replace with the actual utilities package path
	"net/http"
	"reflect"
	"github.com/gin-gonic/gin"
)

// CustomerSignup handles the customer signup process.
//
// It accepts a gin.Context object as the parameter and retrieves the request
// body from it. The request body is expected to be in JSON format and should
// contain the following fields:
// - full_name (string): The full name of the customer.
// - email (string): The email address of the customer.
// - phone_no (string): The phone number of the customer.
// - password (string): The password for the customer's account.
//
// If the request payload is invalid, it returns a JSON response with a
// status code of 400 and a message indicating the invalid payload.
//
// It then proceeds to validate the input fields and ensures that all the
// required data is provided. If any of the fields is missing, it returns a
// JSON response with a status code of 400 and a message indicating that all
// the data should be provided.
//
// Next, it generates a salt for password hashing using the
// utilities.GenerateSalt() function.
//
// It then hashes the provided password using the generated salt using the
// utilities.HashPassword() function. If any error occurs during the hashing
// process, it returns an error message.
//
// The function then calls the utilities.User() function with the phone number
// from the request body to retrieve customer information. If an error occurs
// during this process, it returns a JSON response with a status code of 500
// and a generic error message.
//
// It then defines two structs: Role and Customer, which represent the role and
// customer information, respectively.
//
// If the customer information is not empty, it attempts to unmarshal the
// customer response into the response struct. If an error occurs during the
// unmarshaling process, it returns a JSON response with a status code of 400
// and an error message.
//
// It then checks if the user ID in the response struct is empty. If it is not
// empty, it means that the customer already exists and it returns a JSON
// response with a status code of 400 and a message indicating that the customer
// already exists.
//
// Next, it calls the utilities.FindCustomer() function with the email from the
// request body to check if there is already a customer with the provided email
// address. If an error occurs during this process, it returns a JSON response
// with a status code of 400 and an error message.
// 
// If a customer with the provided email address is found, it returns a JSON
// response with a status code of 400 and a message indicating that the email
// address is already registered.
//
// If no customer with the provided email address is found, it calls the
// utilities.InsertCustomer() function with the email, phone number, and full
// name from the request body to insert a new customer into the database. If an
// error occurs during this process, it returns a JSON response with a status
// code of 400 and an error message.
//
// It then retrieves the customer ID from the result of the previous query and
// assigns it to the user_id variable.
//
// Next, it calls the utilities.InsertUserPassword() function with the user ID
// and the hashed password to insert the user's password into the database. If
// an error occurs during this process, it returns a JSON response with a status
// code of 400 and an error message.
//
// Finally, it returns a JSON response with a status code of 200 and a success
// message indicating that the account was created successfully.
func CustomerSignup(c *gin.Context) {
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

	customerResponse, err := utilities.User(reqBody.Input.PhoneNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	type Role struct {
		RoleName string
		RoleID   string
	}
	type Customer struct {
		UserID   string
		PhoneNo  string
		Status   bool
		Password string
		Role     Role
	}

	var response Customer
	if customerResponse != "" {
		fmt.Println("hhhhhhhhhhhhhhhhhhhhhhhh")
		err = json.Unmarshal([]byte(customerResponse), &response)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		fmt.Println("from Here", response)

		if response.UserID != "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Customer Already Exists"})
			return
		}
	}

	fmt.Println("from above")

	CustomerByEmail, err := utilities.FindCustomer(reqBody.Input.Email)

	fmt.Println("AM great", CustomerByEmail, err)
	if err != nil {
		fmt.Println(CustomerByEmail)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if CustomerByEmail != "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Your Email is Already Registered"})
		return
	}

	CustomerID, err := utilities.InsertCustomer(reqBody.Input.Email, reqBody.Input.PhoneNo, reqBody.Input.FullName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	user_id := CustomerID
	fmt.Println(reflect.TypeOf(CustomerID))  
	fmt.Println("Hashed Password not  Stringified",hashedPassword)
	fmt.Println("Hashed Password  Stringified",string(hashedPassword))
	newCustomer, err := utilities.InsertUserPassword(user_id, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	fmt.Println("Congratulations :", newCustomer)
	c.JSON(http.StatusOK, gin.H{"success": "Account Created Successfully"})
}
