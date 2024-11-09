package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	Addr string
}

var users = []User{}

func (s *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//encode the users slice to JSON
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create users"))

	//decode the request body to a User struct
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}
	validationError := insertUser(u)

	if validationError != nil {
		http.Error(w, validationError.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	//input validation
	if u.FirstName == "" || u.LastName == "" {
		return errors.New("first name and last name are required")
	}

	//duplicate validation
	for _, user := range users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return errors.New("User already exists")
		}
	}
	users = append(users, u)
	return nil
}
