package data_service

import (
	"encoding/json"
	"github.com/go-chassis/openlog"
	"{package_name}/middleware/cachehelper"
	"{package_name}/middleware/constant"
	"{package_name}/middleware/dbhelper"
	"{package_name}/middleware/response"
	"strconv"
	"time"

)

type {ServiceModelLower}Service struct {
}

func New{ServiceModelUpper}Service() *{ServiceModelLower}Service {
	return &{ServiceModelLower}Service{}
}

// cache function start
func (ms *{ServiceModelLower}Service) cacheGet{ServiceModelUpper}(cond *{ServiceModelUpper}{ServiceModelSuffix}) (*response.FWError, *{ServiceModelUpper}{ServiceModelSuffix}) {
	if cond.Id == 0 {
		return constant.ErrCacheNotExist, nil
	}
	err, data := cachehelper.KeyGet("{CacheKeyPrefix}" + strconv.FormatUint(cond.Id,10))
	if err != nil {
		return err, nil
	}
	{ServiceModelLower}{ServiceModelSuffix} := &{ServiceModelUpper}{ServiceModelSuffix}{}
	errJson := json.Unmarshal([]byte(data), {ServiceModelLower}{ServiceModelSuffix})
	if errJson != nil {
		return constant.ErrUnMarshal, nil
	}
	return nil, {ServiceModelLower}{ServiceModelSuffix}
}

func (ms *{ServiceModelLower}Service) cacheSet{ServiceModelUpper}({ServiceModelLower}{ServiceModelSuffix} *{ServiceModelUpper}{ServiceModelSuffix}) *response.FWError {
	cacheData, err := json.Marshal({ServiceModelLower}{ServiceModelSuffix})
	if err != nil {
		return constant.ErrMarshal
	}
	return cachehelper.KeySet("{CacheKeyPrefix}" + strconv.FormatUint({ServiceModelLower}{ServiceModelSuffix}.Id,10), string(cacheData))
}

func (ms *{ServiceModelLower}Service) cacheDel{ServiceModelUpper}({ServiceModelLower}{ServiceModelSuffix} *{ServiceModelUpper}{ServiceModelSuffix}) *response.FWError {
	return cachehelper.KeyDel("{CacheKeyPrefix}" + strconv.FormatUint({ServiceModelLower}{ServiceModelSuffix}.Id,10))
}

// cache function end
// db function start

func (ms *{ServiceModelLower}Service) dbGet{ServiceModelUpper}ByCond(cond *{ServiceModelUpper}{ServiceModelSuffix}) (*response.FWError, *{ServiceModelUpper}{ServiceModelSuffix}) {
	{ServiceModelLower}{ServiceModelSuffix} := &{ServiceModelUpper}{ServiceModelSuffix}{}
	err := dbhelper.GetDataByCond(cond, {ServiceModelLower}{ServiceModelSuffix})
	if err == constant.ErrDBNoRecord {
		return err, {ServiceModelLower}{ServiceModelSuffix}
	}
	return nil, {ServiceModelLower}{ServiceModelSuffix}
}

func (ms *{ServiceModelLower}Service) dbCreate{ServiceModelUpper}({ServiceModelLower}{ServiceModelSuffix} *{ServiceModelUpper}{ServiceModelSuffix}) (*response.FWError, *{ServiceModelUpper}{ServiceModelSuffix}) {
	err, data := dbhelper.CreateData({ServiceModelLower}{ServiceModelSuffix})
	{ServiceModelLower}{ServiceModelSuffix} = data.(*{ServiceModelUpper}{ServiceModelSuffix})
	if err != nil {
		return constant.ErrDb, nil
	}
	return nil, {ServiceModelLower}{ServiceModelSuffix}
}
func (ms *{ServiceModelLower}Service) dbUpdate{ServiceModelUpper}({ServiceModelLower}{ServiceModelSuffix} *{ServiceModelUpper}{ServiceModelSuffix}) (*response.FWError, int64) {
	return dbhelper.UpdateData({ServiceModelLower}{ServiceModelSuffix})
}
func (ms *{ServiceModelLower}Service) dbGetPaged{ServiceModelUpper}ByCond(cond *{ServiceModelUpper}{ServiceModelSuffix}, data *[]*{ServiceModelUpper}{ServiceModelSuffix}, pageNo int64, pageSize int64) (*response.FWError, int64) {
	return dbhelper.GetPagedDataByCond(cond, data, pageNo, pageSize)
}


