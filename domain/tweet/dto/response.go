package tweet

type TweetResponse struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	UserID    string `json:"user_id"`
	SortID    string `json:"sort_id"`
	CreatedAt string `json:"created_at"`
}