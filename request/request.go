package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gohome.4gophers.ru/getapp/boosty/auth"
)

type Request struct {
	url    string
	client *http.Client
	auth   *auth.Auth
}

func New(options ...Option) (*Request, error) {
	auth, err := auth.New()
	if err != nil {
		return nil, err
	}

	r := &Request{
		url:    "https://api.boosty.to",
		client: &http.Client{},
		auth:   auth,
	}

	for _, o := range options {
		if err := o(r); err != nil {
			return nil, err
		}
	}

	return r, nil
}

// check auth code and re request
func (b *Request) Request(method string, u string, body io.Reader) (*http.Response, error) {
	resp, err := b.do(method, b.url+u, body)
	if err != nil {
		return nil, fmt.Errorf("error on do request: %w", err)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		if err := b.refresh(); err != nil {
			return nil, fmt.Errorf("error on refresh token: %w", err)
		}

		resp, err = b.do(method, u, body)
		if err != nil {
			return nil, fmt.Errorf("error on do request: %w", err)
		}
	}

	return resp, nil
}

// create request with auth token
func (b *Request) do(method string, u string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, fmt.Errorf("boosty stats request error: %w", err)
	}

	req.Header.Add("Authorization", b.auth.BearerHeader())

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("boosty stats do error: %w", err)
	}

	return resp, nil
}

type Refresh struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func (b *Request) refresh() error {
	if b.auth.RefreshToken() == "" {
		return errors.New("empty refresh token")
	}

	form := url.Values{}
	form.Add("device_id", b.auth.DeviceId())
	form.Add("device_os", "web")
	form.Add("grant_type", "refresh_token")
	form.Add("refresh_token", b.auth.RefreshToken())

	req, err := http.NewRequest(http.MethodPost, b.url+"/oauth/token/", strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("boosty refresh request error: %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	req.Header.Set("Authorization", b.auth.BearerHeader())
	resp, err := b.client.Do(req)
	if err != nil {
		return fmt.Errorf("boosty refresh do error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("boosty refresh do error: %d", resp.StatusCode)
	}

	refresh := Refresh{}

	if err := json.NewDecoder(resp.Body).Decode(&refresh); err != nil {
		return fmt.Errorf("boosty refresh decode error: %w", err)
	}

	info := auth.Info{
		AccessToken:  refresh.AccessToken,
		RefreshToken: refresh.RefreshToken,
		ExpiresAt:    time.Now().Unix() + refresh.ExpiresIn,
		DeviceId:     b.auth.DeviceId(),
	}

	b.auth.Update(info)
	if err := b.auth.Save(); err != nil {
		return fmt.Errorf("boosty refresh save error: %w", err)
	}

	return nil
}
