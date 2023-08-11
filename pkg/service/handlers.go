package service

import (
	"RESTGate/pkg/reverse"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

func ReverseHandlerFactory(reverseTo string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("reversing to %s", reverseTo)

		path, err := url.Parse(reverseTo)
		if err != nil {
			logrus.Panicf("error when parsing url: %s, %v", reverseTo, err)
		}
		proxy := reverse.NewReverseProxy(path)
		proxy.ServeHTTP(w, r)
	}
}
