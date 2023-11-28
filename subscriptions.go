package boosty

import (
	"fmt"
	"net/http"
	"net/url"
)

type Subscription struct {
	IsArchived bool          `json:"isArchived"`
	Deleted    bool          `json:"deleted"`
	Badges     []interface{} `json:"badges"`
	ID         int           `json:"id"`
	Data       []struct {
		Modificator string `json:"modificator"`
		Type        string `json:"type"`
		Content     string `json:"content"`
	} `json:"data"`
	Name         string `json:"name"`
	ExternalApps struct {
		Telegram struct {
			Groups       []interface{} `json:"groups"`
			IsConfigured bool          `json:"isConfigured"`
		} `json:"telegram"`
		Discord struct {
			Data struct {
				Role struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"role"`
			} `json:"data"`
			IsConfigured bool `json:"isConfigured"`
		} `json:"discord"`
	} `json:"externalApps"`
	Price          int           `json:"price"`
	Promos         []interface{} `json:"promos"`
	OwnerID        int           `json:"ownerId"`
	CreatedAt      int           `json:"createdAt"`
	CurrencyPrices struct {
		USD float64 `json:"USD"`
		RUB int     `json:"RUB"`
	} `json:"currencyPrices"`
}

type Subscriptions struct {
	Offset int            `json:"offset"`
	Total  int            `json:"total"`
	Limit  int            `json:"limit"`
	Data   []Subscription `json:"data"`
}

func (b *Boosty) Subscriptions(values url.Values) (*Subscriptions, error) {
	u := fmt.Sprintf("/v1/blog/%s/subscription_level/?%s", b.blog, values.Encode())
	//u := fmt.Sprintf("/v1/blog/%s/subscription_level/?show_free_level=true&sort_by=on_time&offset=%d&limit=%d&order=gt", b.blog, offset, limit)

	m := Method[Subscriptions]{
		request: b.request,
		method:  http.MethodGet,
		url:     u,
		values:  url.Values{},
	}

	return m.Call(Subscriptions{})
}
