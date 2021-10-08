package main

import (
	"fmt"

	"github.com/kuku1198/nomadcoders-learngo/accounts"
)

func main() {
	account := accounts.NewAccount("YeongGyu")

	fmt.Println(account)
}
