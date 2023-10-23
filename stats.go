package boosty

type Stats struct {
	PaidCount      int `json:"paidCount"`
	FollowersCount int `json:"followersCount"`
	Hold           int `json:"hold"`
	Income         int `json:"income"`
	Balance        int `json:"balance"`
	PayoutSum      int `json:"payoutSum"`
}
