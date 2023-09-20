package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/pages", app.createPageHandler)
	router.HandlerFunc(http.MethodGet, "/v1/pages/:id", app.showPageHandler)
	router.HandlerFunc(http.MethodPost, "/v1/photos", app.createPhotoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/photos/:id", app.showPhotoHandler)
	router.HandlerFunc(http.MethodPut, "/v1/photos/:id", app.updatePhotoHandler)

	return app.recoverPanic(router)
}
