package logger

import (
	"net/http"

	"github.com/DaveBlooman/api-common/Godeps/_workspace/src/github.com/Sirupsen/logrus"
)

var log = logrus.New()

// Log is a HTTP logger abstraction
func Log(inner http.Handler, name string) http.Handler {
	log.Formatter = new(logrus.JSONFormatter)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logrus.Fields{
			"url":    r.RequestURI,
			"method": r.Method,
		}).Info(name)
		inner.ServeHTTP(w, r)
	})
}

// Info is a logger abstraction
func Info(message map[string]interface{}) {
	log.WithFields(message).Info()
}

// Error is a logger abstraction
func Error(message map[string]interface{}) {
	log.WithFields(message).Error("error")
}

// Warn is a logger abstraction
func Warn(message map[string]interface{}) {
	log.WithFields(message).Warn()
}

// Debug is a logger abstraction
func Debug(message map[string]interface{}) {
	log.WithFields(message).Debug()
}

// Panic is a logger abstraction
func Panic(message map[string]interface{}) {
	log.WithFields(message).Panic("panic")
}

// Fatal is a logger abstraction
func Fatal(message map[string]interface{}) {
	log.WithFields(message).Fatal("fatal")
}
