package boosty

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type BoostyTestSuite struct {
	suite.Suite
}

func (s *BoostyTestSuite) SetupTest() {
	//
}

func (s *BoostyTestSuite) TestSubscribers() {
	tests := map[string]struct {
		count int
		body  string
		name  string
		token string
	}{
		"success count 2": {
			count: 2, body: subscribersBody, name: "getapp.store", token: "123",
		},
	}

	for name, test := range tests {
		s.T().Run(name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				auth := r.Header.Get("Authorization")

				s.Assert().Equal(auth, "Bearer "+test.token)

				fmt.Fprintf(w, test.body)
			}))
			defer svr.Close()

			b := Boosty{
				token:  test.token,
				base:   svr.URL,
				client: &http.Client{},
			}

			ss, err := b.Subscribers(0, 10)

			s.Assert().NoError(err)
			s.Assert().Equal(test.count, len(ss))
			if len(ss) > 0 {
				s.Assert().Equal(test.name, ss[0].Name)
			}

		})
	}
}

func (s *BoostyTestSuite) TestSubscriptions() {
	tests := map[string]struct {
		count int
		body  string
		name  string
		token string
	}{
		"success count 3": {
			count: 3, body: subscriptionsBody, name: "Follower", token: "123",
		},
	}

	for name, test := range tests {
		s.T().Run(name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				auth := r.Header.Get("Authorization")

				s.Assert().Equal(auth, "Bearer "+test.token)

				fmt.Fprintf(w, test.body)
			}))
			defer svr.Close()

			b := Boosty{
				token:  test.token,
				base:   svr.URL,
				client: &http.Client{},
			}

			ss, err := b.Subscriptions(0, 10)

			s.Assert().NoError(err)
			s.Assert().Equal(test.count, len(ss))
			if len(ss) > 0 {
				s.Assert().Equal(test.name, ss[0].Name)
			}

		})
	}
}

func (s *BoostyTestSuite) TestStats() {
	tests := map[string]struct {
		followersCount int
		paidCount      int
		body           string
		token          string
	}{
		"success stats": {
			followersCount: 0, paidCount: 0, body: statsBody, token: "123",
		},
	}

	for name, test := range tests {
		s.T().Run(name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				auth := r.Header.Get("Authorization")

				s.Assert().Equal(auth, "Bearer "+test.token)

				fmt.Fprintf(w, test.body)
			}))
			defer svr.Close()

			b := Boosty{
				token:  test.token,
				base:   svr.URL,
				client: &http.Client{},
			}

			stats, err := b.Stats()

			s.Assert().NoError(err)
			s.Assert().Equal(test.followersCount, stats.FollowersCount)
			s.Assert().Equal(test.paidCount, stats.FollowersCount)

		})
	}
}

func TestBoostyTestSuite(t *testing.T) {
	suite.Run(t, new(BoostyTestSuite))
}

