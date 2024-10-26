package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *applicationDependencies) routers() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/signUp", a.signUp)

	router.HandlerFunc(http.MethodGet, "/signUp/read", a.signUpRead)

	router.HandlerFunc(http.MethodPatch, "/signUp/update", a.signUpUpdate)

	router.HandlerFunc(http.MethodDelete, "/signUp/delete", a.signUpDelete)

	return router
}
