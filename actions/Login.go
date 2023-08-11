package actions

import (
	"encoding/json"
	"fmt"
	"go_backend/utilities"
	"net/http"
	"os"
	"time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/vladimiroff/jwt-go/v3"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var reqBody struct {
		Input struct {
			PhoneNo  string `json:"phone_no"`
			Password string `json:"password"`
		} `json:"input"`
	}

	if err := c.ShouldBind(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	input := reqBody.Input
	if input.PhoneNo == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide both Phone Number and Password"})
		return
	}

	type Role struct {
		RoleName string
		RoleID   string
	}


	type User struct {
		UserID   string
		PhoneNo  string
		Status   bool
		Password string
		Role     Role
	}

	var user User
	response, err := utilities.User(input.PhoneNo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}
	err = json.Unmarshal([]byte(response), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err != nil || user.Password == "" || user.Role.RoleName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Credentials"})
		return
	}
	isValidPassword := utilities.CompareHashAndPassword(user.Password, input.Password)
	if !isValidPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Credentials"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["iat"] = time.Now().Unix()

	hasuraClaims := map[string]interface{}{
		"x-hasura-allowed-roles":                          []string{"zadmin", "rider", "customer", "vendor"},
		"x-hasura-default-role":                           user.Role.RoleName,
		fmt.Sprintf("x-hasura-%s-id", user.Role.RoleName): user.UserID,
	}

	claims["https://hasura.io/jwt/claims"] = hasuraClaims
	tokenString, err := token.SignedString([]byte(os.Getenv("HASURA_GRAPHQL_JWT_SECRET")))
	
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": tokenString})
}
