package router

import (
	"9-session-login/config"
	"9-session-login/dto"
	"9-session-login/entity"
	"9-session-login/service"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type HTMLTemplate struct {
	User entity.User
	Err  string
}

type UserRouter struct {
	userService    *service.UserService
	sessionService *service.SessionService
}

func NewUserRouter(userService *service.UserService, sessionService *service.SessionService) *UserRouter {
	return &UserRouter{userService: userService, sessionService: sessionService}
}

func (r *UserRouter) RegisterHandler(c echo.Context) error {
	var inputUser dto.InputUser
	c.Bind(&inputUser)

	newUser, err := r.userService.RegisterUser(inputUser)
	if err != nil {
		log.Warn(err)
		return c.Render(http.StatusOK, "login.html", HTMLTemplate{
			User: entity.User{},
			Err:  err.Error(),
		})
	}

	//Save session
	r.sessionService.SetSession(c, config.SESSION_ID, newUser.UserName)
	return c.Redirect(http.StatusTemporaryRedirect, "/home")
}

func (r *UserRouter) LoginHandler(c echo.Context) error {
	log.Info("Start Login Handler")
	username := c.FormValue("username")
	password := c.FormValue("password")

	inputLogin := dto.InputLogin{
		UserName: username,
		Password: password,
	}
	// authenticate in db
	user, err := r.userService.LoginUser(inputLogin)
	if err != nil {
		log.Warn(err)
		return c.Render(http.StatusOK, "login.html", HTMLTemplate{
			User: entity.User{},
			Err:  "Error when login",
		})
	}
	log.Info("Successfully authenticate user", user)

	//Save session
	r.sessionService.SetSession(c, config.SESSION_ID, user.UserName)

	return c.Redirect(http.StatusTemporaryRedirect, "/home")
}

func (r *UserRouter) HomeHandler(c echo.Context) error {
	username, err := r.sessionService.GetSession(c, config.SESSION_ID)
	if err != nil {
		log.Warn("error when retrieve session")
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return c.Render(http.StatusOK, "home.html", HTMLTemplate{
		User: entity.User{UserName: username},
		Err:  "",
	})
}
