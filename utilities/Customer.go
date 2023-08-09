package utilities

import (
	"context"
	"fmt"
	"go_backend/config"
)

type InsertCustomersOneResponse struct {
	InsertCustomersOne struct {
		CustomerID string `json:"customer_id"`
	} `graphql:"insert_customers_one(object: {email: $email, full_name: $full_name, phone_no: $phone_no})"`
}

func InsertCustomer(email, phoneNo, fullName string) (string, error) {
	client := config.GraphqlClient()
	var variables = map[string]interface{}{
		"email":     email,
		"phone_no":  phoneNo,
		"full_name": fullName,
	}
	var response InsertCustomersOneResponse

	err := client.Mutate(context.Background(), &response, variables)
	if err != nil {
		fmt.Println("An error occurred:", err)
		return "", err
	}

	customerID := response.InsertCustomersOne.CustomerID
	return customerID, nil
}

type CustomerByEmail struct {
	Customer struct {
		CustomerID string `json:"customer_id"`
	} `graphql:"customers(where: {email: {_eq: $email}})"`
}

func FindCustomer(email string) (string, error) {
	client := config.GraphqlClient()
	variables := map[string]interface{}{
		"email": email,
	}
	var response CustomerByEmail

	err := client.Query(context.Background(), &response, variables)
	if err != nil {
		fmt.Println("An error occurred:", err)
		return "", err
	}
	customerId := response.Customer.CustomerID
	return customerId, nil
}
