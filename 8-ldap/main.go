package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/gommon/log"
)

const webServerPort = 9000

const view = `<html>
	<head>
		<title>Template</title>
	</head>
	<body>
		<form method="post" action="/login">
			<div>
				<label>username</label>
				<input type="text" name="username" required/>
			</div>
			<div>
				<label>password</label>
				<input type="password" name="password" required/>
			</div>
			<button type="submit">Login</button>
		</form>
	</body>
	</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("main-template").Parse(view))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		ok, data, err := AuthUsingLDAP(username, password)
		if !ok {
			log.Error("auth using ldap not ok")
			http.Error(w, "invalid username/password", http.StatusUnauthorized)
			return
		}
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		message := fmt.Sprintf("Welcome %s\n", data.FullName)
		w.Write([]byte(message))
	})

	portString := fmt.Sprintf(":%d", webServerPort)
	fmt.Println("server started at", portString)
	http.ListenAndServe(portString, nil)

}
