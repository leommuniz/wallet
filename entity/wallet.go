package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/leommuniz/wallet/pkg/entity"
)

type Wallet struct {
	ID       uuid.UUID `json:"id"`
	CreateIn time.Time `json:"create_in"`
	Balance  float64   `json:"balance"`
	User     *User     `json:"user"`
}

func NewWallet(user *User) (*Wallet, error) {
	wallet := &Wallet{
		ID:       uuid.New(),
		CreateIn: time.Now(),
		User:     user,
		Balance:  0.0,
	}
	return wallet, nil
}

func (wallet *Wallet) Deposit(value float64) (float64, error) {
	if value < 0 {
		return 0, entity.NewCustomErro("O valor deve ser positivo")
	}
	wallet.Balance += value
	NewTransiction(value, false, "E", "Entrada de dinheiro na carteira")
	return wallet.Balance, nil
}

func (wallet *Wallet) Withdraw(value float64) (float64, error) {
	if value < 0 {
		return 0, entity.NewCustomErro("O valor deve ser positivo")
	}
	if wallet.Balance < value {
		return 0, entity.NewCustomErro("O valor a sacar deve ser maior ou igual ao valor disponivel na carteira")
	}
	wallet.Balance -= value
	NewTransiction(value, false, "S", "Saida de dinheiro da carteria")
	return wallet.Balance, nil
}
