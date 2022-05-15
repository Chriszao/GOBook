package controllers

import "net/http"

func InsertUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("InsertingUser"))
}

func FetchUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("FetchingAllUsers"))
}

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("GettingUserById"))
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("UpdatingUser"))
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("DeletingUser"))
}
