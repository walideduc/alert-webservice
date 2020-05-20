package controllers

import (
	"net/http"
	"regexp"
)

type alertController struct {
	alertIDPattern *regexp.Regexp
}

func (uc alertController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the Alert Controller"))
}

func newAlertController() *alertController {
	return &alertController{
		alertIDPattern: regexp.MustCompile(`^/alerts/(\d+)/?`),
	}
}
