package ospafLib

import (
	github "../github"
	//	"fmt"
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
	account.Remains = -1

	url := "https://api.github.com/rate_limit"
	val, code := account.ReadURL(url, "")
	if code == 200 {
		rl, ok := github.RateLimitFrom(val)
		if ok {
			account.Remains = rl.Resources.Core.Remaining
		}
	}
}

func (account *Account) GetRemains() int {
	return account.Remains
}

func (account *Account) ReadURL(url string, param string) (string, int) {
	if account.Remains != -1 && account.Remains < 10 {
		return "System warning: not enough remain access", -1
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
		return "System warning: cannot send get request", -1
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "System warning: cannot read the body", -1
	}

	return string(resp_body), resp.StatusCode
}
