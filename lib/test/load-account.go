package main

import (
	ospaf "../"
	"fmt"
)

func AccountTest(typeA string, userA string, passwordA string) {
	var account ospaf.Account
	account.Init(typeA, userA, passwordA)
	fmt.Println("Account remaining: ", account.GetRemains())
	test_user := "initlove"
	url := fmt.Sprintf("https://api.github.com/users/%s", test_user)
	info, code := account.ReadURL(url, "")
	if code != 200 {
		fmt.Println(test_user, info)
	} else {
		fmt.Println(info)
	}
}
func main() {
	accounts, err := ospaf.LoadAccounts("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(accounts)
		for index := 0; index < len(accounts); index++ {
			fmt.Println("\nType: ", accounts[index].Type, " Name: ", accounts[index].User)
			AccountTest(accounts[index].Type, accounts[index].User, accounts[index].Password)
		}
	}
}
