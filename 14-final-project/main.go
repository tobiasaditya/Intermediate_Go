package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type WebSocketConnection struct {
	*websocket.Conn
	Name string
}

var connections = make([]*WebSocketConnection, 0)

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

	e.Any("/ws", func(ctx echo.Context) error {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		currentSession, err := upgrader.Upgrade(ctx.Response().Writer, ctx.Request(), nil)

		if err != nil {
			return ctx.String(http.StatusBadRequest, "Could not open web socket!")
		}

		name := ctx.QueryParams().Get("name")
		currentConn := WebSocketConnection{Conn: currentSession, Name: name}
		connections = append(connections, &currentConn)

		return nil
	})

	e.Start(":8080")
}
