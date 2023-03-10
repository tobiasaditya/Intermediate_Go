package main

import (
	"12-middleware/customMiddleware"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type M map[string]interface{}

func main() {
	//Load config
	viper.AddConfigPath("./config")
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseGlob("./*.html"))

	e := echo.New()

	const CSRFTokenHeader = "X-CSRF-TOKEN"
	const CSRFKey = "csrf"

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header: " + CSRFTokenHeader,
		ContextKey:  CSRFKey,
	}))

	// e.Use(middlewareOne)
	// e.Use(middlewareTwo)
	// e.Use(echo.WrapMiddleware(middlewareNonEcho))
	e.Use(middlewareLogrus)
	e.HTTPErrorHandler = errorHandler

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}",
	// }))

	e.GET("/index", func(ctx echo.Context) error {
		fmt.Println("Masuk /index")

		return ctx.JSON(http.StatusOK, true)
	})

	private := e.Group("/private")
	private.Use(middleware.BasicAuth(customMiddleware.BasicAuthMiddleware))
	private.GET("/index", func(c echo.Context) (err error) {
		fmt.Println("threeeeee!")

		return c.JSON(http.StatusOK, true)
	})

	//CSRF
	e.GET("/ui/index", func(ctx echo.Context) error {
		data := make(M)
		data[CSRFKey] = ctx.Get(CSRFKey)
		return tmpl.Execute(ctx.Response(), data)
	})

	e.POST("/ui/sayhello", func(ctx echo.Context) error {
		data := make(M)
		if err := ctx.Bind(&data); err != nil {
			return err
		}
		message := fmt.Sprintf("hellow %s", data["name"])
		return ctx.JSON(http.StatusOK, message)
	})

	lock := make(chan error)
	go func(lock chan error) {
		lock <- e.Start(":9000")
	}(lock)

	time.Sleep(1 * time.Millisecond)
	makeLogEntry(nil).Warn("application started without ssl/tls enabled")

	err := <-lock
	if err != nil {
		makeLogEntry(nil).Panic("failed to start application")
	}
	// e.Logger.Fatal(e.Start(":9000"))
}

func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("Masuk middleware one")
		return next(ctx)
	}
}

func middlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("Masuk middleware two")
		return next(ctx)
	}
}

func middlewareNonEcho(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Masuk middleware non echo")
		next.ServeHTTP(w, r)
	})
}

func middlewareLogrus(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		makeLogEntry(ctx).Info("Incoming request")
		return next(ctx)
	}
}

func errorHandler(err error, ctx echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	makeLogEntry(ctx).Error(report.Message)
	ctx.HTML(report.Code, report.Message.(string))
}

func makeLogEntry(c echo.Context) *logrus.Entry {
	if c == nil {
		return logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:00"),
		})
	}

	return logrus.WithFields(logrus.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:00"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}
