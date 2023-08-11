package main

import (
	"RESTGate/pkg/service"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

var RESTGateService service.RESTGateService

func init() {
	config, err := os.ReadFile(os.Getenv("RESTGATE_CONFIG_PATH"))
	if err != nil {
		logrus.Fatalf("error reading RESTGate config YAML file: %v", err)
	}

	err = yaml.Unmarshal(config, &RESTGateService)
	if err != nil {
		logrus.Fatalf("error parsing RESTGate config YAML file: %v", err)
	}
}

func main() {
	RESTGateService.StartRESTGate()
}
