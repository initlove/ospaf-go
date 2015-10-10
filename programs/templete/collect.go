package main

import (
	github "../../github"
	ospaf "../../lib"
)

func main() {
	_, err := ospaf.InitPool()
	if err != nil {
		fmt.Println(err)
		return
	}

}
