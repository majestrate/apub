package main

import (
	"github.com/gin-gonic/gin"
	"github.com/majestrate/apub"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type APubDemo struct {
	apub.APubHandler
}

// implements apub.UserFinder
func (demo *APubDemo) LocalUser(query string) (user *apub.UserInfo, err error) {
	logrus.Infof("find user: %s", query)
	return
}

func (demo *APubDemo) ListFollowers(query string) (users []*apub.UserInfo, err error) {
	logrus.Infof("Find followers: %s", query)
	return
}

func (demo *APubDemo) ListFollowing(query string) (users []*apub.UserInfo, err error) {
	logrus.Infof("Find following: %s", query)
	return
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	r := gin.Default()
	addr := os.Args[1]

	var demo APubDemo
	demo.Finder = &demo
	demo.Setup(func(path string, handler http.Handler) {
		r.Any(path, gin.WrapH(handler))
	}, func(subpath string, handler http.Handler) {
		r.Group(subpath).Any("/", gin.WrapH(handler))
	})
	logrus.Infof("listening on %s", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		logrus.Errorf("http.ListenAndServe(): %s", err.Error())
	}
}
