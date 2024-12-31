package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func (t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.GetBalance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("deposit", func (t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(10))

		assertError(t, err)
		assertBalance(t, wallet, Bitcoin(0))
	})

	t.Run("withdraw insufficient funds", func (t *testing.T) {
		startingBalance := Bitcoin(30)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(300))

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}