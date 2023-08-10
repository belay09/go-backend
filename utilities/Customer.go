package utilities

import (
	"context"
	"fmt"
	"go_backend/config"
)

type InsertCustomersOneResponse struct {
	InsertCustomersOne struct {
		CustomerId string `graphql:"customer_id"`
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
	CustomerId := response.InsertCustomersOne.CustomerId
	fmt.Println("Registration Done",CustomerId)
	return CustomerId, nil
}
type CustomerByEmailResponse struct {
	Customers []struct {
		CustomerId string `graphql:"customer_id"`
	} `graphql:"customers(where: {email: {_eq: $email}})"`
}

func FindCustomer(email string) (string, error) {
	fmt.Println("am an Email", email)
	client := config.GraphqlClient()
	variables := map[string]interface{}{
		"email": email,
	}
	var response CustomerByEmailResponse

	err := client.Query(context.Background(), &response, variables)
	if err != nil {
		fmt.Println("An error occurred here:", err)
		return "", err
	}
	if len(response.Customers) > 0 {
		CustomerId := response.Customers[0].CustomerId
		return CustomerId, nil
	}
	return "", nil
}

