package follow

type FollowResponse struct {
	ID         string `json:"id"`
	FollowerID string `json:"follower_id"`
	FolloweeID string `json:"followee_id"`
	CreatedAt  string `json:"created_at"`
}
