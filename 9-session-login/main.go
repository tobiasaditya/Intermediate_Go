package main

import (
	"9-session-login/config"
	"9-session-login/repository"
	"9-session-login/router"
	"9-session-login/service"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

const SESSION_ID = "test-session-id"

func main() {
	db, err := sql.Open("postgres", config.DB_URI)
	if err != nil {
		fmt.Println(err.Error())
	}
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	sessionService := service.NewSessionService()

	userHandler := router.NewUserRouter(userService, sessionService)

	r := echo.New()
	r.Renderer = config.NewRenderer("template/*", true)
	r.GET("/", func(ctx echo.Context) error {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
		return nil
	})
	r.GET("/home", userHandler.HomeHandler)
	r.POST("/home", userHandler.HomeHandler)

	r.POST("/register", userHandler.RegisterHandler)
	r.POST("/login", userHandler.LoginHandler)
	r.GET("/login", userHandler.LoginHandler)
	r.POST("/logout", userHandler.LogoutHandler)
	r.Start(":8000")
}
