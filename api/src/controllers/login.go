package controllers

import (
	"api/src/config"
	"api/src/database"
	"api/src/models"
	"api/src/providers"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		responses.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err := json.Unmarshal(requestBody, &user); err != nil {
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

	userExists, err := repository.FindByEmail(user.Email)

	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	if err := providers.ValidatePassword(user.Password, userExists.Password); err != nil {
		responses.Error(writer, http.StatusUnauthorized, err)
		return
	}

	token, err := config.GenerateToken(userExists.ID)

	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	writer.Write([]byte(token))

}
