package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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
	nameOrNick := strings.ToLower(request.URL.Query().Get("user"))

	db, err := database.Connect()

	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, err := repository.FindAll(nameOrNick)

	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
	}

	responses.JSON(writer, http.StatusOK, users)
}

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
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

	user, err := repository.FindById(userId)

	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(user.ID)

	if user.ID == 0 {
		responses.Error(writer, http.StatusNotFound, errors.New("user not found"))
		return
	}

	responses.JSON(writer, http.StatusOK, user)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("UpdatingUser"))
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("DeletingUser"))
}
