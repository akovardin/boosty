package boosty

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Point struct {
	Day   int `json:"day"`
	Year  int `json:"year"`
	Count int `json:"count"`
	Month int `json:"month"`
}
type Stats struct {
	PostSaleMoney       []Point `json:"postSaleMoney"`
	UpSubscribers       []Point `json:"upSubscribers"`
	MessagesSale        []Point `json:"messagesSale"`
	DecSubscribers      []Point `json:"decSubscribers"`
	PostsSale           []Point `json:"postsSale"`
	DonationsMoney      []Point `json:"donationsMoney"`
	GiftsSaleSaleMoney  []Point `json:"giftsSaleSaleMoney"`
	MessagesSaleMoney   []Point `json:"messagesSaleMoney"`
	TotalMoney          []Point `json:"totalMoney"`
	DecFollowers        []Point `json:"decFollowers"`
	IncSubscribersMoney []Point `json:"incSubscribersMoney"`
	RecurrentsMoney     []Point `json:"recurrentsMoney"`
	Recurrents          []Point `json:"recurrents"`
	ReferalMoney        []Point `json:"referalMoney"`
	ReferalMoneyOut     []Point `json:"referalMoneyOut"`
	IncFollowers        []Point `json:"incFollowers"`
	Referal             []Point `json:"referal"`
	Donations           []Point `json:"donations"`
	IncSubscribers      []Point `json:"incSubscribers"`
	GiftsSale           []Point `json:"giftsSale"`
	UpSubscribersMoney  []Point `json:"upSubscribersMoney"`
	Holds               []Point `json:"holds"`
}

func (b *Boosty) Stats() (*Stats, error) {
	u := fmt.Sprintf("/v1/target/%s/", b.blog)
	resp, err := b.request.Request(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("error on do request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("boosty stats status error")
	}

	res := &Stats{
		PostSaleMoney:       []Point{},
		UpSubscribers:       []Point{},
		MessagesSale:        []Point{},
		DecSubscribers:      []Point{},
		PostsSale:           []Point{},
		DonationsMoney:      []Point{},
		GiftsSaleSaleMoney:  []Point{},
		MessagesSaleMoney:   []Point{},
		TotalMoney:          []Point{},
		DecFollowers:        []Point{},
		IncSubscribersMoney: []Point{},
		RecurrentsMoney:     []Point{},
		Recurrents:          []Point{},
		ReferalMoney:        []Point{},
		ReferalMoneyOut:     []Point{},
		IncFollowers:        []Point{},
		Referal:             []Point{},
		Donations:           []Point{},
		IncSubscribers:      []Point{},
		GiftsSale:           []Point{},
		UpSubscribersMoney:  []Point{},
		Holds:               []Point{},
	}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, fmt.Errorf("boosty stats decode error: %w", err)
	}

	return res, nil
}
