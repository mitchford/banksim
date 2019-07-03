package account

import (
	txn "banksim/transaction"
	"errors"
	"math/rand"
	"time"
)

type Account struct  {
	ID int
	Balance	int
	Blocked bool
	Transactions map[int]txn.Transaction
}

type Deposit struct {
	ID int
	AccountId int
	DepositAmount int
	Credited bool
}

var Accounts = make(map[int]Account)

func FindAccountById(id int) (Account, error) {
	return Accounts[id], nil
}

func (a *Account) Deposit(amount int) {
	acc := Accounts[a.ID]
	acc.Balance += amount
	Accounts[a.ID] = acc
}

func (a *Account) Withdraw(amount int) {
	acc := Accounts[a.ID]
	acc.Balance -= amount
	Accounts[a.ID] = acc
}

func (a Account) HasAvailableFunds(amount int) bool{
	if a.Balance >= amount {
		return true
	}
	return false
}

func (a *Account) AddTxn (txn txn.Transaction) {
	a.Transactions[txn.TxnId] = txn
}

func (a *Account) ProcessDeposit (deposit Deposit) error {
	if !deposit.Credited {
		a.Balance += deposit.DepositAmount
		txn := txn.Transaction{TxnId:rand.Int(), ExternalSource:true, ToAccountId:a.ID, Amount:deposit.DepositAmount,
			Date:time.Now().UTC()}
		a.AddTxn(txn)
		deposit.Credited = true
		return nil
	}
	return errors.New("deposit already processed")
}