package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/pages", app.createPageHandler)
	router.HandlerFunc(http.MethodGet, "/v1/pages/:id", app.showPageHandler)
	router.HandlerFunc(http.MethodPost, "/v1/photos", app.createPhotoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/photos/:id", app.showPhotoHandler)

	return router
}
