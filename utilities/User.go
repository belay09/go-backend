package utilities

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go_backend/config"
	// "reflect"
)

type Authentication struct {
	Username  string `graphql:"username"`
	UserID  int `graphql:"id"`
	Password string `graphql:"password"`
}
type Response struct {
	Authentications []Authentication `graphql:"users(limit: 1, where: {email: {_eq: $email}})"`
}

func User(email string) (string, error) {
	fmt.Println(email)
	client := config.GraphqlClient()
	variables := map[string]interface{}{
		"email": email,
	}
	var response Response
	err := client.Query(context.Background(), &response, variables)
	fmt.Println("hihihii",len(response.Authentications))

	if err != nil {
		fmt.Println("An error occurred:", err)
		return "", errors.New("error fetching user data")
	}
	if len(response.Authentications) > 0 {
		firstUser := response.Authentications[0]
		fmt.Println(firstUser)

		responseJSON, _ := json.Marshal(firstUser)
		fmt.Println(string(responseJSON))
		return string(responseJSON), nil
	} 

	if len(response.Authentications)== 0 {
		hi:=""
		return string(hi), nil
	}else {
		return "", nil
	}
	
}

type mutation struct {
	UpdateUsers struct {
		AffectedRows int `graphql:"affected_rows"`
	} `graphql:"update_users(where: {username: {_eq: $username}}, _set: {password: $password})"`
}
func InsertUserPassword(user_name, password string) (string, error) {
	client := config.GraphqlClient()
	type uuid string
	variables := map[string]interface{}{
		"username":  user_name,
		"password": password,
	}

	fmt.Println(user_name)
	fmt.Println("Type Of String", variables["password"])
	var response mutation

	err := client.Mutate(context.Background(), &response, variables)
	fmt.Println("hjikjij", response)
	if err != nil {
		fmt.Println("error occurred:", err)
		return "", errors.New("error fetching user data")
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println(responseJSON)
	return "", nil
}
