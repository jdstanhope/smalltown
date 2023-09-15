package main

import (
	"fmt"
	"net/http"
)

func (app *application) createPageHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintln(writer, "creating a new page from a picture")
}

func (app *application) showPageHandler(writer http.ResponseWriter, request *http.Request) {
	id, err := app.readIDParam(request)
	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}
	_, _ = fmt.Fprintf(writer, "show details of page %d\n", id)
}
