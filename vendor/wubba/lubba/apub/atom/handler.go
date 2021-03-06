package atom

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"wubba/lubba/apub"
	"wubba/lubba/apub/xml"
)

const RoutePath = "/apub/feeds"

const PerPage = 10

type Handler struct {
	apub.BaseHandler
}

func (h *Handler) RoutePath() string {
	return RoutePath
}

func (h *Handler) ServeUser(info apub.User, w http.ResponseWriter, r *http.Request) {
	log.Printf("serve nil atom feed: %s", info == nil)
	name := info.User()
	offsetStr := r.URL.Query().Get("offset")
	var offset int64
	if offsetStr != "" {
		off, _ := strconv.ParseInt(offsetStr, 64, 10)
		if off > 0 {
			offset = off
		}
	}
	var next string
	var feed apub.UserFeed
	posts, err := info.Posts(offset, PerPage)
	if err == nil {
		if len(posts) < PerPage {
			// no next url
		} else {
			nextURL, _ := url.Parse(r.URL.String())
			q := nextURL.Query()
			q.Set("offset", fmt.Sprintf("%d", (offset+1)*PerPage))
			nextURL.RawQuery = q.Encode()
			nextURL.Host = r.Host
			nextURL.Scheme = "https"
			next = nextURL.String()
		}
		feed, err = info.ToAtomFeed("atom feed for "+name, next)
		if err == nil {
			if err == nil {
				for idx := range posts {
					feed.AppendPost(posts[idx])
				}
			}
		}
	}
	if err == nil {
		w.Header().Set("Content-Type", apub.AtomMime)
		err = xml.MarshalHTTP(w, feed)
		if err != nil {
			log.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
	}
}
