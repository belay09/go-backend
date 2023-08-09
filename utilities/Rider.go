package utilities

import (
	"context"
	"fmt"
	"go_backend/config"
)

type InsertRidersOneResponse struct {
	InsertRidersOne struct {
		RiderId string `graphql:"rider_id"`
	} `graphql:"insert_riders_one(object: {email: $email, full_name: $full_name, phone_no: $phone_no})"`
}

func InsertRider(email, phoneNo, fullName string) (string, error) {
	fmt.Println("wowoownjkbvkjdbs")
	client := config.GraphqlClient()
	var variables = map[string]interface{}{
		"email":     email,
		"phone_no":  phoneNo,
		"full_name": fullName,
	}
	var response InsertRidersOneResponse

	err := client.Mutate(context.Background(), &response, variables)
	if err != nil {

		fmt.Println("An error occurred:", err)
		return "", err
	}

	RiderId := response.InsertRidersOne.RiderId
	fmt.Println("Registration Done",RiderId)
	return RiderId, nil
}

type RiderByEmailResponse struct {
	Riders []struct {
		RiderId string `graphql:"rider_id"`
	} `graphql:"riders(where: {email: {_eq: $email}})"`
}

func FindRider(email string) (string, error) {
	fmt.Println("am an Email", email)
	client := config.GraphqlClient()
	variables := map[string]interface{}{
		"email": email,
	}
	var response RiderByEmailResponse

	err := client.Query(context.Background(), &response, variables)
	if err != nil {
		fmt.Println("An error occurred here:", err)
		return "", err
	}

	if len(response.Riders) > 0 {
		riderId := response.Riders[0].RiderId
		return riderId, nil
	}

	return "", nil
}

