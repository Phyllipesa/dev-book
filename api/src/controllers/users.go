package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// request body
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	// unmarshal request body
	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	// create a db connection
	db, erro := db.Connection()
	if erro != nil {
		log.Fatal(erro)
	}

	// passing db connection to repository
	repository := repository.NewRepositoryUsers(db)
	userID, erro := repository.Create(user)

	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("User created with id %d", userID)))
}

// FindAllUsers finds all users
func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find all users"))
}

// FindUserById finds a user by id
func FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find user by id"))
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
