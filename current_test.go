package boosty

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"gohome.4gophers.ru/getapp/boosty/auth"
	"gohome.4gophers.ru/getapp/boosty/request"
)

type CurrentTestSuite struct {
	suite.Suite
}

func (s *CurrentTestSuite) SetupTest() {
	//
}

func (s *BoostyTestSuite) TestStats() {
	tests := map[string]struct {
		followersCount int
		paidCount      int
		body           string
		token          string
	}{
		"success stats": {
			followersCount: 0, paidCount: 0, body: currentBody, token: "123",
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

			stats, err := b.Current()

			s.Assert().NoError(err)
			s.Assert().Equal(test.followersCount, stats.FollowersCount)
			s.Assert().Equal(test.paidCount, stats.FollowersCount)

		})
	}
}

const currentBody = `
{
  "followersCount": 0,
  "income": 6210,
  "balance": 0,
  "payoutSum": 6210,
  "paidCount": 2,
  "hold": 0
}
`
