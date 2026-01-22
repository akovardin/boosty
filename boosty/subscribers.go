package boosty

import (
	"fmt"
	"net/http"
	"net/url"
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

func (b *Boosty) Subscribers(values url.Values) (*Subscribers, error) {
	u := fmt.Sprintf("/v1/blog/%s/subscribers?%s", b.blog, values.Encode())
	//u := fmt.Sprintf("/v1/blog/%s/subscribers?sort_by=on_time&offset=%d&limit=%d&order=gt", b.blog, offset, limit)

	m := Method[Subscribers]{
		request: b.request,
		method:  http.MethodGet,
		url:     u,
		values:  url.Values{},
	}

	return m.Call(Subscribers{})
}
