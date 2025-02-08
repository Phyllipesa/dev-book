package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// request body
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	// unmarshal request body
	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	// prepare user data
	if erro = user.Prepare(); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	// create a db connection
	db, erro := db.Connection()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	// close db connection after function ends
	defer db.Close()

	// passing db connection to repository
	repository := repository.NewRepositoryUsers(db)
	user.ID, erro = repository.Create(user)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// FindAllWithNameOrNick finds all users with name of nick
func FindAllWithNameOrNick(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := db.Connection()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsers(db)
	users, erro := repository.FindAllWithNameOrNick(nameOrNick)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, users)
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
