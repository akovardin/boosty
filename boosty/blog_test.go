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

type BlogTestSuite struct {
	suite.Suite
}

func (s *BlogTestSuite) SetupTest() {}

func (s *BlogTestSuite) TestBlog() {
	tests := map[string]struct {
		blogName string
		body     string
		token    string
	}{
		"success": {
			blogName: "test-blog",
			body:     blogBody,
			token:    "123",
		},
	}

	for name, test := range tests {
		s.T().Run(name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				a := r.Header.Get("Authorization")
				s.Assert().Equal(a, "Bearer "+test.token)

				s.Assert().Equal(fmt.Sprintf("/v1/blog/%s", test.blogName), r.URL.Path)

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

			b, err := New(test.blogName, WithRequest(req))
			s.Assert().NoError(err)

			blog, err := b.Blog()
			s.Assert().NoError(err)
			s.Assert().Equal(test.blogName, blog.Owner.Name)
		})
	}
}

const blogBody = `
{
  "owner": {
    "id": 12345,
    "hasAvatar": true,
    "name": "test-blog",
    "avatarUrl": "https://example.com/avatar.jpg"
  },
  "title": "Test Blog",
  "isReadOnly": false,
  "flags": {
    "showPostDonations": true,
    "allowGoogleIndex": true,
    "hasTargets": true,
    "acceptDonationMessages": true,
    "allowIndex": true,
    "isRssFeedEnabled": true,
    "hasSubscriptionLevels": true
  },
  "signedQuery": "test-query",
  "isBlackListedByUser": false,
  "isSubscribed": true,
  "subscription": null,
  "isTotalBaned": false,
  "accessRights": {
    "canSetPayout": true,
    "canCreateComments": true,
    "canEdit": true,
    "canView": true,
    "canDeleteComments": true,
    "canCreate": true
  },
  "count": {
    "subscribers": 100,
    "posts": 50
  },
  "blogUrl": "https://boosty.to/test-blog",
  "isOwner": true,
  "publicWebSocketChannel": "test-channel",
  "subscriptionKind": "active",
  "isBlackListed": false,
  "allowedPromoTypes": ["type1", "type2"],
  "description": [
    {
      "type": "text",
      "content": "Test description"
    }
  ],
  "socialLinks": [
    {
      "url": "https://example.com",
      "type": "website"
    }
  ],
  "hasAdultContent": false,
  "coverUrl": "https://example.com/cover.jpg"
}
`

func TestBlogTestSuite(t *testing.T) {
	suite.Run(t, new(BlogTestSuite))
}
