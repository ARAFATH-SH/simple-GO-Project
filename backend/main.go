package main

import (
	"ecommerce/util"
	"fmt"
)

func main() {
	// cmd.Serve()
	jwt, err := util.CreateJwt("my-secret", util.Payload{
		Sub:         45,
		FirstName:   "Arafath",
		LastName:    "Rahman",
		Email:       "arafath@gmail.com",
		IsShopOwner: false,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jwt)
}
