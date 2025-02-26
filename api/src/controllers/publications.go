package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePublication add a publication in database
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userIdFromToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publication models.Publication
	if erro = json.Unmarshal(bodyRequest, &publication); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	publication.AuthorID = userIdFromToken

	if erro = publication.Prepare(); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryPublications(db)
	publication.ID, erro = repository.Create(publication)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}

// FindPublication find publications that would appear in the user's feed
func FindPublications(w http.ResponseWriter, r *http.Request) {
	userIdFromToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryPublications(db)
	publications, erro := repository.FindPublications(userIdFromToken)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, publications)
}

// FindPublicationById find a publication
func FindPublicationById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(parameters["publicationId"], 10, 64)
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

	repository := repository.NewRepositoryPublications(db)
	publication, erro := repository.FindById(publicationId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, publication)
}

// UpdatePublication update a publication
func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userIdFromToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(parameters["publicationId"], 10, 64)
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

	repository := repository.NewRepositoryPublications(db)
	publicationSavedInDB, erro := repository.FindById(publicationId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if publicationSavedInDB.AuthorID != userIdFromToken {
		responses.Error(w, http.StatusForbidden, errors.New(
			"invalid action: you cannot update publications that are not yours"),
		)
		return
	}

	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publication models.Publication
	if erro = json.Unmarshal(requestBody, &publication); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publication.Prepare(); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
	}

	if erro = repository.Update(publicationId, publication); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletePublication delete a publication
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	userIdFromToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(parameters["publicationId"], 10, 64)
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

	repository := repository.NewRepositoryPublications(db)
	publicationSavedInDB, erro := repository.FindById(publicationId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if publicationSavedInDB.AuthorID != userIdFromToken {
		responses.Error(w, http.StatusForbidden, errors.New(
			"invalid action: you cannot delete publications that are not yours"),
		)
		return
	}

	if erro = repository.Delete(publicationId); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FindPublicationsByUser find publications by user
func FindPublicationsByUser(w http.ResponseWriter, r *http.Request) {
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

	repository := repository.NewRepositoryPublications(db)
	publications, erro := repository.FindByUserId(userID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}

// LikePublication like an publication
func LikePublication(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, erro := strconv.ParseUint(parameters["publicationId"], 10, 64)
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

	repository := repository.NewRepositoryPublications(db)
	if erro = repository.Like(publicationID); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// UnlikePublication unlike an publication
func UnlikePublication(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, erro := strconv.ParseUint(parameters["publicationId"], 10, 64)
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

	repository := repository.NewRepositoryPublications(db)
	if erro = repository.Unlike(publicationID); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
