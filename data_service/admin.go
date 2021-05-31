package data_service

import (
	"encoding/json"
	"github.com/go-chassis/openlog"
	"open_period_cards/middleware/cachehelper"
	"open_period_cards/middleware/constant"
	"open_period_cards/middleware/dbhelper"
	"open_period_cards/middleware/response"
	"strconv"
	"time"
)

type adminService struct {
}

func NewAdminService() *adminService {
	return &adminService{}
}

// cache function start
func (ms *adminService) cacheGetAdmin(cond *AdminInfo) (*response.FWError, *AdminInfo) {
	if cond.Id == 0 {
		return constant.ErrCacheNotExist, nil
	}
	err, data := cachehelper.KeyGet("admin_" + strconv.FormatInt(cond.Id, 10))
	if err != nil {
		return err, nil
	}
	adminInfo := &AdminInfo{}
	errJson := json.Unmarshal([]byte(data), adminInfo)
	if errJson != nil {
		return constant.ErrUnMarshal, nil
	}
	return nil, adminInfo
}

func (ms *adminService) cacheSetAdmin(adminInfo *AdminInfo) *response.FWError {
	cacheData, err := json.Marshal(adminInfo)
	if err != nil {
		return constant.ErrMarshal
	}
	return cachehelper.KeySet("admin_"+strconv.FormatInt(int64(adminInfo.Id), 10), string(cacheData))
}

func (ms *adminService) cacheDelAdmin(adminInfo *AdminInfo) *response.FWError {
	return cachehelper.KeyDel("admin_" + strconv.FormatInt(int64(adminInfo.Id), 10))
}

// cache function end
// db function start

func (ms *adminService) dbGetAdminByCond(cond *AdminInfo) (*response.FWError, *AdminInfo) {
	adminInfo := &AdminInfo{}
	err := dbhelper.GetDataByCond(cond, adminInfo)
	if err == constant.ErrDBNoRecord {
		return err, adminInfo
	}
	return nil, adminInfo
}

func (ms *adminService) dbCreateAdmin(adminInfo *AdminInfo) (*response.FWError, *AdminInfo) {
	err, data := dbhelper.CreateData(adminInfo)
	adminInfo = data.(*AdminInfo)
	if err != nil {
		return constant.ErrDb, nil
	}
	return nil, adminInfo
}
func (ms *adminService) dbUpdateAdmin(adminInfo *AdminInfo) (*response.FWError, int64) {
	err, row := dbhelper.UpdateData(adminInfo)
	return err, row
}

// db function end

// service function start
func (ms *adminService) GetAdminByCond(cond *AdminInfo) (*response.FWError, *AdminInfo) {
	// get cache
	adminInfo := &AdminInfo{}
	err, adminInfo := ms.cacheGetAdmin(cond)
	if err == nil {
		return nil, adminInfo
	}
	if err == constant.ErrCacheNotExist {
		// get db
		err, adminInfo = ms.dbGetAdminByCond(cond)
		if err == constant.ErrCacheNotExist {
			return err, nil
		}
		if err != nil {
			return err, nil
		} else {
			return nil, adminInfo
		}
	} else {
		return err, nil
	}
}

func (ms *adminService) CreateAdmin(adminInfo *AdminInfo) (*response.FWError, *AdminInfo) {
	adminInfo.Ctime = uint32(time.Now().Unix())
	adminInfo.Mtime = uint32(time.Now().Unix())
	err, adminInfo := ms.dbCreateAdmin(adminInfo)
	if err != nil {
		openlog.Error(err.String() + " db create admin failed.")
		return err, nil
	}
	err = ms.cacheSetAdmin(adminInfo)
	if err != nil {
		openlog.Error(err.String() + " cache set admin failed.")
	}
	return nil, adminInfo
}

func (ms *adminService) UpdateAdmin(adminInfo *AdminInfo) (*response.FWError, int64) {
	adminInfo.Mtime = uint32(time.Now().Unix())
	err, row := ms.dbUpdateAdmin(adminInfo)
	if err != nil {
		return err, row
	}

	err = cachehelper.KeyDel("admin_" + strconv.FormatInt(int64(adminInfo.Id), 10))
	if err != nil {
		openlog.Error("admin_" + strconv.FormatInt(int64(adminInfo.Id), 10) + " cache del failed.")
	}
	return nil, row
}

// service function end
