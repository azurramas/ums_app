package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Users struct{}

func (u Users) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (u Users) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (u Users) Edit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (u Users) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
