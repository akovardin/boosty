package boosty

import "gohome.4gophers.ru/getapp/boosty/request"

type Option func(b *Boosty) error

func WithRequest(request *request.Request) Option {
	return func(b *Boosty) error {
		b.request = request

		return nil
	}
}
