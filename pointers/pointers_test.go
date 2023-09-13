package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("Balance test, got %s but want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		if got == nil {
			t.Error("wanted an error but didn't get one")
		}

		if got.Error() != want.Error() {
			t.Errorf("got %v but wanted %v", got, want)
		}
	}
	t.Run("Show Balance correctly", func(t *testing.T) {

		wallet := Wallet{}

		//Make a deposit
		wallet.Deposit(10)
		//Show balance
		got := wallet.Balance()
		want := Bitcoin(10)
		fmt.Printf("\n address of balance in test is %p \n", &wallet)

		if got != want {
			t.Errorf("Balance test, got %s but want %s", got, want)
		}
	})

	t.Run("withdraw with sufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		want := Bitcoin(20)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, want)
	})

}
