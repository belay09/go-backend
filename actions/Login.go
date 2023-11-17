package actions

import (
	"encoding/json"
	"fmt"
	"go_backend/utilities"
	"net/http"
	"os"
	"time"
	"strconv"
	// "github.com/dgrijalva/jwt-go"
	"github.com/vladimiroff/jwt-go/v3"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var reqBody struct {
		Input struct {
			Email  string `json:"email"`
			Password string `json:"password"`
		} `json:"input"`
	}

	if err := c.ShouldBind(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	input := reqBody.Input
	if input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide both Phone Number and Password"})
		return
	}



	type User struct {
		Username   string
		UserID  int
		Password string

	}

	var user User
	response, err := utilities.User(input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}
	fmt.Println("hheloo",response)
	if response =="hi" {
		c.JSON(http.StatusBadRequest, gin.H{"message":"no accaunt associated with this one"})

	}



if response !="hi" {
	

	err = json.Unmarshal([]byte(response), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	fmt.Println("thiere")

	if err != nil || user.Password == "" {
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
	userIDString := strconv.Itoa(user.UserID)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["iat"] = time.Now().Unix()

	hasuraClaims := map[string]interface{}{
		"x-hasura-allowed-roles":                          []string{"user","admin"},
		"x-hasura-default-role":                           "user",
		"x-hasura-user-id": userIDString,
		}

	claims["https://hasura.io/jwt/claims"] = hasuraClaims
	tokenString, err := token.SignedString([]byte(os.Getenv("HASURA_GRAPHQL_JWT_SECRET")))
	
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": tokenString})
}
}
