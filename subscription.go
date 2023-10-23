package boosty

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