const subscribersBody = `
{
  "limit": 2,
  "data": [
    {
      "subscribed": true,
      "price": 300,
      "payments": 2970,
      "hasAvatar": true,
      "nextPayTime": 1699209830,
      "id": 3684586,
      "level": {
        "ownerId": 10435460,
        "price": 300,
        "data": [
          {
            "type": "text",
            "modificator": "",
            "content": "[\"Узнавай про фикс багов самый первый\",\"unstyled\",[]]"
          },
          {
            "type": "text",
            "modificator": "BLOCK_END",
            "content": ""
          },
          {
            "type": "text",
            "modificator": "BLOCK_END",
            "content": ""
          }
        ],
        "isArchived": false,
        "name": "Стандартная подписка",
        "createdAt": 1664319534,
        "id": 1091773,
        "currencyPrices": {
          "RUB": 300,
          "USD": 3.2
        },
        "deleted": false
      },
      "name": "getapp.store",
      "avatarUrl": "https://images.boosty.to/user/3684586/avatar?change_time=1664275665",
      "email": "artem.kovardin@gmail.com",
      "onTime": 1670697830,
      "isBlackListed": false
    },
    {
      "price": 300,
      "subscribed": true,
      "id": 11222871,
      "nextPayTime": 1697748053,
      "hasAvatar": true,
      "payments": 3240,
      "email": "",
      "avatarUrl": "https://images.boosty.to/user/11222871/avatar?change_time=1666643907",
      "name": "Lena Nesterenko",
      "level": {
        "data": [
          {
            "type": "text",
            "modificator": "",
            "content": "[\"Узнавай про фикс багов самый первый\",\"unstyled\",[]]"
          },
          {
            "type": "text",
            "modificator": "BLOCK_END",
            "content": ""
          },
          {
            "modificator": "BLOCK_END",
            "type": "text",
            "content": ""
          }
        ],
        "isArchived": false,
        "createdAt": 1664319534,
        "name": "Стандартная подписка",
        "ownerId": 10435460,
        "price": 300,
        "currencyPrices": {
          "RUB": 300,
          "USD": 3.2
        },
        "deleted": false,
        "id": 1091773
      },
      "isBlackListed": false,
      "onTime": 1666644053
    }
  ],
  "offset": 2,
  "total": 2
}
`

const subscriptionsBody = `
{
  "currentId": null,
  "subscriptions": [],
  "nextId": null,
  "data": [
    {
      "currencyPrices": {
        "USD": 0,
        "RUB": 0
      },
      "promos": [],
      "externalApps": {
        "telegram": {
          "groups": [],
          "isConfigured": false
        },
        "discord": {
          "isConfigured": false,
          "data": {
            "role": {
              "name": "",
              "id": ""
            }
          }
        }
      },
      "id": 1091770,
      "price": 0,
      "ownerId": 10435460,
      "isArchived": false,
      "createdAt": 1664319306,
      "name": "Follower",
      "badges": [],
      "data": [],
      "deleted": false
    },
    {
      "promos": [],
      "currencyPrices": {
        "RUB": 300,
        "USD": 3.2
      },
      "ownerId": 10435460,
      "isArchived": false,
      "externalApps": {
        "telegram": {
          "isConfigured": false,
          "groups": []
        },
        "discord": {
          "isConfigured": false,
          "data": {
            "role": {
              "name": "",
              "id": ""
            }
          }
        }
      },
      "price": 300,
      "id": 1091773,
      "createdAt": 1664319534,
      "deleted": false,
      "data": [
        {
          "modificator": "",
          "type": "text",
          "content": "[\"Узнавай про фикс багов самый первый\",\"unstyled\",[]]"
        },
        {
          "modificator": "BLOCK_END",
          "type": "text",
          "content": ""
        },
        {
          "content": "",
          "type": "text",
          "modificator": "BLOCK_END"
        }
      ],
      "name": "Стандартная подписка",
      "badges": []
    },
    {
      "isArchived": false,
      "ownerId": 10435460,
      "price": 700,
      "id": 1122632,
      "externalApps": {
        "discord": {
          "isConfigured": false,
          "data": {
            "role": {
              "id": "",
              "name": ""
            }
          }
        },
        "telegram": {
          "groups": [],
          "isConfigured": false
        }
      },
      "promos": [],
      "currencyPrices": {
        "USD": 7.5,
        "RUB": 700
      },
      "deleted": false,
      "data": [
        {
          "modificator": "",
          "type": "text",
          "content": "[\"Заказывай новые фичи\",\"unstyled\",[]]"
        },
        {
          "content": "",
          "type": "text",
          "modificator": "BLOCK_END"
        }
      ],
      "badges": [],
      "name": "Расширенная подписка",
      "createdAt": 1665493091
    }
  ]
}
`

const statsBody = `
{
  "followersCount": 0,
  "income": 6210,
  "balance": 0,
  "payoutSum": 6210,
  "paidCount": 2,
  "hold": 0
}
`
