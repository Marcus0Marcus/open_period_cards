package service

import (
	"encoding/json"
	"github.com/go-chassis/openlog"
	"merchant/middleware/cachehelper"
	"merchant/middleware/constant"
	"merchant/middleware/dbhelper"
	"merchant/middleware/response"
	"time"
)

type merchantService struct {
}

func NewMerchantService() *merchantService {
	return &merchantService{}
}

// cache function start
func (ms *merchantService) cacheGetMerchant(cond *MerchantInfo) (*response.FWError, *MerchantInfo) {
	if cond.Phone == "" {
		return constant.ErrCacheNotExist, nil
	}
	err, data := cachehelper.KeyGet(cond.Phone)
	if err != nil {
		return err, nil
	}
	merchantInfo := &MerchantInfo{}
	errJson := json.Unmarshal([]byte(data), merchantInfo)
	if errJson != nil {
		return constant.ErrUnMarshal, nil
	}
	return nil, merchantInfo
}

func (ms *merchantService) cacheSetMerchant(merchantInfo *MerchantInfo) *response.FWError {
	cacheData, err := json.Marshal(merchantInfo)
	if err != nil {
		return constant.ErrMarshal
	}
	return cachehelper.KeySet(merchantInfo.Phone, string(cacheData))
}

func (ms *merchantService) cacheDelMerchant(merchantInfo *MerchantInfo) *response.FWError {
	return cachehelper.KeyDel(merchantInfo.Phone)
}

// cache function end
// db function start

func (ms *merchantService) dbGetMerchantByCond(cond *MerchantInfo) (*response.FWError, *MerchantInfo) {
	merchantInfo := &MerchantInfo{}
	err := dbhelper.GetDataByCond(cond, merchantInfo)
	if err == constant.ErrDBNoRecord {
		return err, merchantInfo
	}
	return nil, merchantInfo
}

func (ms *merchantService) dbCreateMerchant(merchantInfo *MerchantInfo) (*response.FWError, *MerchantInfo) {
	err, data := dbhelper.CreateData(merchantInfo)
	merchantInfo = data.(*MerchantInfo)
	if err != nil {
		openlog.Info(err.Message)
		return constant.ErrDb, nil
	}
	return nil, merchantInfo
}
func (ms *merchantService) dbUpdateMerchant(merchantInfo *MerchantInfo) (*response.FWError, int64) {
	err, row := dbhelper.UpdateData(merchantInfo)
	return err, row
}

// db function end

// service function start
func (ms *merchantService) GetMerchantByCond(cond *MerchantInfo) (*response.FWError, *MerchantInfo) {
	// get cache
	merchantInfo := &MerchantInfo{}
	err, merchantInfo := ms.cacheGetMerchant(cond)
	if err == nil {
		return nil, merchantInfo
	}
	if err == constant.ErrCacheNotExist {
		// get db
		err, merchantInfo = ms.dbGetMerchantByCond(cond)
		if err == constant.ErrCacheNotExist {
			return err, nil
		}
		if err != nil {
			return err, nil
		} else {
			return nil, merchantInfo
		}
	} else {
		return err, nil
	}
}

func (ms *merchantService) CreateMerchant(merchantInfo *MerchantInfo) (*response.FWError, *MerchantInfo) {
	merchantInfo.Ctime = time.Now().Unix()
	merchantInfo.Mtime = time.Now().Unix()
	err, merchantInfo := ms.dbCreateMerchant(merchantInfo)
	if err != nil {
		openlog.Error(err.String() + " db create merchant failed.")
		return err, nil
	}
	err = ms.cacheSetMerchant(merchantInfo)
	if err != nil {
		openlog.Error(err.String() + " cache set merchant failed.")
	}
	return nil, merchantInfo
}

func (ms *merchantService) UpdateMerchant(merchantInfo *MerchantInfo) (*response.FWError, int64) {
	merchantInfo.Mtime = time.Now().Unix()
	openlog.Debug(merchantInfo.IndustryName)
	err, row := ms.dbUpdateMerchant(merchantInfo)
	if err != nil {
		return err, row
	}
	err = cachehelper.KeyDel(merchantInfo.Phone)
	if err != nil {
		openlog.Error(merchantInfo.Phone + " cache del failed.")
	}
	return nil, row
}

// service function end
