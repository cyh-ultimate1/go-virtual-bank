package apiControllers

import (
	"encoding/json"
	"govirtualbank/controllers"
	"govirtualbank/models"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(controllers.GetAllUsers())
}

func GetUserByFirstName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	firstname := vars["firstname"]
	json.NewEncoder(w).Encode(controllers.GetUserByFirstName(firstname))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user models.APIUser
	json.Unmarshal(reqBody, &user)

	userOut := controllers.GetUserByName(user.FirstName, user.LastName)
	if userOut.ID <= 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(userOut)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user models.APIUserRegister
	json.Unmarshal(reqBody, &user)

	err2 := controllers.RegisterUser(user)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
}
