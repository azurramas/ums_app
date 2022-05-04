package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/azurramas/ums_app/models"
	"github.com/azurramas/ums_app/services"
	"github.com/julienschmidt/httprouter"
)

type Permissions struct{}

func (p Permissions) Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var permission models.Permission

	err := json.NewDecoder(r.Body).Decode(&permission)
	if err != nil {
		errorMsg := errors.New("error while decoding request body")
		services.WriteError(w, errorMsg)
		return
	}

	//Data validation
	if permission.Description == "" || permission.Code == "" {
		errorMsg := errors.New("mandatory data missing")
		services.WriteError(w, errorMsg)
		return
	}

	userExists, err := models.UserExistsByID(permission.UserID)
	if err != nil {
		services.WriteError(w, err)
		return
	}

	if !userExists {
		errorMsg := errors.New("user does not exist")
		services.WriteError(w, errorMsg)
		return
	}

	err = permission.Add()
	if err != nil {
		services.WriteError(w, err)
		return
	}

	services.WriteResponse(w, permission, http.StatusOK)
}

func (p Permissions) Remove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	permExists, _ := models.PermissionExists(id)
	if !permExists {
		errorMsg := errors.New("record does not exist")
		services.WriteError(w, errorMsg)
		return
	}

	err := models.DeletePermissionByID(id)
	if err != nil {
		services.WriteError(w, err)
		return
	}

	services.WriteResponse(w, "Successfuly deleted permission.", http.StatusOK)
}
