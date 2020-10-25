package main

import (
	"context"
	"gonebook/app/webserver"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"

	"gonebook/internal/ent"
	"gonebook/internal/services"
)

func main() {
	client, err := ent.Open("sqlite3", "file:db.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatal().Err(err).Msg("failed opening connection to sqlite")
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

	service := services.Service{Database: client}
	webserver := webserver.NewWebServer(&service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/register", webserver.RegisterUser)
	r.Post("/login", webserver.Login)

	r.Route("/contacts", func(r chi.Router) {
		r.Get("/", webserver.ContactsList)
		r.Post("/", webserver.CreateContact)
		r.Get("/{contactId}", webserver.ContactDetails)
		r.Put("/{contactId}", webserver.UpdateContact)
	})

	http.ListenAndServe(":3000", r)
}
