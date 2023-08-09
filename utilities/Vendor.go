package utilities
import (
	"context"
	"fmt"
	"go_backend/config"
)
type InsertVendorsOneResponse struct {
	InsertVendorsOne struct {
		VendorId string `json:"Vendor_id"`
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
	VendorId := response.InsertVendorsOne.VendorId
	return VendorId, nil
}

type VendorByEmail struct {
	Vendor struct {
		VendorId string `json:"vendor_id"`
	} `graphql:"vendors(where: {email: {_eq: $email}})"`
}

func FindVendor(email string) (string, error) {
	client := config.GraphqlClient()
	variables := map[string]interface{}{
		"email": email,
	}
	var response VendorByEmail

	err := client.Query(context.Background(), &response, variables)
	if err != nil {
		fmt.Println("An error occurred:", err)
		return "", err
	}
	VendorId := response.Vendor.VendorId
	return VendorId, nil
}
