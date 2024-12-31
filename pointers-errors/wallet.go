package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(bitcoinAmount Bitcoin) {
	w.balance += bitcoinAmount
	return
}

func (w *Wallet) Withdraw(bitcoinAmount Bitcoin) error {
	if bitcoinAmount > w.balance {
		err := errors.New("withdraw amount is greater than")
		return err
	}

	w.balance -= bitcoinAmount
	return nil
}

func (w *Wallet) GetBalance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
