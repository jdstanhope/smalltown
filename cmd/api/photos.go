package main

import (
	"fmt"
	"github.com/jdstanhope/smalltown/internal/data"
	"github.com/jdstanhope/smalltown/internal/validator"
	"net/http"
	"time"
)

func (app *application) createPhotoHandler(writer http.ResponseWriter, request *http.Request) {
	var input struct {
		Name string `json:"name"`
	}
	// read input json
	err := app.readJSON(writer, request, &input)
	if err != nil {
		app.badRequestResponse(writer, request, err)
		return
	}

	// convert to model to validate
	photo := &data.Photo{
		ID:         1,
		CreatedAt:  time.Now(),
		Name:       input.Name,
		StorageURL: "https://placekitten.com/320/320?image=5",
		UserID:     1,
	}
	v := validator.New()
	if data.ValidatePhoto(v, photo); !v.Valid() {
		app.failedValidationResponse(writer, request, v.Errors)
		return
	}

	// insert into models db
	err = app.models.Photos.Insert(photo)
	if err != nil {
		app.serverErrorResponse(writer, request, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/photos/%d", photo.ID))

	// write response
	err = app.writeJSON(writer, http.StatusCreated, photo, "photo", headers)
	if err != nil {
		app.serverErrorResponse(writer, request, err)
	}
}

func (app *application) showPhotoHandler(writer http.ResponseWriter, request *http.Request) {
	id, err := app.readIDParam(request)
	if err != nil || id < 1 {
		app.notFoundResponse(writer, request)
		return
	}

	photo := data.Photo{
		ID:         id,
		CreatedAt:  time.Now(),
		Name:       "A cat",
		StorageURL: "https://placekitten.com/320/320?image=5",
		UserID:     1,
	}

	err = app.writeJSON(writer, http.StatusOK, photo, "photo", nil)
	if err != nil {
		app.serverErrorResponse(writer, request, err)
	}
}
