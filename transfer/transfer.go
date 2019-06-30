package transfer

import (
	"bankSim/account"
	"errors"
	"fmt"
)

type TransferRequest struct {
	FromAccountId int
	ToAccountId int
	Value int
}

type Transfer struct {
	FromAccount account.Account
	ToAccount account.Account
	Value int
}

func transferIsAllowed(t Transfer) (bool, error) {
	fmt.Println(t.FromAccount)
	if t.FromAccount.Blocked && t.ToAccount.Blocked {
		return false, errors.New("one of the accounts is blocked from transactions")
	}
	if !t.FromAccount.HasAvailableFunds(t.Value) {
		return false, fmt.Errorf("account with id: %d has insuffcient funds", t.FromAccount.ID)
	}

	return true, nil
}

func (t *Transfer) Move() (bool, error){
	if _, err := transferIsAllowed(*t); err != nil {
		return false, err
	} else {
		fmt.Println(t.FromAccount)
		t.FromAccount.Withdraw(t.Value)
		t.ToAccount.Deposit(t.Value)
		return true, nil
	}
}
