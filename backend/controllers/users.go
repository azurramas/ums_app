package controllers

import (
	"net/http"

	"github.com/azurramas/ums_app/services"
	"github.com/julienschmidt/httprouter"
)

type Users struct{}

func (u Users) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	services.WriteResponse(w, "test", http.StatusOK)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (u Users) Edit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (u Users) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
