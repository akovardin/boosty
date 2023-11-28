package boosty

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (b *Boosty) Subscriptions(offset, limit int) ([]Subscription, error) {
	u := fmt.Sprintf("/v1/blog/%s/subscription_level/?show_free_level=true&sort_by=on_time&offset=%d&limit=%d&order=gt", b.blog, offset, limit)
	resp, err := b.request.Request(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("error on do request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("boosty subscriptions status error")
	}

	res := Subscriptions{
		Data: []Subscription{},
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("boosty subscriptions decode error: %w", err)
	}

	return res.Data, nil
}
