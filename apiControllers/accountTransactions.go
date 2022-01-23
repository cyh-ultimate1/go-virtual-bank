package apiControllers

import (
	"encoding/json"
	"govirtualbank/controllers"
	"govirtualbank/models"
	"io/ioutil"
	"net/http"
)

func Deposit(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var tx models.APIAccountTransactionRequest
	err2 := json.Unmarshal(reqBody, &tx)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	err3 := controllers.DepositToAccount(tx.Amount, tx.ReceipientFirstName, tx.ReceipientLastName)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var tx models.APIAccountTransactionRequest
	err2 := json.Unmarshal(reqBody, &tx)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	err3 := controllers.WithdrawFromAccount(tx.Amount, tx.ReceipientFirstName, tx.ReceipientLastName)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var tx models.APIAccountTransactionRequest
	err2 := json.Unmarshal(reqBody, &tx)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	err3 := controllers.TransferToAccount(tx.Amount, tx.SourceFirstName, tx.SourceLastName, tx.ReceipientFirstName, tx.ReceipientLastName)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}
}
