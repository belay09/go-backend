package utilities

import (
	"context"
	"fmt"
	"go_backend/config"
)

type InsertVendorsOneResponse struct {
	InsertVendorsOne struct {
		VendorId string `graphql:"vendor_id"`
	} `graphql:"insert_vendors_one(object: {email: $email, full_name: $full_name, phone_no: $phone_no})"`
}

func InsertVendor(email, phoneNo, fullName string) (string, error) {
	client := config.GraphqlClient()
	var variables = map[string]interface{}{
		"email":     email,
		"phone_no":  phoneNo,
		"full_name": fullName,
	}
	var response InsertVendorsOneResponse

	err := client.Mutate(context.Background(), &response, variables)
	if err != nil {

		fmt.Println("An error occurred:", err)
		return "", err
	}
	vendorId := response.InsertVendorsOne.VendorId
	fmt.Println("Registration Done",vendorId)
	return vendorId, nil
}
type vendorByEmailResponse struct {
	Vendors []struct {
		vendorId string `graphql:"vendor_id"`
	} `graphql:"vendors(where: {email: {_eq: $email}})"`
}

func FindVendor(email string) (string, error) {
	fmt.Println("am an Email", email)
	client := config.GraphqlClient()
	variables := map[string]interface{}{
		"email": email,
	}
	var response vendorByEmailResponse

	err := client.Query(context.Background(), &response, variables)
	if err != nil {
		fmt.Println("An error occurred here:", err)
		return "", err
	}
	if len(response.Vendors) > 0 {
		vendorId := response.Vendors[0].vendorId
		return vendorId, nil
	}
	return "", nil
}

