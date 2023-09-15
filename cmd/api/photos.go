package main

import (
	"fmt"
	"github.com/jdstanhope/smalltown/internal/data"
	"net/http"
	"time"
)

func (app *application) createPhotoHandler(writer http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(writer, "upload a new photo")
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
