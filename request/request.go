package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"kovardin.ru/projects/boosty/auth"
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

	req.Header.Add("Authorization", b.auth.Bearer())

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("boosty stats do error: %w", err)
	}

	return resp, nil
}

func (b *Request) refresh() error {
	if b.auth.Refresh() == "" {
		return errors.New("empty refresh token")
	}

	// TODO: change device id
	body := `
{
	"device_id": "91cd83c5-7130-4a85-8ef2-66898d4ced5b;",
	"device_os": "web",
	"grant_type": "refresh_token",
	"refresh_token": "` + b.auth.Refresh() + `",
}
`

	req, err := http.NewRequest(http.MethodPost, b.url+"/oauth/token/", bytes.NewReader([]byte(body)))
	if err != nil {
		return fmt.Errorf("boosty refresh request error: %w", err)
	}

	resp, err := b.client.Do(req)
	if err != nil {
		return fmt.Errorf("boosty refresh do error: %w", err)
	}

	info := auth.Info{}

	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return fmt.Errorf("boosty refresh decode error: %w", err)
	}

	b.auth.Update(info)
	if err := b.auth.Save(); err != nil {
		return fmt.Errorf("boosty refresh save error: %w", err)
	}

	return nil
}
