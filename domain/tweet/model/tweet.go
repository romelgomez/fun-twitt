package tweet

import (
	dto "funtwitt/domain/tweet/dto"
	"funtwitt/helper"
	"funtwitt/prisma/db"
)

type Tweet struct {
	db.InnerTweet
}

func (t *Tweet) Dto() dto.TweetResponse {
	return dto.TweetResponse{
		ID:        t.ID,
		Content:   t.Content,
		UserID:    t.UserID,
		CreatedAt: *helper.DateTimeToStr(&t.CreatedAt),
	}
}
