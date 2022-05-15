package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func InsertUser(writer http.ResponseWriter, request *http.Request) {
	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		responses.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(); err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	user.ID, err = repository.Create(user)

	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(writer, http.StatusOK, user)
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
