package apiControllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/testapiping", TestApiPing)
	myRouter.HandleFunc("/users", GetAllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{firstname}", GetUserByFirstName).Methods("GET")
	myRouter.HandleFunc("/user", GetUser).Methods("POST")
	myRouter.HandleFunc("/user/register", RegisterUser).Methods("POST")
	myRouter.HandleFunc("/deposit", Deposit).Methods("POST")
	myRouter.HandleFunc("/withdraw", Withdraw).Methods("POST")
	myRouter.HandleFunc("/transfer", Transfer).Methods("POST")
	myRouter.HandleFunc("/getaccountbalance", GetAccountBalance).Methods("POST")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argumen
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
