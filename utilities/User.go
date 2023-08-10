package utilities

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go_backend/config"
	"reflect"
)

type Role struct {
	RoleName string `graphql:"role_name"`
	RoleID   string `graphql:"role_id"`
}

type Authentication struct {
	PhoneNo  string `graphql:"phone_no"`
	UserID   string `graphql:"user_id"`
	Status   bool   `graphql:"status"`
	Role     Role   `graphql:"role"`
	Password string `graphql:"password"`
}
type Response struct {
	Authentications []Authentication `graphql:"authentications(where: {phone_no: {_eq: $phone_no}})"`
}

func User(phone_no string) (string, error) {
	fmt.Println(phone_no)
	client := config.GraphqlClient()
	variables := map[string]interface{}{
		"phone_no": phone_no,
	}
	var response Response
	err := client.Query(context.Background(), &response, variables)
	if err != nil {
		fmt.Println("An error occurred:", err)
		return "", errors.New("error fetching user data")
	}
	if len(response.Authentications) > 0 {
		firstUser := response.Authentications[0]
		responseJSON, _ := json.Marshal(firstUser)
		fmt.Println(string(responseJSON))
		return string(responseJSON), nil
	} else {
		return "", nil
	}
}

type mutation struct {
	UpdateUserPassword struct {
		UserID string `graphql:"user_id"`
	} `graphql:"update_authentications_by_pk(pk_columns: {user_id: $user_id}, _set: {password: $password})"`
}

func InsertUserPassword(user_id, password string) (string, error) {
	client := config.GraphqlClient()
	type uuid string
	variables := map[string]interface{}{
		"user_id":  uuid(user_id),
		"password": password,
	}

	fmt.Println("Type Of UUID", reflect.TypeOf(variables["user_id"]))
	fmt.Println("Type Of String", reflect.TypeOf(variables["password"]))
	var response mutation

	err := client.Mutate(context.Background(), &response, variables)
	fmt.Println("", response)
	if err != nil {
		fmt.Println("An error occurred:", err)
		return "", errors.New("error fetching user data")
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println(string(responseJSON))
	return "", nil
}
