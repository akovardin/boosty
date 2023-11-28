package boosty

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (b *Boosty) Targets(deleted bool) (*Targets, error) {
	u := fmt.Sprintf("/v1/target/%s/", b.blog)
	if deleted {
		u = u + "?show_deleted=true"
	}

	resp, err := b.request.Request(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("error on do request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("boosty stats status error")
	}

	res := &Targets{
		Data: []Target{},
	}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, fmt.Errorf("boosty stats decode error: %w", err)
	}

	return res, nil
}
