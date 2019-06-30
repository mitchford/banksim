package main

import (
	a "bankSim/account"
	"bankSim/transfer"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

func parseTransfer(fromAccountID int, toAccountID int, value int) (error, transfer.Transfer) {
	fromAccount, err := a.FindAccountById(fromAccountID)
	toAccount, err := a.FindAccountById(toAccountID)
	if err != nil {
		log.Fatal("Error finding account")
	}
	transfer := transfer.Transfer{fromAccount, toAccount, value}
	return err, transfer
}

func showAllAccounts(w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a.Accounts)
}

func main() {

	a.Accounts[1] = a.Account{ID: 1, Balance: 200, Blocked: false}
	a.Accounts[2] = a.Account{ID: 2, Balance: 4550, Blocked: false}
	a.Accounts[3] = a.Account{ID: 3, Balance: 40, Blocked: true}

	router := mux.NewRouter()
	router.HandleFunc("/moveMoney", moveMoney).Methods("POST")
	router.HandleFunc("/accounts", showAllAccounts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}