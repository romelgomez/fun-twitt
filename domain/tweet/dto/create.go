package tweet

import validation "github.com/go-ozzo/ozzo-validation/v4"

type TweetCreate struct {
	Content string `json:"content"`
	UserID  string `json:"user_id"`
}

func (t *TweetCreate) Validate() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Content, validation.Required, validation.Length(1, 280)),
		validation.Field(&t.UserID, validation.Required),
	)
}
