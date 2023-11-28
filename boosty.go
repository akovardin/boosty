package boosty

import (
	"kovardin.ru/projects/boosty/request"
)

type Boosty struct {
	blog    string
	request *request.Request
}

func New(blog string, options ...Option) (*Boosty, error) {
	request, err := request.New()
	if err != nil {
		return nil, err
	}

	b := &Boosty{
		blog:    blog,
		request: request,
	}

	for _, o := range options {
		if err := o(b); err != nil {
			return nil, err
		}
	}

	return b, nil
}
