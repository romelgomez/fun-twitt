package user

type UserResponse struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	SortID  string  `json:"sort_id"`
	Email   string  `json:"email"`
	Created string  `json:"created"`
	Updated string  `json:"updated,omitempty"`
	Deleted *string `json:"deleted,omitempty"`
}
