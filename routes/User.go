package routes

import (
	"encoding/json"
	"example/Go/models"
	"example/Go/resolvers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request, resolvers *resolvers.Resolver) {
	var theUser models.User

	var userId string = mux.Vars(r)["id"]

	for _, up := range resolvers.Users {
		user, err := up.Get()

		if err != nil {
			log.Fatal(err)
		}
		if user.ID == userId {
			theUser = user
		}
	}

	json.NewEncoder(w).Encode(theUser)
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

func UpdateUserHandler(w http.ResponseWriter, r *http.Request, resolvers *resolvers.Resolver) {
	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
	// response body. A request body larger than that will now result in
	// Decode() returning a "http: request body too large" error.
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var reqBody models.User
	decErr := dec.Decode(&reqBody)

	if decErr != nil {
		log.Fatalf("Error in parsing user data: %s", decErr)
	}

	var userId string = mux.Vars(r)["id"]

	var response models.User
	for _, up := range resolvers.Users {
		user, err := up.Get()
		if err != nil {
			log.Fatalf("Error in retrieving a user: %s", err)
		}

		if user.ID == userId {
			user, _ = up.Update(models.User{
				ID:      user.ID,
				Name:    reqBody.Name,
				Age:     reqBody.Age,
				Address: reqBody.Address,
			})
			response = user
		}
	}

	json.NewEncoder(w).Encode(&response)
}
