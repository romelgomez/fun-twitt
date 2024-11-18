package follow

import validation "github.com/go-ozzo/ozzo-validation/v4"

type FollowCreate struct {
	// FollowerID: The user who follows someone.
	FollowerID string `json:"follower_id"`
	// FolloweeID: The user being followed.
	FolloweeID string `json:"followee_id"`
}

func (f *FollowCreate) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.FollowerID, validation.Required),
		validation.Field(&f.FolloweeID, validation.Required),
	)
}
