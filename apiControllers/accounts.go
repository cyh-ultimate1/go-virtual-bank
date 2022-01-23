package apiControllers

import (
	"encoding/json"
	"govirtualbank/controllers"
	"govirtualbank/models"
	"io/ioutil"
	"net/http"
)

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var userInput models.APIUser
	err2 := json.Unmarshal(reqBody, &userInput)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	var user = controllers.GetUserByName(userInput.FirstName, userInput.LastName)
	if user.ID <= 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var account = controllers.GetAccountByUserName(userInput.FirstName, userInput.LastName)
	if account.ID <= 0 {
		http.Error(w, "User account not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(account.GetBalance())
}

func GetAccountByUserName(w http.ResponseWriter, r *http.Request) {

}
