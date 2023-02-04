package main

import "errors"

func main() {

}

func LuasPersegi(sisi int) int {
	return sisi * sisi
}

func Register(username string, password string) error {
	if username == "" {
		return errors.New("username empty")
	}

	if password == "" {
		return errors.New("password empty")
	}

	return nil
}
