package controllers

import "net/http"

func RegisterControllers() {
	ac := newAlertController()

	http.Handle("/alerts", *ac)
	http.Handle("/alerts/", *ac)
}
