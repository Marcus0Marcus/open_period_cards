package service

import (
	"admin/middleware/constant"
	"admin/middleware/dbhelper"
	"admin/middleware/response"
)

type adminService struct {
}

func NewAdminService() *adminService {
	return &adminService{}
}
func (r *adminService) GetAdminById(id int32) (*response.FWError, *AdminInfo) {
	adminInfo := &AdminInfo{}
	err := dbhelper.GetDataById(id, adminInfo)
	if err != nil {
		return constant.ErrDb, nil
	}
	return nil, adminInfo
}
func (r *adminService) GetAdminByCond(cond *AdminInfo) (*response.FWError, *AdminInfo) {
	adminInfo := &AdminInfo{}
	err := dbhelper.GetDataByCond(cond, adminInfo)
	if err != nil {
		return constant.ErrDb, nil
	}
	return nil, adminInfo
}
