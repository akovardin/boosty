package boosty

import (
	"fmt"
	"net/http"
	"net/url"
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

func (b *Boosty) Stats(values url.Values) (*Stats, error) {
	u := fmt.Sprintf("v1/blog/%s/stat/data/?%s", b.blog, values.Encode())

	m := Method[Stats]{
		request: b.request,
		method:  http.MethodGet,
		url:     u,
		values:  url.Values{},
	}

	return m.Call(Stats{})
}
