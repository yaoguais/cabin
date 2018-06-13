package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/sirupsen/logrus"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		post := r.PostForm
		get := r.URL.Query()
		header := r.Header
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		logger := logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
			"header": header,
			"post":   post,
			"get":    get,
			"ip":     ip,
		})

		c := httptest.NewRecorder()
		next.ServeHTTP(c, r)

		response, _ := ioutil.ReadAll(c.Body)

		ct := c.HeaderMap.Get("Content-Type")
		respBody := ct
		if ct != "" {
			for _, v := range []string{"text", "json", "xml"} {
				if strings.Contains(ct, v) {
					respBody = string(response)
					break
				}
			}
		} else {
			respBody = string(response)
		}

		logger.WithFields(logrus.Fields{
			"response": respBody,
			"code":     c.Code,
		}).Info("http request")

		for k, v := range c.HeaderMap {
			w.Header()[k] = v
		}
		w.WriteHeader(c.Code)
		w.Write(response)
	})
}
