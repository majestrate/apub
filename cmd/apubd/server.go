package main

import (
	"github.com/gin-gonic/gin"
	"github.com/majestrate/apub"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
)

func findUser(query string) (*url.URL, error) {
	logrus.Infof("find user: %s", query)
	return nil, nil
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	addr := os.Args[1]
	hostname := os.Args[2]
	var handler apub.APubHandler
	r := gin.Default()
	handler.Setup(findUser, hostname, func(path string, handler http.Handler) {
		r.Any(path, gin.WrapH(handler))
	})
	logrus.Infof("listening on %s as %s", addr, hostname)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		logrus.Errorf("http.ListenAndServe(): %s", err.Error())
	}
}
