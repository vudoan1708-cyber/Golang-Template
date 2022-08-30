package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"example/Go/resolvers"
)

func HomeHandler(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }).Methods("GET")
}

func UserHandler(r *mux.Router, resolvers *resolvers.Resolver) {
	r.HandleFunc("/user/{id}", func(res http.ResponseWriter, req *http.Request) {
		GetUserHandler(res, req, resolvers)
	}).Methods("GET")
	r.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		GetUsersHandler(res, req, resolvers)
	}).Methods("GET")
	r.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		AddUserHandler(res, req, resolvers)
	}).Methods("POST")
	r.HandleFunc("/user/{id}", func(res http.ResponseWriter, req *http.Request) {
		UpdateUserHandler(res, req, resolvers)
	}).Methods("PUT")
}
