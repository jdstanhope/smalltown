package main

import (
	"fmt"
	"github.com/jdstanhope/smalltown/internal/data"
	"net/http"
	"time"
)

func (app *application) createPageHandler(writer http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(writer, "creating a new page from a picture")
}

func (app *application) showPageHandler(writer http.ResponseWriter, request *http.Request) {
	id, err := app.readIDParam(request)
	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}

	page := data.Page{
		ID:         id,
		CreatedAt:  time.Now(),
		Name:       "A cat",
		UserID:     1,
		StorageURL: "https://placekitten.com/320/320?image=5",
		PhotoID:    1,
	}
	err = app.writeJSON(writer, http.StatusOK, page, "page", nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}
