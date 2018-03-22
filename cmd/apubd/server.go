package main

import (
	"github.com/gin-gonic/gin"
	"github.com/majestrate/apub"
	"github.com/majestrate/apub/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
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
	demo.User.GetPosts = func(int64, int) ([]*apub.Post, error) {
		now := time.Now()
		return []*apub.Post{
			&apub.Post{
				Message: "test post",
				From:    user.String(),
				Posted:  now,
				Updated: now,
				Self:    "https://" + user.Server() + "/posts/1",
			},
		}, nil
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

// implements apub.UserFinder
func (demo *APubDemo) LocalUser(query string) (user *apub.UserInfo, err error) {
	logrus.Infof("find user: %s", query)
	queryUser := apub.UserName("@" + query)
	if queryUser.Server() == "" {
		queryUser = apub.UserName("@" + query + "@" + demo.User.ServerName)
	}
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
	demo.Finder = &demo
	demo.InitLocalUser(user)

	err := demo.EnsureKeys()
	if err != nil {
		logrus.Fatalf("failed to ensure private key: %s", err.Error())
	}
	// set up apub routes
	demo.Setup(func(path string, handler http.Handler) {
		r.Any(path, gin.WrapH(handler))
	}, func(subpath string, handler http.Handler) {
		r.Group(subpath).Any("/:username", gin.WrapH(handler))
	})

	// serve profile page
	r.GET("/profile/:username", func(c *gin.Context) {
		username := c.Param("username")
		user, err := demo.Finder.LocalUser(username)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if user == nil {
			c.String(http.StatusNotFound, "no such user")
			return
		}
		c.HTML(http.StatusOK, "profile.html", gin.H{
			"user": user,
		})
	})

	logrus.Infof("listening on %s", addr)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		logrus.Errorf("http.ListenAndServe(): %s", err.Error())
	}
}
