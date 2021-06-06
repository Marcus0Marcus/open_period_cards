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

type userService struct {
}

func NewUserService() *userService {
	return &userService{}
}

// cache function start
func (ms *userService) cacheGetUser(cond *UserInfo) (*response.FWError, *UserInfo) {
	if cond.Id == 0 {
		return constant.ErrCacheNotExist, nil
	}
	err, data := cachehelper.KeyGet("user_" + strconv.FormatUint(cond.Id,10))
	if err != nil {
		return err, nil
	}
	userInfo := &UserInfo{}
	errJson := json.Unmarshal([]byte(data), userInfo)
	if errJson != nil {
		return constant.ErrUnMarshal, nil
	}
	return nil, userInfo
}

func (ms *userService) cacheSetUser(userInfo *UserInfo) *response.FWError {
	cacheData, err := json.Marshal(userInfo)
	if err != nil {
		return constant.ErrMarshal
	}
	return cachehelper.KeySet("user_" + strconv.FormatUint(userInfo.Id,10), string(cacheData))
}

func (ms *userService) cacheDelUser(userInfo *UserInfo) *response.FWError {
	return cachehelper.KeyDel("user_" + strconv.FormatUint(userInfo.Id,10))
}

// cache function end
// db function start

func (ms *userService) dbGetUserByCond(cond *UserInfo) (*response.FWError, *UserInfo) {
	userInfo := &UserInfo{}
	err := dbhelper.GetDataByCond(cond, userInfo)
	if err == constant.ErrDBNoRecord {
		return err, userInfo
	}
	return nil, userInfo
}

func (ms *userService) dbCreateUser(userInfo *UserInfo) (*response.FWError, *UserInfo) {
	err, data := dbhelper.CreateData(userInfo)
	userInfo = data.(*UserInfo)
	if err != nil {
		return constant.ErrDb, nil
	}
	return nil, userInfo
}
func (ms *userService) dbUpdateUser(userInfo *UserInfo) (*response.FWError, int64) {
	return dbhelper.UpdateData(userInfo)
}
func (ms *userService) dbGetPagedUserByCond(cond *UserInfo, data *[]*UserInfo, pageNo int64, pageSize int64) (*response.FWError, int64) {
	return dbhelper.GetPagedDataByCond(cond, data, pageNo, pageSize)
}


// db function end

// service function start
func (ms *userService) GetUserByCond(cond *UserInfo) (*response.FWError, *UserInfo) {
	// get cache
	userInfo := &UserInfo{}
	err, userInfo := ms.cacheGetUser(cond)
	if err == nil {
		return nil, userInfo
	}
	if err == constant.ErrCacheNotExist {
		// get db
		err, userInfo = ms.dbGetUserByCond(cond)
		if err == constant.ErrCacheNotExist {
			return err, nil
		}
		if err != nil {
			return err, nil
		} else {
			return nil, userInfo
		}
	} else {
		return err, nil
	}
}

func (ms *userService) CreateUser(userInfo *UserInfo) (*response.FWError, *UserInfo) {
	userInfo.Ctime = uint32(time.Now().Unix())
	userInfo.Mtime = uint32(time.Now().Unix())
	err, userInfo := ms.dbCreateUser(userInfo)
	if err != nil {
		openlog.Error(err.String() + " db create user failed.")
		return err, nil
	}
	err = ms.cacheSetUser(userInfo)
	if err != nil {
		openlog.Error(err.String() + " cache set user failed.")
	}
	return nil, userInfo
}

func (ms *userService) UpdateUser(userInfo *UserInfo) (*response.FWError, int64) {
	userInfo.Mtime = uint32(time.Now().Unix())
	err, row := ms.dbUpdateUser(userInfo)
	if err != nil {
		return err, row
	}

	err = cachehelper.KeyDel("user_" + strconv.FormatUint(userInfo.Id,10))
	if err != nil {
		openlog.Error("user_" + strconv.FormatUint(userInfo.Id,10) + " cache del failed.")
	}
	return nil, row
}
func (ms *userService) GetPagedUserByCond(cond *UserInfo, data *[]*UserInfo, pageNo int64, pageSize int64) (*response.FWError, int64) {
	return ms.dbGetPagedUserByCond(cond, data, pageNo, pageSize)
}

// service function end
