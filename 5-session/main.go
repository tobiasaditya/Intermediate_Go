package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/rs/cors"
)

const SESSION_ID = "test-session-id"

func newCookieStore() *sessions.CookieStore {
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7
	store.Options.HttpOnly = true

	return store
}

func newPostgresStore() *pgstore.PGStore {
	url := "postgres://postgresuser:postgrespassword@127.0.0.1:5432/postgres?sslmode=disable"
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}

	return store
}

func main() {
	r := echo.New()

	//Set up middleware for CORST
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://hacktiv8.com", "https://www.google.com"},
		AllowedMethods: []string{"POST", "GET"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-TOKEN"},
	})

	r.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	// store := newCookieStore()
	store := newPostgresStore()
	r.GET("/set", func(ctx echo.Context) error {
		session, err := store.Get(ctx.Request(), SESSION_ID)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		session.Values["message1"] = "hello"
		session.Values["message2"] = "world"
		err = session.Save(ctx.Request(), ctx.Response())
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.String(http.StatusOK, "session set")
	})

	r.GET("/get", func(ctx echo.Context) error {
		session, err := store.Get(ctx.Request(), SESSION_ID)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
		}

		if len(session.Values) == 0 {
			return ctx.String(http.StatusOK, "empty values for sessions")
		}

		data := fmt.Sprintf("%s %s", session.Values["message1"], session.Values["message2"])
		return ctx.String(http.StatusOK, data)

	})

	r.GET("/delete", func(ctx echo.Context) error {
		session, err := store.Get(ctx.Request(), SESSION_ID)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		session.Options.MaxAge = -1 //forced to be expired
		session.Save(ctx.Request(), ctx.Response())
		return ctx.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	r.Start("localhost:8080")
}
