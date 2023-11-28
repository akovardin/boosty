package boosty

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Blog struct {
	Owner struct {
		ID        int    `json:"id"`
		HasAvatar bool   `json:"hasAvatar"`
		Name      string `json:"name"`
		AvatarURL string `json:"avatarUrl"`
	} `json:"owner"`
	Title      string `json:"title"`
	IsReadOnly bool   `json:"isReadOnly"`
	Flags      struct {
		ShowPostDonations      bool `json:"showPostDonations"`
		AllowGoogleIndex       bool `json:"allowGoogleIndex"`
		HasTargets             bool `json:"hasTargets"`
		AcceptDonationMessages bool `json:"acceptDonationMessages"`
		AllowIndex             bool `json:"allowIndex"`
		IsRssFeedEnabled       bool `json:"isRssFeedEnabled"`
		HasSubscriptionLevels  bool `json:"hasSubscriptionLevels"`
	} `json:"flags"`
	SignedQuery         string      `json:"signedQuery"`
	IsBlackListedByUser bool        `json:"isBlackListedByUser"`
	IsSubscribed        bool        `json:"isSubscribed"`
	Subscription        interface{} `json:"subscription"`
	IsTotalBaned        bool        `json:"isTotalBaned"`
	AccessRights        struct {
		CanSetPayout      bool `json:"canSetPayout"`
		CanCreateComments bool `json:"canCreateComments"`
		CanEdit           bool `json:"canEdit"`
		CanView           bool `json:"canView"`
		CanDeleteComments bool `json:"canDeleteComments"`
		CanCreate         bool `json:"canCreate"`
	} `json:"accessRights"`
	Count struct {
		Subscribers int `json:"subscribers"`
		Posts       int `json:"posts"`
	} `json:"count"`
	BlogURL                string   `json:"blogUrl"`
	IsOwner                bool     `json:"isOwner"`
	PublicWebSocketChannel string   `json:"publicWebSocketChannel"`
	SubscriptionKind       string   `json:"subscriptionKind"`
	IsBlackListed          bool     `json:"isBlackListed"`
	AllowedPromoTypes      []string `json:"allowedPromoTypes"`
	Description            []struct {
		Type        string `json:"type"`
		Content     string `json:"content"`
		Explicit    bool   `json:"explicit,omitempty"`
		URL         string `json:"url,omitempty"`
		Modificator string `json:"modificator,omitempty"`
	} `json:"description"`
	SocialLinks []struct {
		URL  string `json:"url"`
		Type string `json:"type"`
	} `json:"socialLinks"`
	HasAdultContent bool   `json:"hasAdultContent"`
	CoverURL        string `json:"coverUrl"`
}

func (b *Boosty) Blog() (*Blog, error) {
	u := fmt.Sprintf("/v1/blog/%s", b.blog)
	resp, err := b.request.Request(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("error on do request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("boosty blog status error")
	}

	res := &Blog{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, fmt.Errorf("boosty blog decode error: %w", err)
	}

	return res, nil
}
