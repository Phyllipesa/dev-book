package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
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

}

// DeletePublication delete a publication
func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
