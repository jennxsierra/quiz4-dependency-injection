package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("home handler hit")
	w.Write([]byte("Welcome (refactored, DI)\n"))
}
