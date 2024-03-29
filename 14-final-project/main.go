package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

const TYPE_NEW_USER = "NEW_USER"
const TYPE_CHAT = "CHAT"
const TYPE_DISCONNECT = "DISCONNECT"

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	From    string
	Type    string
	Message string
}
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
		fmt.Println(name)
		currentConn := WebSocketConnection{Conn: currentSession, Name: name}
		connections = append(connections, &currentConn)

		go handleIO(&currentConn, connections)

		return nil
	})

	e.Start(":8080")
}

func handleIO(currentConn *WebSocketConnection, connections []*WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	broadcastMessage(currentConn, TYPE_NEW_USER, "")

	for {
		payload := SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				broadcastMessage(currentConn, TYPE_DISCONNECT, "")
				ejectConnection(currentConn)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		broadcastMessage(currentConn, TYPE_CHAT, payload.Message)
	}
}

func ejectConnection(currentConn *WebSocketConnection) {
	var newConn []*WebSocketConnection
	for _, conn := range connections {
		if conn != currentConn {
			newConn = append(newConn, conn)
		}
	}
	connections = newConn
}

func broadcastMessage(currentConn *WebSocketConnection, kind, message string) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		eachConn.WriteJSON(SocketResponse{
			From:    fmt.Sprintf(currentConn.Name),
			Type:    kind,
			Message: message,
		})
	}
}
