package main

import (
	"banksim/account"
	a "banksim/account"
	"banksim/transaction"
	"banksim/transfer"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func moveMoney(w http.ResponseWriter, r*http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var transferReq transfer.TransferRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&transferReq)
	if err != nil {
		panic(err)
	}

	_, transfer := parseTransfer(transferReq.FromAccountId, transferReq.ToAccountId, transferReq.Value)

	if _, err := transfer.Move(); err != nil {
		fmt.Println(err)
	}
}

func deposit(w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var deposit account.Deposit
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&deposit)
	if err != nil {
		panic(err)
	}
	acc := a.Accounts[deposit.AccountId]
	acc.ProcessDeposit(deposit)
}

func showAllAccounts(w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a.Accounts)
}

func findAccount (w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	accId, _ := strconv.Atoi(params["id"])
	account := a.Accounts[accId]
	json.NewEncoder(w).Encode(account)
}

func parseTransfer(fromAccountID int, toAccountID int, value int) (error, transfer.Transfer) {
	fromAccount, err := a.FindAccountById(fromAccountID)
	toAccount, err := a.FindAccountById(toAccountID)
	if err != nil {
		log.Fatal("Error finding account")
	}
	transfer := transfer.Transfer{fromAccount, toAccount, value}
	return err, transfer
}

func main() {

	a.Accounts[1] = a.Account{ID: 1, Balance: 200, Blocked: false, Transactions: make(map[int]transaction.Transaction)}
	a.Accounts[2] = a.Account{ID: 2, Balance: 4550, Blocked: false, Transactions: make(map[int]transaction.Transaction)}
	a.Accounts[3] = a.Account{ID: 3, Balance: 12220, Blocked: true, Transactions: make(map[int]transaction.Transaction)}
	a.Accounts[4] = a.Account{ID: 4, Balance: 4, Blocked: true, Transactions: make(map[int]transaction.Transaction)}
	a.Accounts[5] = a.Account{ID: 5, Balance: 40, Blocked: false, Transactions: make(map[int]transaction.Transaction)}
	a.Accounts[6] = a.Account{ID: 6, Balance: 35, Blocked: false, Transactions: make(map[int]transaction.Transaction)}
	a.Accounts[7] = a.Account{ID: 7, Balance: 365, Blocked: false, Transactions: make(map[int]transaction.Transaction)}
	a.Accounts[8] = a.Account{ID: 8, Balance: 40, Blocked: false, Transactions: make(map[int]transaction.Transaction)}

	router := mux.NewRouter()
	router.HandleFunc("/moveMoney", moveMoney).Methods("POST")
	router.HandleFunc("/deposit", deposit).Methods("POST")
	router.HandleFunc("/accounts", showAllAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}", findAccount).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}