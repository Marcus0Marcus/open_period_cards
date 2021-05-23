package service

import (
	"merchant/middleware/dbhelper"
	"merchant/middleware/response"
)

type userService struct {
}

func NewUserService() *userService {
	return &userService{}
}
func (r *userService) GetUserById(id int32) (*response.FWError, *UserInfo) {
	userInfo := &UserInfo{}
	err := dbhelper.GetDataById(id, userInfo)
	if err != nil {
		return err, nil
	}
	return nil, userInfo
}
