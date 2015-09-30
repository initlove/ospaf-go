package main

import (
	ospaf "../"
	"encoding/json"
	"fmt"
)

func main() {
	var account ospaf.Account
	account.Init("Basic", "golang001", "qwe123456")

	var accounts []ospaf.Account
	accounts = append(accounts, account)
	accounts = append(accounts, account)
	accounts = append(accounts, account)
	val, _ := json.MarshalIndent(accounts, "", "\t")
	fmt.Println(string(val))
	return
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
