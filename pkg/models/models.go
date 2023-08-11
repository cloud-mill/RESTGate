package models

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Service struct {
	Name       string
	ServiceUrl string
	Routes     []Route
}
