package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
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

type User struct {
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=80"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	r := echo.New()

	r.Validator = &CustomValidator{validator: validator.New()}

	//Error handler
	r.HTTPErrorHandler = func(err error, ctx echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		castedObject, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not a valid email", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s must be greater than %s", err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s must be less than %s", err.Field(), err.Param())
				}
				break

			}

		}

		ctx.Logger().Error(report)
		ctx.JSON(report.Code, report)

	}

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

	//Wrap handler
	r.GET("/action/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	r.GET("/action/home", echo.WrapHandler(ActionHome))
	r.GET("/action/about", ActionAbout)

	//Static
	r.Static("/static", "assets")

	//Parsing payload
	r.Any("/user", func(ctx echo.Context) error {
		u := new(User)
		err := ctx.Bind(u)
		if err != nil {
			return err
		}

		if err = ctx.Validate(u); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, u)
	})

	r.Start("localhost:8080")
}
