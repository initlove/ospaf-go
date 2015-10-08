package ospafLib

import (
	"fmt"
)

type Pool struct {
	Accounts []Account
}

func InitPool() (pool Pool, err error) {
	pool.Accounts, err = LoadAccounts("")
	if err != nil {
		fmt.Println("Cannot Using pool due to: ", err)
		return pool, err
	}

	for index := 0; index < len(pool.Accounts); index++ {
		pool.Accounts[index].Load()
	}
	return pool, err
}

//Could have different algorithm
func (pool *Pool) PickAccount() int {
	maxIndex := -1
	maxRemain := 0

	for index := 0; index < len(pool.Accounts); index++ {
		if pool.Accounts[index].Remains > maxRemain {
			maxRemain = pool.Accounts[index].Remains
			maxIndex = index
		}
	}
	if maxRemain < 5 {
		fmt.Println("The pool is unhealthy!")
		maxIndex = -1
	}
	return maxIndex
}

func (pool *Pool) ReadURL(url string, param string) (string, int) {
	index := pool.PickAccount()
	if index == -1 {
		return "No avaiable account in the pool", -1
	} else {
		//	fmt.Println("Using ", pool.Accounts[index].User)
	}
	return pool.Accounts[index].ReadURL(url, param)
}
