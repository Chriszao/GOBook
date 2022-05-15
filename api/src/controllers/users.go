package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func InsertUser(writer http.ResponseWriter, request *http.Request) {
	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUserRepository(db)

	userId, err := repository.Create(user)

	if err != nil {
		log.Fatal(err)
	}

	createdUserId := map[string]uint64{
		"id": userId,
	}

	writer.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(writer).Encode(createdUserId); err != nil {
		log.Fatal(err)
	}
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
