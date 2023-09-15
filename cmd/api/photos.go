package main

import (
	"fmt"
	"github.com/jdstanhope/smalltown/internal/data"
	"net/http"
	"time"
)

func (app *application) createPhotoHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintln(writer, "upload a new photo")
}

func (app *application) showPhotoHandler(writer http.ResponseWriter, request *http.Request) {
	id, err := app.readIDParam(request)
	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}

	photo := data.Photo{
		ID:         id,
		CreatedAt:  time.Now(),
		Comment:    "A cat",
		StorageURL: "https://placekitten.com/320/320?image=5",
		UserID:     1,
	}

	err = app.writeJSON(writer, http.StatusOK, photo, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}
