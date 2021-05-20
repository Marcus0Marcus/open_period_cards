package service

import "admin/middleware/dbhelper"

type userService struct {
}

func NewUserService() *userService {
	return &userService{}
}
func (r *userService) GetUserById(id int32) (error, *UserInfo) {
	userInfo := &UserInfo{}
	err := dbhelper.GetDataById(id, userInfo)
	if err != nil {
		return err, nil
	}
	return nil, userInfo
}
