package request

import (
	"net/http"

	"kovardin.ru/projects/boosty/auth"
)

type Option func(b *Request) error

func WithClient(client *http.Client) Option {
	return func(r *Request) error {
		r.client = client

		return nil
	}
}

func WithAuth(auth *auth.Auth) Option {
	return func(r *Request) error {
		r.auth = auth

		return nil
	}
}

func WithUrl(url string) Option {
	return func(r *Request) error {
		r.url = url

		return nil
	}
}
