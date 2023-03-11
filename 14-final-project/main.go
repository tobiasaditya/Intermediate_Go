package main

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		content, err := ioutil.ReadFile("template/chat.html")
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "could not open html")
		}

		return ctx.HTML(http.StatusOK, string(content))
	})

	e.Static("/template", "template")

	e.Start(":8080")
}
