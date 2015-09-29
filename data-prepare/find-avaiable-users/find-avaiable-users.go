package main

import (
	ospaf "../../lib"
	"fmt"
)

func main() {
	var account ospaf.Account
	account.Init("Basic", "initlove", "")

	chars := "abcdefghijklmnopqrstuvwxyz0123456789"

	for a_index := 0; a_index < len(chars); a_index++ {
		for b_index := 0; b_index < len(chars); b_index++ {
			for c_index := 0; c_index < len(chars); c_index++ {
				for d_index := 0; d_index < len(chars); d_index++ {
					test_user := fmt.Sprintf("%c%c%c%c", chars[a_index], chars[b_index], chars[c_index], chars[d_index])
					url := fmt.Sprintf("https://api.github.com/users/%s", test_user)
					ok, _ := account.ReadURL(url, "")
					if ok == false {
						fmt.Println(test_user)
					}
				}
			}
		}
	}

}
