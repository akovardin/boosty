package boosty

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
	"gohome.4gophers.ru/getapp/boosty/auth"
	"gohome.4gophers.ru/getapp/boosty/request"
)

type SubscribersTestSuite struct {
	suite.Suite
}

func (s *SubscribersTestSuite) SetupTest() {
	//
}

func (s *SubscribersTestSuite) TestSubscribers() {
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
				a := r.Header.Get("Authorization")

				s.Assert().Equal(a, "Bearer "+test.token)

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
			v.Add("offset", "0")
			v.Add("limit", "10")
			v.Add("order", "gt")
			v.Add("sort_by", "on_time")

			ss, err := b.Subscribers(v)

			s.Assert().NoError(err)
			s.Assert().Equal(test.count, len(ss.Data))
			if len(ss.Data) > 0 {
				s.Assert().Equal(test.name, ss.Data[0].Name)
			}
		})
	}
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
