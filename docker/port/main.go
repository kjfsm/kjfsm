package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

func main() {
	tRouter := chi.NewRouter()
	tRouter.Use(middleware.Logger)

	tRouter.Get("/", func(tWriter http.ResponseWriter, tRequest *http.Request) {
		tWriter.Write([]byte("root"))
	})

	http.ListenAndServe(":3333", tRouter)
}
