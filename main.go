package main

import (
	"fmt"

	"github.com/kuku1198/nomadcoders-learngo/accounts"
)

func main() {
	account := accounts.NewAccount("YeongGyu")
	account.Deposit(10)
	fmt.Println(account.Balance())

	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
