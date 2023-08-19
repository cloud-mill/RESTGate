package service

import (
	"RESTGate/internal/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type RESTGateService struct {
	Port             int              `yaml:"port"`
	ObservedServices []models.Service `yaml:"observed_services"`
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
