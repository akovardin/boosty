package boosty

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
