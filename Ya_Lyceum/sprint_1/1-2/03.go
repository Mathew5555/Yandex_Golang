package main

import (
	"errors"
	"fmt"
)

var (
	NegativeBalance = errors.New("Negative balance on the account")
)

type Account struct {
	balance float64
	owner   string
}

func NewAccount(ownerAcc string) *Account {
	return &Account{balance: 0, owner: ownerAcc}
}

func (acc *Account) SetBalance(money float64) error {
	if money < 0 {
		return NegativeBalance
	}
	acc.balance = money
	return nil
}

func (acc *Account) Deposit(money float64) error {
	if money < 0 {
		return NegativeBalance
	}
	acc.balance += money
	return nil
}

func (acc *Account) Withdraw(money float64) error {
	if acc.balance-money < 0 || money < 0 {
		return NegativeBalance
	}
	acc.balance -= money
	return nil
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}
func main() {
	account := NewAccount("Alice")

	// Устанавливаем баланс
	err := account.SetBalance(1000.0)
	if err != nil {
		_ = fmt.Errorf("Ошибка при установке баланса: %v", err)
	}

	// Вносим деньги
	err = account.Deposit(500.0)
	if err != nil {
		_ = fmt.Errorf("Ошибка при внесении денег: %v", err)
	}

	// Снимаем деньги
	err = account.Withdraw(200.0)
	if err != nil {
		_ = fmt.Errorf("Ошибка при снятии денег: %v", err)
	}

	// Получаем текущий баланс
	balance := account.GetBalance()
	expectedBalance := 1300.0
	if balance != expectedBalance {
		_ = fmt.Errorf("Ожидается баланс %.2f, получено %.2f", expectedBalance, balance)
	}
}
