package controllers

import (
	"api/src/config"
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

	if err := user.Prepare("new"); err != nil {
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

	user.Password = ""

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

	// if users slice length is equals to zero, init an empty slice
	// with length zero, to return an empty array, instead nil
	if len(users) == 0 {
		users = []models.User{}
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
	params := mux.Vars(request)

	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	tokeUserId, err := config.ExtractUserId(request)

	if err != nil {
		responses.Error(writer, http.StatusUnauthorized, err)
		return
	}

	if userId != tokeUserId {
		responses.Error(
			writer,
			http.StatusForbidden,
			errors.New("it is not allowed to update a user that is not yours"),
		)
		return
	}

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

	if err = user.Prepare("edit"); err != nil {
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

	if err = repository.Update(userId, user); err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(writer, http.StatusNoContent, nil)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	tokeUserId, err := config.ExtractUserId(request)

	if err != nil {
		responses.Error(writer, http.StatusUnauthorized, err)
		return
	}

	if userId != tokeUserId {
		responses.Error(
			writer,
			http.StatusForbidden,
			errors.New("it is not allowed to delete a user that is not yours"),
		)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	if err = repository.Delete(userId); err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(writer, http.StatusNoContent, nil)
}

func FollowUser(writer http.ResponseWriter, request *http.Request) {

	followerId, err := config.ExtractUserId(request)

	if err != nil {
		responses.Error(writer, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(request)

	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	if userId == followerId {
		responses.Error(writer, http.StatusForbidden, errors.New("it is not allowed to follow yourself"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	if err := repository.FollowUser(userId, followerId); err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(writer, http.StatusNoContent, nil)
}

func UnFollowUser(writer http.ResponseWriter, request *http.Request) {
	followerId, err := config.ExtractUserId(request)

	if err != nil {
		responses.Error(writer, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(request)

	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	if userId == followerId {
		responses.Error(writer, http.StatusForbidden, errors.New("it is not allowed to unfollow yourself"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	if err := repository.UnFollowUser(userId, followerId); err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(writer, http.StatusNoContent, nil)
}
