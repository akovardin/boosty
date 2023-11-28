package boosty

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"kovardin.ru/projects/boosty/auth"
	"kovardin.ru/projects/boosty/request"
)

type TargetsTestSuite struct {
	suite.Suite
}

func (s *TargetsTestSuite) SetupTest() {
	//
}

func (s *TargetsTestSuite) TestStats() {
	tests := map[string]struct {
		targetsCount int
		bloggerId    int
		targetSum    int
		body         string
		token        string
	}{
		"success stats": {
			targetsCount: 1, bloggerId: 10435460, targetSum: 1000, body: targetsBody, token: "123",
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

			stats, err := b.Targets(true)

			s.Assert().NoError(err)
			s.Assert().Equal(test.targetsCount, len(stats.Data))
			s.Assert().Equal(test.bloggerId, stats.Data[0].BloggerID)
			s.Assert().Equal(test.targetSum, stats.Data[0].TargetSum)
		})
	}
}

const targetsBody = `
{
  "data": [
    {
      "currentSum": 4,
      "bloggerId": 10435460,
      "bloggerUrl": "getapp",
      "id": 242127,
      "type": "subscribers",
      "createdAt": 1665345059,
      "targetSum": 1000,
      "description": "С 1 000 подписчиков будет значительно проще заниматься проектом. Появится возможность сделать демо версию и оказывать поддержку.",
      "priority": 0,
      "finishTime": null
    }
  ]
}
`
