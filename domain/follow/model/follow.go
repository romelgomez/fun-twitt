package follow

import (
	dto "funtwitt/domain/follow/dto"
	"funtwitt/helper"
	"funtwitt/prisma/db"
)

type Follow struct {
	db.InnerFollow
}

func (f *Follow) Dto() dto.FollowResponse {
	return dto.FollowResponse{
		ID:         f.ID,
		FollowerID: f.FollowerID,
		FolloweeID: f.FolloweeID,
		CreatedAt:  *helper.DateTimeToStr(&f.CreatedAt),
	}
}
