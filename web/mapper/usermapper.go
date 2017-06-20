package mapper

import (
	"gitlab.com/parkby/parkby-service/model"
	"gitlab.com/parkby/parkby-service/web/dto"
)

type userMapper struct {
}

// UserMapper maps user model to dto
var UserMapper userMapper

func (mapper userMapper) ToModel(dto dto.UserRegisterDTO) *model.User {
	var user model.User
	user.LastName = dto.LatName
	user.FirstName = dto.FirstName
	user.PhoneNumber = dto.PhoneNumber
	user.Email = dto.Email
	user.Password = dto.Password
	return &user
}
