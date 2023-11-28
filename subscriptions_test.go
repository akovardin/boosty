package boosty

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
	"kovardin.ru/projects/boosty/auth"
	"kovardin.ru/projects/boosty/request"
)

type SubscriptionsTestSuite struct {
	suite.Suite
}

func (s *SubscriptionsTestSuite) SetupTest() {
	//
}

func (s *SubscriptionsTestSuite) TestSubscriptions() {
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

			auth, err := auth.New(auth.WithInfo(auth.Info{
				AccessToken: test.token,
			}))
			s.Assert().NoError(err)

			req, err := request.New(
				request.WithUrl(svr.URL),
				request.WithAuth(auth),
				request.WithClient(&http.Client{}),
			)
			s.Assert().NoError(err)

			b, err := New("", WithRequest(req))
			s.Assert().NoError(err)

			v := url.Values{}
			v.Add("show_free_level", "true")
			v.Add("sort_by", "on_time")
			v.Add("offset", "0")
			v.Add("limit", "10")
			v.Add("order", "gt")

			ss, err := b.Subscriptions(v)

			s.Assert().NoError(err)
			s.Assert().Equal(test.count, len(ss.Data))
			if len(ss.Data) > 0 {
				s.Assert().Equal(test.name, ss.Data[0].Name)
			}
		})
	}
}

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