// db function end

// service function start
func (ms *{ServiceModelLower}Service) Get{ServiceModelUpper}ByCond(cond *{ServiceModelUpper}{ServiceModelSuffix}) (*response.FWError, *{ServiceModelUpper}{ServiceModelSuffix}) {
	// get cache
	{ServiceModelLower}{ServiceModelSuffix} := &{ServiceModelUpper}{ServiceModelSuffix}{}
	err, {ServiceModelLower}{ServiceModelSuffix} := ms.cacheGet{ServiceModelUpper}(cond)
	if err == nil {
		return nil, {ServiceModelLower}{ServiceModelSuffix}
	}
	if err == constant.ErrCacheNotExist {
		// get db
		err, {ServiceModelLower}{ServiceModelSuffix} = ms.dbGet{ServiceModelUpper}ByCond(cond)
		if err == constant.ErrCacheNotExist {
			return err, nil
		}
		if err != nil {
			return err, nil
		} else {
			return nil, {ServiceModelLower}{ServiceModelSuffix}
		}
	} else {
		return err, nil
	}
}

func (ms *{ServiceModelLower}Service) Create{ServiceModelUpper}({ServiceModelLower}{ServiceModelSuffix} *{ServiceModelUpper}{ServiceModelSuffix}) (*response.FWError, *{ServiceModelUpper}{ServiceModelSuffix}) {
	{ServiceModelLower}{ServiceModelSuffix}.Ctime = uint32(time.Now().Unix())
	{ServiceModelLower}{ServiceModelSuffix}.Mtime = uint32(time.Now().Unix())
	err, {ServiceModelLower}{ServiceModelSuffix} := ms.dbCreate{ServiceModelUpper}({ServiceModelLower}{ServiceModelSuffix})
	if err != nil {
		openlog.Error(err.String() + " db create {ServiceModelLower} failed.")
		return err, nil
	}
	err = ms.cacheSet{ServiceModelUpper}({ServiceModelLower}{ServiceModelSuffix})
	if err != nil {
		openlog.Error(err.String() + " cache set {ServiceModelLower} failed.")
	}
	return nil, {ServiceModelLower}{ServiceModelSuffix}
}

func (ms *{ServiceModelLower}Service) Update{ServiceModelUpper}({ServiceModelLower}{ServiceModelSuffix} *{ServiceModelUpper}{ServiceModelSuffix}) (*response.FWError, int64) {
	{ServiceModelLower}{ServiceModelSuffix}.Mtime = uint32(time.Now().Unix())
	err, row := ms.dbUpdate{ServiceModelUpper}({ServiceModelLower}{ServiceModelSuffix})
	if err != nil {
		return err, row
	}

	err = cachehelper.KeyDel("{CacheKeyPrefix}" + strconv.FormatUint({ServiceModelLower}{ServiceModelSuffix}.Id,10))
	if err != nil {
		openlog.Error("{CacheKeyPrefix}" + strconv.FormatUint({ServiceModelLower}{ServiceModelSuffix}.Id,10) + " cache del failed.")
	}
	return nil, row
}
func (ms *{ServiceModelLower}Service) GetPaged{ServiceModelUpper}ByCond(cond *{ServiceModelUpper}{ServiceModelSuffix}, data *[]*{ServiceModelUpper}{ServiceModelSuffix}, pageNo int64, pageSize int64) (*response.FWError, int64) {
	return ms.dbGetPaged{ServiceModelUpper}ByCond(cond, data, pageNo, pageSize)
}

// service function end
