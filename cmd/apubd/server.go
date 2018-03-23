package main

import (
	"github.com/gin-gonic/gin"
	"github.com/majestrate/apub"
	"github.com/majestrate/apub/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type APubDemo struct {
	apub.APubHandler
	User apub.UserInfo
}

func (demo *APubDemo) InitLocalUser(name string) {
	user := apub.UserName(name)
	demo.User.PreferedName = "test user"
	demo.User.Summary = "test user"
	demo.User.ServerName = user.Server()
	demo.User.UserName = user.User()
	demo.User.Avatar = "https://" + user.Server() + "/avatars/" + user.User() + ".jpeg"
	demo.User.Header = "https://" + user.Server() + "/headers/" + user.User() + ".jpeg"
	demo.User.Feed = "https://" + user.Server() + "/apub/feeds/" + user.User()
	demo.User.Profile = "https://" + user.Server() + "/profile/" + user.User()
	demo.User.GetPosts = func(offset int64, limit int) ([]*apub.Post, error) {
		return demo.LocalUserPosts(user.User(), offset, limit)
	}
}

func (demo *APubDemo) EnsureKeys() error {
	keyfile := demo.User.User() + ".pem"
	_, err := os.Stat(keyfile)
	if os.IsNotExist(err) {
		logrus.Info("regenerating signing key...")
		err = demo.User.RegenerateSigningKey()
		if err == nil {
			err = util.StorePrivateKey(keyfile, demo.User.SigningKey)
		}
	}
	demo.User.SigningKey, err = util.LoadPrivateKey(keyfile)
	return err
}

func (demo *APubDemo) LocalUserPosts(postid string, offset int64, limit int) (posts []*apub.Post, err error) {
	return
}

func (demo *APubDemo) LocalPost(postid string) (post *apub.Post, err error) {
	return
}

func (demo *APubDemo) LocalUser(query string) (user *apub.UserInfo, err error) {
	queryUser := apub.NormalizeUser(query, demo.User.ServerName)
	logrus.Infof("find user: %s", queryUser)
	if queryUser.String() == demo.User.User() {
		user = &demo.User
	}
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
	user := os.Args[2]

	var demo APubDemo
	demo.Database = &demo
	demo.InitLocalUser(user)

	err := demo.EnsureKeys()
	if err != nil {
		logrus.Fatalf("failed to ensure private key: %s", err.Error())
	}
	// set up apub routes
	demo.SetupRoutes(func(path string, handler http.Handler) {
		r.Any(path, gin.WrapH(handler))
	}, func(subpath string, handler http.Handler) {
		r.Group(subpath).Any("/:extra", gin.WrapH(handler))
	})

	// serve profile page
	r.GET("/profile/:username", func(c *gin.Context) {
		username := c.Param("username")
		user, err := demo.LocalUser(username)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if user == nil {
			c.String(http.StatusNotFound, "no such user")
			return
		}
		c.String(http.StatusOK, user.User())
	})

	logrus.Infof("listening on %s", addr)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		logrus.Errorf("http.ListenAndServe(): %s", err.Error())
	}
}
