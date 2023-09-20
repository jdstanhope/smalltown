package main

import (
	"errors"
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

	photo, err := app.models.Photos.Get(id)
	if err != nil {
		switch {
		case err == data.ErrRecordNotFound:
			app.notFoundResponse(writer, request)
		default:
			app.serverErrorResponse(writer, request, err)
		}
		return
	}
	err = app.writeJSON(writer, http.StatusOK, photo, "photo", nil)
	if err != nil {
		app.serverErrorResponse(writer, request, err)
	}
}

func (app *application) updatePhotoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	photo, err := app.models.Photos.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
	}

	var input struct {
		Name string `json:"name"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	photo.Name = input.Name

	v := validator.New()
	if data.ValidatePhoto(v, photo); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Photos.Update(photo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, photo, "photo", nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
