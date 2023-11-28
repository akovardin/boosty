package boosty

import "kovardin.ru/projects/boosty/request"

type Option func(b *Boosty) error

func WithRequest(request *request.Request) Option {
	return func(b *Boosty) error {
		b.request = request

		return nil
	}
}
