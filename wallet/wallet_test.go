package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit to wallet", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		assertBallance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw from wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(22)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBallance(t, wallet, 17)
	})

	t.Run("insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBallance(t, wallet, startingBalance)

	})

}

func assertBallance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("\ngot %s\nwant %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("hot an error but didn't want one")
	}

}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
