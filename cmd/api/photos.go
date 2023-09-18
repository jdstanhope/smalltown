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
		Comment string `json:"comment"`
	}

	err := app.readJSON(writer, request, &input)
	if err != nil {
		app.badRequestResponse(writer, request, err)
		return
	}

	photo := &data.Photo{
		ID:         1,
		CreatedAt:  time.Now(),
		Comment:    input.Comment,
		StorageURL: "https://placekitten.com/320/320?image=5",
		UserID:     1,
	}
	v := validator.New()
	if data.ValidatePhoto(v, photo); !v.Valid() {
		app.failedValidationResponse(writer, request, v.Errors)
		return
	}

	_, _ = fmt.Fprintf(writer, "Got %+v\n\n", input)
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
		Comment:    "A cat",
		StorageURL: "https://placekitten.com/320/320?image=5",
		UserID:     1,
	}

	err = app.writeJSON(writer, http.StatusOK, photo, "photo", nil)
	if err != nil {
		app.serverErrorResponse(writer, request, err)
	}
}
