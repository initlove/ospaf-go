package ospafLib

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Account struct {
	//Basic/Oauth
	Type     string
	User     string
	Password string

	Remains int
}

//TODO: add lock to Remains

func (account *Account) Init(accountType string, accountUser string, accountPassword string) {
	account.Type = accountType
	account.User = accountUser
	account.Password = accountPassword
	account.Remains = 5000
	//	account.Remains = -1

	url := "https://api.github.com/rate_limit"
	_, val := account.ReadURL(url, "")

	fmt.Println(val)
}

func (account *Account) GetRemains() int {
	return account.Remains
}

func (account *Account) ReadURL(url string, param string) (int, string) {
	if account.Remains != -1 && account.Remains < 10 {
		return -1, "System warning: not enough remain access"
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	switch account.Type {
	case "Basic":
		req.SetBasicAuth(account.User, account.Password)
		break
	}
	resp, err := client.Do(req)
	account.Remains -= 1
	if err != nil {
		return -1, "System warning: cannot send get request"
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, "System warning: cannot read the body"
	}

	return resp.StatusCode, string(resp_body)
}
