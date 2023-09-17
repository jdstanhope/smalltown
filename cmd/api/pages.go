package main

import (
	"fmt"
	"github.com/jdstanhope/smalltown/internal/data"
	"net/http"
	"time"
)

func (app *application) createPageHandler(writer http.ResponseWriter, request *http.Request) {
	var input struct {
		Name string `json:"name"`
	}
	err := app.readJSON(writer, request, &input)
	if err != nil {
		app.badRequestResponse(writer, request, err)
		return
	}
	_, _ = fmt.Fprintf(writer, "Got %+v\n\n", input)
}

func (app *application) showPageHandler(writer http.ResponseWriter, request *http.Request) {
	id, err := app.readIDParam(request)
	if err != nil || id < 1 {
		app.notFoundResponse(writer, request)
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
		app.serverErrorResponse(writer, request, err)
	}
}
