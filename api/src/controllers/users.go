package controllers

import "net/http"

// CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
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
