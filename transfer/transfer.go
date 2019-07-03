package transfer

import (
	"banksim/account"
	"banksim/transaction"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type TransferRequest struct {
	Id int
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
	_, err := areAccountsFrozen(t)
	if err != nil {
		return false, errors.New("one or more of the accounts has been blocked")
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
		txn := createTxn(*t)
		t.FromAccount.Withdraw(t.Value)
		t.FromAccount.AddTxn(txn)
		t.ToAccount.Deposit(t.Value)
		t.ToAccount.AddTxn(txn)
		return true, nil
	}
}

func areAccountsFrozen(t Transfer) (bool, error) {
	if t.FromAccount.Blocked || t.ToAccount.Blocked {
		return true, errors.New("one of the accounts is blocked from transactions")
	}
	return false, nil
}

func createTxn (t Transfer) (transaction.Transaction) {
	txn := transaction.Transaction{TxnId:rand.Int(), ExternalSource:false, FromAccountId:t.FromAccount.ID, ToAccountId:t.ToAccount.ID,
		Amount:t.Value, Date:time.Now().UTC()}
	return txn
}
