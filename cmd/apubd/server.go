package main

import (
	"github.com/gin-gonic/gin"
	"github.com/majestrate/apub"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type APubDemo struct {
}

func (demo *APubDemo) FindUser(query string) (*apub.UserInfo, error) {
	logrus.Infof("find user: %s", query)
	return nil, nil
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	addr := os.Args[1]
	var demo APubDemo
	var handler apub.APubHandler
	r := gin.Default()
	handler.Setup(&demo, func(path string, handler http.Handler) {
		r.Any(path, gin.WrapH(handler))
	})
	logrus.Infof("listening on %s", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		logrus.Errorf("http.ListenAndServe(): %s", err.Error())
	}
}
