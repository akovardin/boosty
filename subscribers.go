package boosty

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Subscriber struct {
	HasAvatar bool `json:"hasAvatar"`
	Payments  int  `json:"payments"`
	Level     struct {
		Deleted bool   `json:"deleted"`
		Name    string `json:"name"`
		Price   int    `json:"price"`
		OwnerID int    `json:"ownerId"`
		Data    []struct {
			Type        string `json:"type"`
			Content     string `json:"content"`
			Modificator string `json:"modificator"`
		} `json:"data"`
		ID             int `json:"id"`
		CurrencyPrices struct {
			RUB int     `json:"RUB"`
			USD float64 `json:"USD"`
		} `json:"currencyPrices"`
		CreatedAt  int  `json:"createdAt"`
		IsArchived bool `json:"isArchived"`
	} `json:"level"`
	Email         string `json:"email"`
	IsBlackListed bool   `json:"isBlackListed"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	OnTime        int    `json:"onTime"`
	Subscribed    bool   `json:"subscribed"`
	NextPayTime   int    `json:"nextPayTime"`
	Price         int    `json:"price"`
	AvatarURL     string `json:"avatarUrl"`
}

type Subscribers struct {
	Offset int          `json:"offset"`
	Total  int          `json:"total"`
	Limit  int          `json:"limit"`
	Data   []Subscriber `json:"data"`
}

func (b *Boosty) Subscribers(offset, limit int) ([]Subscriber, error) {
	u := fmt.Sprintf("/v1/blog/%s/subscribers?sort_by=on_time&offset=%d&limit=%d&order=gt", b.blog, offset, limit)
	resp, err := b.request.Request(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("error on do request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("boosty subscribers status error")
	}

	res := Subscribers{
		Data: []Subscriber{},
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("boosty subscribers decode error: %w", err)
	}

	return res.Data, nil
}
