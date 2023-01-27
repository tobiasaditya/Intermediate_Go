package main

import (
	"8-ldap/config"
	"8-ldap/repository"
	"8-ldap/router"
	"8-ldap/service"
	"fmt"
	"net/http"

	"github.com/go-ldap/ldap"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

const webServerPort = "9000"

func main() {

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", config.LdapServer, config.LdapPort))
	if err != nil {
		log.Panic(err)
	}

	defer l.Close()
	err = l.Bind(config.LdapBindDN, config.LdapPassword)
	if err != nil {
		log.Panic(err)
	}

	ldapRepo := repository.NewLDAPRepo(l)
	loginService := service.NewLoginService(ldapRepo)
	routerHandler := router.NewRouterHandler(loginService)

	r := mux.NewRouter()
	r.HandleFunc("/", routerHandler.MainPageRouter)
	r.HandleFunc("/login", routerHandler.LoginRouter)

	baseUrl := "0.0.0.0:" + webServerPort
	fmt.Println("server started at", baseUrl)
	http.ListenAndServe(baseUrl, r)

}
