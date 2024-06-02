package main

import (
	"fmt"
	"time"

	"github.com/leommuniz/wallet/entity"
)

func main() {
	user, _ := entity.NewUser("leonardo", time.Now(), "leonardo@leonardo.com", "1234")
	wallet, _ := entity.NewWallet(user)
	wallet.Deposit(100.0)
	fmt.Println(wallet.Balance)
}
