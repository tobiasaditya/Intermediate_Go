package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type M map[string]interface{}

var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action index"))
}

var ActionHome = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("from action home"))
	},
)

var ActionAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("from action about"))
		},
	),
)

func main() {
	r := echo.New()

	r.GET("/index", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/html", func(ctx echo.Context) error {
		data := "<html><head></head><body><h1>Hello html</h1></body> </html>"
		return ctx.HTML(http.StatusOK, data)
	})

	r.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/index")
	})

	r.GET("/json", func(ctx echo.Context) error {
		data := M{"message": "hello world", "counter": "2"}
		return ctx.JSON(http.StatusOK, data)
	})

	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("Hello %s", name)
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s", name)
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")
		data := fmt.Sprintf("Hello %s. Message for you : %s", name, message)
		return ctx.String(http.StatusOK, data)
	})

	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")
		message = strings.Replace(message, "/", " ", 1)
		data := fmt.Sprintf("Hello %s. Message for you : %s", name, message)

		return ctx.String(http.StatusOK, data)
	})

	r.GET("/action/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	r.GET("/action/home", echo.WrapHandler(ActionHome))
	r.GET("/action/about", ActionAbout)

	r.Start("localhost:8080")
}
