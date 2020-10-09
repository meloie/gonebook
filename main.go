package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	"gonebook/ent"
)

var client *ent.Client

func init() {
	var err error
	client, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		logrus.Fatalf("failed opening connection to sqlite: %v", err)
	}
}
func createUser(ctx context.Context, client *ent.Client, username string, password string) (*ent.User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	u, err := client.User.Create().SetUsername(username).SetPassword(string(pass)).Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}
	return u, nil
}

type user struct {
	Username string
	Password string
}

func newUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var data user
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		logrus.WithError(err).Fatal("failed to decode request body")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Error!"))
		return
	}
	u, err := createUser(ctx, client, data.Username, data.Password)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create user")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Error!"))
		return
	}
	w.Write([]byte(fmt.Sprintf("user with ID %d created", u.ID)))
}

func main() {
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/", newUser)

	http.ListenAndServe(":3000", r)
}
