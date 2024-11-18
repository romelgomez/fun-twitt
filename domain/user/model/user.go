package user

import (
	dto "funtwitt/domain/user/dto"
	"funtwitt/helper"
	"funtwitt/prisma/db"
)

type User struct {
	db.InnerUser
}

func (u *User) Dto() dto.UserResponse {

	return dto.UserResponse{
		ID:      u.ID,
		Name:    u.Name,
		Email:   u.Email,
		Created: *helper.DateTimeToStr(&u.Created),
		Updated: *helper.DateTimeToStr(&u.Updated),
		Deleted: helper.DateTimeToStr(u.Deleted),
	}
}
