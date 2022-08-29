package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"example/Go/models"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userName := vars["name"]
	userAge, er := strconv.ParseUint(vars["age"], 10, 8)
	userAddr := vars["address"]

	if er != nil {
		log.Fatalf("Error in parsing into Uint datatype: %s", er)
	}
	var user models.User = models.User(models.User{
		Name:    userName,
		Age:     uint8(userAge),
		Address: userAddr,
	})

	w.WriteHeader(http.StatusOK)
	// w.Write(payload)
	json.NewEncoder(w).Encode(user)
}
