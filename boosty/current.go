package boosty

import (
	"fmt"
	"net/http"
	"net/url"
)

type Current struct {
	PaidCount      int `json:"paidCount"`
	FollowersCount int `json:"followersCount"`
	Hold           int `json:"hold"`
	Income         int `json:"income"`
	Balance        int `json:"balance"`
	PayoutSum      int `json:"payoutSum"`
}

func (b *Boosty) Current() (*Current, error) {
	u := fmt.Sprintf("/v1/blog/stat/%s/current", b.blog)

	m := Method[Current]{
		request: b.request,
		method:  http.MethodGet,
		url:     u,
		values:  url.Values{},
	}

	return m.Call(Current{})
}
