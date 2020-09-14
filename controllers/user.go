package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from the User controller!"))
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile("lipu"),

	}
}