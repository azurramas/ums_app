package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/azurramas/ums_app/models"
	"github.com/azurramas/ums_app/services"
	"github.com/julienschmidt/httprouter"
)

type Permissions struct{}

func (p Permissions) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var permissions []models.Permission
	permissions, err := models.ListPermissions()
	if err != nil {
		services.WriteError(w, err)
		return
	}

	services.WriteResponse(w, permissions, http.StatusOK)
}

func (p Permissions) GetUserPerms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, _ := strconv.Atoi(ps.ByName("user_id"))
	var permissions []models.UserPermissions

	userExists, err := models.UserExistsByID(int64(userID))
	if err != nil {
		services.WriteError(w, err)
		return
	}

	if !userExists {
		errorMsg := errors.New("user does not exist")
		services.WriteError(w, errorMsg)
		return
	}

	permissions, err = models.GetUserPermissions(userID)
	if err != nil {
		services.WriteError(w, err)
		return
	}

	services.WriteResponse(w, permissions, http.StatusOK)
}

func (p Permissions) Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, _ := strconv.Atoi(ps.ByName("user_id"))
	permID, _ := strconv.Atoi(ps.ByName("perm_id"))

	userExists, err := models.UserExistsByID(int64(userID))
	if err != nil {
		services.WriteError(w, err)
		return
	}

	if !userExists {
		errorMsg := errors.New("user does not exist")
		services.WriteError(w, errorMsg)
		return
	}

	err = models.AddPermToUser(permID, userID)
	if err != nil {
		services.WriteError(w, err)
		return
	}

	services.WriteResponse(w, "Successfuly added permission.", http.StatusOK)
}

func (p Permissions) Remove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, _ := strconv.Atoi(ps.ByName("user_id"))
	permID, _ := strconv.Atoi(ps.ByName("perm_id"))

	userExists, err := models.UserExistsByID(int64(userID))
	if err != nil {
		services.WriteError(w, err)
		return
	}

	if !userExists {
		errorMsg := errors.New("user does not exist")
		services.WriteError(w, errorMsg)
		return
	}

	err = models.DeletePermissionByID(permID, userID)
	if err != nil {
		services.WriteError(w, err)
		return
	}

	services.WriteResponse(w, "Successfuly deleted permission.", http.StatusOK)
}
