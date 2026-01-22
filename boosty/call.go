package boosty

import (
	"encoding/json"
	"fmt"
	"net/url"

	"gohome.4gophers.ru/getapp/boosty/request"
)

type Method[T interface{}] struct {
	request *request.Request
	method  string
	url     string
	values  url.Values
}

func (m *Method[T]) Call(model T) (*T, error) {
	u := m.url + m.values.Encode()

	resp, err := m.request.Request(m.method, u, nil)
	if err != nil {
		return nil, fmt.Errorf("error on do request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("boosty request status error")
	}

	if err := json.NewDecoder(resp.Body).Decode(&model); err != nil {
		return nil, fmt.Errorf("boosty request decode error: %w", err)
	}

	return &model, nil
}
