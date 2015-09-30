package main

import (
	ospaf "../../lib"
	"fmt"
)

func main() {
	pool, err := ospaf.InitPool()
	if err != nil {
		fmt.Println(err)
		return
	}

	chars := "abcdefghijklmnopqrstuvwxyz"

	for a_index := 0; a_index < len(chars); a_index++ {
		url := fmt.Sprintf("https://api.github.com/users/ab%c", chars[a_index])
		_, statusCode := pool.ReadURL(url, "")
		if statusCode == -1 {
			return
		}
		fmt.Println("ab", chars[a_index], statusCode)
	}

}
