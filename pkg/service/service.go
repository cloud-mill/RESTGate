package service

import (
	"RESTGate/pkg/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type RESTGateService struct {
	Port             int
	ObservedServices []models.Service
}

func (RESTGateService *RESTGateService) StartRESTGate() {
	r := NewRouter(RESTGateService.ObservedServices)
	http.Handle("/", r)

	logrus.Infof("Starting RESTGateService at %d", RESTGateService.Port)
	err := http.ListenAndServe(":"+strconv.Itoa(RESTGateService.Port), nil)

	if err != nil {
		logrus.Panicf("error occurred starting RESTGateService: %v", err)
	}
}
