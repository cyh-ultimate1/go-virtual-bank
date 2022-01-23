package apiControllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func TestApiPing(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("this is TestApiPing")
	json.NewEncoder(w).Encode("test api")
}
