package service

import (
	"RESTGate/internal/models"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func NewRouter(services []models.Service) *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	for _, service := range services {
		for _, route := range service.Routes {
			r.Methods(strings.Split(route.Methods, ",")...).
				Path(route.Pattern).
				Name(route.Name).
				Handler(http.HandlerFunc(ReverseHandlerFactory(service.ServiceUrl)))
		}
	}

	return r
}
