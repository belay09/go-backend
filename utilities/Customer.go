package utilities

import (
	"context"
	"fmt"
	"go_backend/config"
)

type InsertCustomersOneResponse struct {
	InsertCustomersOne struct {
		CustomerId string `graphql:"username"`
	} `graphql:"insert_users_one(object: {fname: $fname, lname: $lname, password: $password, username: $username, email: $email})"`
}

func InsertCustomer(email, fname,lname, password,username string) (string, error) {
	client := config.GraphqlClient()
	var variables = map[string]interface{}{
		"email":     email,
		"fname":  fname,
		"lname": lname,
		"password": password,
		"username": username,
	}
	var response InsertCustomersOneResponse

	err := client.Mutate(context.Background(), &response, variables)
	if err != nil {
		fmt.Println("An sehetet occurred:", err)
		return "", err
	}
	CustomerId := response.InsertCustomersOne.CustomerId
	fmt.Println("Registration Done",CustomerId)
	return CustomerId, nil
}
