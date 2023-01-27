package router

import (
	"8-ldap/service"
	"fmt"
	"net/http"
	"text/template"
)

type RouterHandler struct {
	loginService *service.LoginService
}

func NewRouterHandler(loginService *service.LoginService) *RouterHandler {
	return &RouterHandler{loginService: loginService}
}
func (rh *RouterHandler) MainPageRouter(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("page.html")
	if err := parsedTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rh *RouterHandler) LoginRouter(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	data, err := rh.loginService.Login(username, password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	message := fmt.Sprintf("Welcome %s\n", data.FullName)
	w.Write([]byte(message))
}
