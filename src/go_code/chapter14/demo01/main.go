package main

import (
	"fmt"
)

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (account *Account) Deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("Password is wrong")
		return
	}
	if money <= 0 {
		fmt.Println("money input is wrong")
		return
	}
	account.Balance += money
	fmt.Println("deposite money success~~~~")
}

func (account *Account) WithDraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("Password is wrong")
		return
	}
	if money <= 0 || money > account.Balance {
		fmt.Println("money input is wrong")
		return
	}
	account.Balance -= money
	fmt.Println("withdraw money success~~~~")
}

func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("Password is wrong")
		return
	}
	fmt.Printf("你的账号为%v, 余额为%v\n", account.AccountNo, account.Balance)
}

func main() {
	account := Account{
		AccountNo: "gs1111111",
		Pwd:       "666666",
		Balance:   100.0,
	}
	account.Query("666666")
	account.Deposite(200.0, "666666")
	account.Query("666666")
	account.WithDraw(100.0, "666666")
	account.Query("666666")
}
