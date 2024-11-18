package user

type UserResponse struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Created string  `json:"created"`
	Updated string  `json:"updated,omitempty"`
	Deleted *string `json:"deleted,omitempty"`
}
