package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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
	if erro = user.Prepare("cadastro"); erro != nil {
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
	parameters := mux.Vars(r)

	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsers(db)
	user, erro := repository.FindUserById(userID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	userIdFromToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIdFromToken {
		responses.Error(w, http.StatusForbidden, errors.New(
			"you don't have permission to update another users"),
		)
		return
	}

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("edicao"); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsers(db)
	erro = repository.UpdateUser(userID, user)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	userIdFromToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIdFromToken {
		responses.Error(w, http.StatusForbidden, errors.New(
			"you don't have permission to delete another users"),
		)
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsers(db)
	if erro = repository.DeleteUser(userID); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FollowUser allows a user follow other
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if followerId == userID {
		responses.Error(w, http.StatusForbidden, errors.New(
			"invalid action: you cannot follow yourself"),
		)
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsers(db)
	if erro = repository.FollowUser(userID, followerId); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser allows a user unfollow other
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New(
			"invalid action: you cannot unfollow yourself"),
		)
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsers(db)
	if erro = repository.UnfollowUser(userID, followerID); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FindFollowers find all followers of a user
func FindFollowers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsers(db)
	followers, erro := repository.FindFollowers(userID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// FindFollowing find for all users you are following
func FindFollowing(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsers(db)
	following, erro := repository.FindFollowing(userID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, following)
}

// UpdatePassword update a user's password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIdFromToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if userIdFromToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New(
			"invalid action! You can't update another account's password"),
		)
		return
	}

	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var password models.Password
	if erro = json.Unmarshal(requestBody, &password); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsers(db)
	savedPassword, erro := repository.FindPassword(userID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(
		savedPassword,
		password.CurrentPassword,
	); erro != nil {
		responses.Error(
			w,
			http.StatusInternalServerError,
			errors.New("invalid password"))
		return
	}

	newPasswordWithHash, erro := security.Hash(password.NewPassword)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repository.UpdatePassword(userID, string(newPasswordWithHash)); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
