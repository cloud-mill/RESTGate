package models

import "net/http"

type Route struct {
	Name        string `yaml:"name"`
	Methods     string `yaml:"methods"`
	Pattern     string `yaml:"pattern"`
	HandlerFunc http.HandlerFunc
}

type Service struct {
	Name       string  `yaml:"name"`
	ServiceUrl string  `yaml:"service_url"`
	Routes     []Route `yaml:"routes"`
}
