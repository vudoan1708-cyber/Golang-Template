package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"example/Go/models"
	"example/Go/resolvers"
)

func HomeHandler(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }).Methods("GET")
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request, resolvers *resolvers.Resolver) {
	var allUsers []models.User
	for _, up := range resolvers.Users {
		u, err := up.Get()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		allUsers = append(allUsers, u)
	}
	json.NewEncoder(w).Encode(allUsers)
}

func AddUserHandler(w http.ResponseWriter, r *http.Request, resolvers *resolvers.Resolver) {
	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
	// response body. A request body larger than that will now result in
	// Decode() returning a "http: request body too large" error.
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var userModel models.User
	parsingErr := dec.Decode(&userModel)

	if parsingErr != nil {
		http.Error(w, parsingErr.Error(), http.StatusBadRequest)
		return
	}

	up, err := resolvers.UserFactory.NewUser(
		userModel.Name,
		userModel.Age,
		userModel.Address,
	)

	if err != nil {
		log.Fatalf("Error in creating a new user: %s", err)
	}

	// Get the user info
	user, userErr := up.Get()
	if userErr != nil {
		log.Fatalf("Error in retrieving user data: %s", userErr)
	}

	resolvers.Users = append(resolvers.Users, up)

	json.NewEncoder(w).Encode(user)
}

func UserHandler(r *mux.Router, resolvers *resolvers.Resolver) {
	r.HandleFunc("/user/{name}/{age}/{address}", Handler).Methods("GET")
	r.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		GetUsersHandler(res, req, resolvers)
	}).Methods("GET")
	r.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		AddUserHandler(res, req, resolvers)
	}).Methods("POST")
}
