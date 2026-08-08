package controller

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating users"))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User by ID"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user by ID"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating users"))
}
