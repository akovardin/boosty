package boosty

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Option func(b *Boosty)

type Boosty struct {
	base   string
	blog   string
	token  string
	client *http.Client
}

func New(blog, token string, options ...Option) *Boosty {
	b := &Boosty{
		base:   "https://api.boosty.to",
		blog:   blog,
		token:  token,
		client: &http.Client{},
	}

	for _, o := range options {
		o(b)
	}

	return b
}

func WithClient(client *http.Client) Option {
	return func(b *Boosty) {
		b.client = client
	}
}

func (b *Boosty) Stats() (*Stats, error) {
	u := fmt.Sprintf("%s/v1/blog/stat/%s/current", b.base, b.blog)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("boosty stats request error: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+b.token)

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("boosty stats do error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("boosty stats status error")
	}

	res := &Stats{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, fmt.Errorf("boosty stats decode error: %w", err)
	}

	return res, nil
}

type SubscriptionsResponse struct {
	Offset int            `json:"offset"`
	Total  int            `json:"total"`
	Limit  int            `json:"limit"`
	Data   []Subscription `json:"data"`
}

func (b *Boosty) Subscriptions(offset, limit int) ([]Subscription, error) {
	u := fmt.Sprintf("%s/v1/blog/%s/subscription_level/?show_free_level=true&sort_by=on_time&offset=%d&limit=%d&order=gt", b.base, b.blog, offset, limit)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("boosty subscriptions request error: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+b.token)

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("boosty subscriptions do error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("boosty subscriptions status error")
	}

	res := SubscriptionsResponse{
		Data: []Subscription{},
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("boosty subscriptions decode error: %w", err)
	}

	return res.Data, nil
}

type SubscribersResponse struct {
	Offset int          `json:"offset"`
	Total  int          `json:"total"`
	Limit  int          `json:"limit"`
	Data   []Subscriber `json:"data"`
}

func (b *Boosty) Subscribers(offset, limit int) ([]Subscriber, error) {
	u := fmt.Sprintf("%s/v1/blog/%s/subscribers?sort_by=on_time&offset=%d&limit=%d&order=gt", b.base, b.blog, offset, limit)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("boosty subscribers request error: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+b.token)

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("boosty subscribers do error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("boosty subscribers status error")
	}

	res := SubscribersResponse{
		Data: []Subscriber{},
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("boosty subscribers decode error: %w", err)
	}

	return res.Data, nil
}
