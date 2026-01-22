package boosty

import (
	"fmt"
	"net/http"
	"net/url"
)

type Targets struct {
	Data []Target `json:"data"`
}

type Target struct {
	CreatedAt   int         `json:"createdAt"`
	Type        string      `json:"type"`
	Priority    int         `json:"priority"`
	TargetSum   int         `json:"targetSum"`
	Description string      `json:"description"`
	FinishTime  interface{} `json:"finishTime"`
	BloggerID   int         `json:"bloggerId"`
	CurrentSum  int         `json:"currentSum"`
	ID          int         `json:"id"`
	BloggerURL  string      `json:"bloggerUrl"`
}

func (b *Boosty) Targets(values url.Values) (*Targets, error) {
	u := fmt.Sprintf("/v1/target/%s/?%s", b.blog, values.Encode())

	m := Method[Targets]{
		request: b.request,
		method:  http.MethodGet,
		url:     u,
		values:  values,
	}

	return m.Call(Targets{})
}
