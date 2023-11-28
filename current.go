package boosty

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	resp, err := b.request.Request(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("error on do request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("boosty current status error")
	}

	res := &Current{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, fmt.Errorf("boosty current decode error: %w", err)
	}

	return res, nil
}
