package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
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

func main() {
	r := echo.New()
	store := newCookieStore()
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
