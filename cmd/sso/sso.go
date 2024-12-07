package main

import (
	"fmt"
	"sso/pkg/contracts/ssocontract"
)

func main() {
	fmt.Println(ssocontract.AuthenticateRequest{
		Username: "username",
		Password: "password",
	})
}
