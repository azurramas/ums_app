package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/azurramas/ums_app/models"
	"github.com/azurramas/ums_app/services"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type Users struct{}

func (u Users) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := models.ListUsers()
	if err != nil {
		services.WriteError(w, err)
		return
	}

	services.WriteResponse(w, users, http.StatusOK)
}

func (u Users) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	user, err := models.GetUser(id)
	if err != nil {
		services.WriteError(w, err)
		return
	}

	user.Password = ""

	services.WriteResponse(w, user, http.StatusOK)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorMsg := errors.New("error while decoding request body")
		services.WriteError(w, errorMsg)
		return
	}

	//Data validation
	if user.Username == "" || user.FirstName == "" || user.LastName == "" ||
		user.Email == "" || user.Password == "" || user.Status == "" {
		errorMsg := errors.New("mandatory data missing")
		services.WriteError(w, errorMsg)
		return
	}

	//Check if user exists with same username or email
	userExists, err := user.ExistsByUsernameOrEmail()
	if err != nil {
		services.WriteError(w, err)
		return
	}

	if userExists {
		errorMsg := errors.New("user already exists")
		services.WriteError(w, errorMsg)
		return
	}

	//Add password hashing
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		services.WriteError(w, err)
		return
	}

	user.Password = string(bytes)

	err = user.Create()
	if err != nil {
		services.WriteError(w, err)
		return
	}

	services.WriteResponse(w, user, http.StatusOK)
}

func (u Users) Edit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorMsg := errors.New("error while decoding request body")
		services.WriteError(w, errorMsg)
		return
	}

	userExists, err := models.UserExistsByID(user.ID)
	if err != nil {
		services.WriteError(w, err)
		return
	}

	if !userExists {
		errorMsg := errors.New("record does not exist")
		services.WriteError(w, errorMsg)
		return
	}

	//Data validation
	if user.Username == "" || user.FirstName == "" || user.LastName == "" ||
		user.Email == "" || user.Status == "" {
		errorMsg := errors.New("mandatory data missing")
		services.WriteError(w, errorMsg)
		return
	}

	//Check if user exists with same username or email
	userExists, err = user.ExistsByUsernameOrEmailWhereID()
	if err != nil {
		services.WriteError(w, err)
		return
	}

	if userExists {
		errorMsg := errors.New("username or email already in use")
		services.WriteError(w, errorMsg)
		return
	}

	err = user.Update()
	if err != nil {
		services.WriteError(w, err)
		return
	}

	services.WriteResponse(w, user, http.StatusOK)

}

func (u Users) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	user, err := models.GetUser(id)
	if err != nil {
		services.WriteError(w, err)
		return
	}

	err = user.Delete()
	if err != nil {
		services.WriteError(w, err)
		return
	}

	services.WriteResponse(w, "Successfuly deleted user.", http.StatusOK)
}
