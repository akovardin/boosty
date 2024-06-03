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

type StatsTestSuite struct {
	suite.Suite
}

func (s *StatsTestSuite) SetupTest() {
	//
}

func (s *StatsTestSuite) TestStats() {
	tests := map[string]struct {
		postSaleMoney int
		decFollowers  int
		donations     int
		body          string
		token         string
	}{
		"success stats": {
			postSaleMoney: 5, decFollowers: 5, donations: 5, body: statsBody, token: "123",
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

			stats, err := b.Stats(url.Values{})

			s.Assert().NoError(err)
			s.Assert().Equal(test.postSaleMoney, len(stats.PostSaleMoney))
			s.Assert().Equal(test.decFollowers, len(stats.DecFollowers))
			s.Assert().Equal(test.donations, len(stats.Donations))
		})
	}
}

const statsBody = `
{
  "postSaleMoney": [
    {
      "day": 20,
      "year": 2023,
      "count": 0,
      "month": 8
    },
    {
      "day": 5,
      "year": 2023,
      "month": 9,
      "count": 0
    },
    {
      "day": 6,
      "year": 2023,
      "count": 0,
      "month": 9
    },
    {
      "month": 9,
      "count": 0,
      "year": 2023,
      "day": 7
    },
    {
      "year": 2023,
      "day": 18,
      "month": 9,
      "count": 0
    }
  ],
  "upSubscribers": [
    {
      "year": 2023,
      "day": 20,
      "count": 0,
      "month": 8
    },
    {
      "month": 9,
      "count": 0,
      "day": 5,
      "year": 2023
    },
    {
      "count": 0,
      "month": 9,
      "year": 2023,
      "day": 6
    },
    {
      "count": 0,
      "month": 9,
      "day": 7,
      "year": 2023
    },
    {
      "year": 2023,
      "day": 18,
      "count": 0,
      "month": 9
    }
  ],
  "messagesSale": [
    {
      "count": 0,
      "month": 8,
      "year": 2023,
      "day": 20
    },
    {
      "month": 9,
      "count": 0,
      "year": 2023,
      "day": 5
    },
    {
      "count": 0,
      "month": 9,
      "day": 6,
      "year": 2023
    },
    {
      "month": 9,
      "count": 0,
      "year": 2023,
      "day": 7
    },
    {
      "day": 18,
      "year": 2023,
      "count": 0,
      "month": 9
    }
  ],
  "decSubscribers": [
    {
      "count": 0,
      "month": 8,
      "year": 2023,
      "day": 20
    },
    {
      "year": 2023,
      "day": 5,
      "count": 0,
      "month": 9
    },
    {
      "month": 9,
      "count": 0,
      "day": 6,
      "year": 2023
    },
    {
      "count": 0,
      "month": 9,
      "day": 7,
      "year": 2023
    },
    {
      "month": 9,
      "count": 0,
      "day": 18,
      "year": 2023
    }
  ],
  "postsSale": [
    {
      "day": 20,
      "year": 2023,
      "month": 8,
      "count": 0
    },
    {
      "count": 0,
      "month": 9,
      "year": 2023,
      "day": 5
    },
    {
      "day": 6,
      "year": 2023,
      "month": 9,
      "count": 0
    },
    {
      "month": 9,
      "count": 0,
      "year": 2023,
      "day": 7
    },
    {
      "day": 18,
      "year": 2023,
      "month": 9,
      "count": 0
    }
  ],
  "donationsMoney": [
    {
      "count": 0,
      "month": 8,
      "year": 2023,
      "day": 20
    },
    {
      "year": 2023,
      "day": 5,
      "month": 9,
      "count": 0
    },
    {
      "month": 9,
      "count": 0,
      "year": 2023,
      "day": 6
    },
    {
      "day": 7,
      "year": 2023,
      "count": 0,
      "month": 9
    },
    {
      "month": 9,
      "count": 0,
      "day": 18,
      "year": 2023
    }
  ],
  "giftsSaleSaleMoney": [
    {
      "month": 8,
      "count": 0,
      "day": 20,
      "year": 2023
    },
    {
      "month": 9,
      "count": 0,
      "day": 5,
      "year": 2023
    },
    {
      "year": 2023,
      "day": 6,
      "month": 9,
      "count": 0
    },
    {
      "month": 9,
      "count": 0,
      "day": 7,
      "year": 2023
    },
    {
      "day": 18,
      "year": 2023,
      "count": 0,
      "month": 9
    }
  ],
  "messagesSaleMoney": [
    {
      "month": 8,
      "count": 0,
      "year": 2023,
      "day": 20
    },
    {
      "year": 2023,
      "day": 5,
      "count": 0,
      "month": 9
    },
    {
      "year": 2023,
      "day": 6,
      "count": 0,
      "month": 9
    },
    {
      "day": 7,
      "year": 2023,
      "count": 0,
      "month": 9
    },
    {
      "day": 18,
      "year": 2023,
      "month": 9,
      "count": 0
    }
  ],
  "totalMoney": [
    {
      "day": 20,
      "year": 2023,
      "count": 0,
      "month": 8
    },
    {
      "month": 9,
      "count": 0,
      "year": 2023,
      "day": 5
    },
    {
      "year": 2023,
      "day": 6,
      "count": 300,
      "month": 9
    },
    {
      "year": 2023,
      "day": 7,
      "month": 9,
      "count": 0
    },
    {
      "count": 0,
      "month": 9,
      "day": 18,
      "year": 2023
    }
  ],
  "decFollowers": [
    {
      "count": 0,
      "month": 8,
      "year": 2023,
      "day": 20
    },
    {
      "day": 5,
      "year": 2023,
      "count": 0,
      "month": 9
    },
    {
      "month": 9,
      "count": 0,
      "day": 6,
      "year": 2023
    },
    {
      "year": 2023,
      "day": 7,
      "month": 9,
      "count": 0
    },
    {
      "count": 0,
      "month": 9,
      "year": 2023,
      "day": 18
    }
  ],
  "incSubscribersMoney": [
    {
      "year": 2023,
      "day": 20,
      "month": 8,
      "count": 0
    },
    {
      "count": 0,
      "month": 9,
      "day": 5,
      "year": 2023
    },
    {
      "day": 6,
      "year": 2023,
      "month": 9,
      "count": 0
    },
    {
      "month": 9,
      "count": 0,
      "year": 2023,
      "day": 7
    },
    {
      "year": 2023,
      "day": 18,
      "count": 0,
      "month": 9
    }
  ],
  "recurrentsMoney": [
    {
      "year": 2023,
      "day": 20,
      "month": 8,
      "count": 0
    },
    {
      "day": 5,
      "year": 2023,
      "count": 0,
      "month": 9
    },
    {
      "count": 300,
      "month": 9,
      "day": 6,
      "year": 2023
    },
    {
      "year": 2023,
      "day": 7,
      "count": 0,
      "month": 9
    },
    {
      "year": 2023,
      "day": 18,
      "count": 0,
      "month": 9
    }
  ],
  "recurrents": [
    {
      "month": 8,
      "count": 0,
      "day": 20,
      "year": 2023
    },
    {
      "year": 2023,
      "day": 5,
      "count": 0,
      "month": 9
    },
    {
      "day": 6,
      "year": 2023,
      "count": 1,
      "month": 9
    },
    {
      "year": 2023,
      "day": 7,
      "count": 0,
      "month": 9
    },
    {
      "month": 9,
      "count": 0,
      "day": 18,
      "year": 2023
    }
  ],
  "referalMoney": [
    {
      "year": 2023,
      "day": 20,
      "month": 8,
      "count": 0
    },
    {
      "year": 2023,
      "day": 5,
      "count": 0,
      "month": 9
    },
    {
      "month": 9,
      "count": 0,
      "year": 2023,
      "day": 6
    },
    {
      "year": 2023,
      "day": 7,
      "month": 9,
      "count": 0
    },
    {
      "count": 0,
      "month": 9,
      "day": 18,
      "year": 2023
    }
  ],
  "referalMoneyOut": [
    {
      "day": 20,
      "year": 2023,
      "count": 0,
      "month": 8
    },
    {
      "day": 5,
      "year": 2023,
      "count": 0,
      "month": 9
    },
    {
      "month": 9,
      "count": 0,
      "day": 6,
      "year": 2023
    },
    {
      "year": 2023,
      "day": 7,
      "count": 0,
      "month": 9
    },
    {
      "count": 0,
      "month": 9,
      "year": 2023,
      "day": 18
    }
  ],
  "incFollowers": [
    {
      "count": 0,
      "month": 8,
      "day": 20,
      "year": 2023
    },
    {
      "day": 5,
      "year": 2023,
      "month": 9,
      "count": 0
    },
    {
      "day": 6,
      "year": 2023,
      "month": 9,
      "count": 0
    },
    {
      "month": 9,
      "count": 0,
      "year": 2023,
      "day": 7
    },
    {
      "month": 9,
      "count": 0,
      "year": 2023,
      "day": 18
    }
  ],
  "referal": [
    {
      "count": 0,
      "month": 8,
      "year": 2023,
      "day": 20
    },
    {
      "year": 2023,
      "day": 5,
      "month": 9,
      "count": 0
    },
    {
      "day": 6,
      "year": 2023,
      "count": 0,
      "month": 9
    },
    {
      "year": 2023,
      "day": 7,
      "month": 9,
      "count": 0
    },
    {
      "year": 2023,
      "day": 18,
      "count": 0,
      "month": 9
    }
  ],
  "donations": [
    {
      "day": 20,
      "year": 2023,
      "month": 8,
      "count": 0
    },
    {
      "month": 9,
      "count": 0,
      "year": 2023,
      "day": 5
    },
    {
      "year": 2023,
      "day": 6,
      "month": 9,
      "count": 0
    },
    {
      "day": 7,
      "year": 2023,
      "month": 9,
      "count": 0
    },
    {
      "year": 2023,
      "day": 18,
      "count": 0,
      "month": 9
    }
  ],
  "incSubscribers": [
    {
      "year": 2023,
      "day": 20,
      "count": 0,
      "month": 8
    },
    {
      "day": 5,
      "year": 2023,
      "count": 0,
      "month": 9
    },
    {
      "year": 2023,
      "day": 6,
      "count": 0,
      "month": 9
    },
    {
      "day": 7,
      "year": 2023,
      "month": 9,
      "count": 0
    },
    {
      "year": 2023,
      "day": 18,
      "count": 0,
      "month": 9
    }
  ],
  "giftsSale": [
    {
      "day": 20,
      "year": 2023,
      "month": 8,
      "count": 0
    },
    {
      "day": 5,
      "year": 2023,
      "month": 9,
      "count": 0
    },
    {
      "count": 0,
      "month": 9,
      "day": 6,
      "year": 2023
    },
    {
      "count": 0,
      "month": 9,
      "day": 7,
      "year": 2023
    },
    {
      "month": 9,
      "count": 0,
      "day": 18,
      "year": 2023
    }
  ],
  "upSubscribersMoney": [
    {
      "year": 2023,
      "day": 20,
      "month": 8,
      "count": 0
    },
    {
      "day": 5,
      "year": 2023,
      "count": 0,
      "month": 9
    },
    {
      "month": 9,
      "count": 0,
      "year": 2023,
      "day": 6
    },
    {
      "year": 2023,
      "day": 7,
      "count": 0,
      "month": 9
    },
    {
      "day": 18,
      "year": 2023,
      "count": 0,
      "month": 9
    }
  ],
  "holds": [
    {
      "year": 2023,
      "day": 20,
      "month": 8,
      "count": 0
    },
    {
      "year": 2023,
      "day": 5,
      "month": 9,
      "count": 0
    },
    {
      "month": 9,
      "count": 0,
      "day": 6,
      "year": 2023
    },
    {
      "count": 0,
      "month": 9,
      "year": 2023,
      "day": 7
    },
    {
      "count": 0,
      "month": 9,
      "year": 2023,
      "day": 18
    }
  ]
}
`
