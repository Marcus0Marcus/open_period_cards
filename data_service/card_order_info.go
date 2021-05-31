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

type cardOrderInfoService struct {
}

func NewCardOrderInfoService() *cardOrderInfoService {
	return &cardOrderInfoService{}
}

// cache function start
func (ms *cardOrderInfoService) cacheGetCardOrderInfo(cond *CardOrderInfoInfo) (*response.FWError, *CardOrderInfoInfo) {
	if cond.Id == 0 {
		return constant.ErrCacheNotExist, nil
	}
	err, data := cachehelper.KeyGet("card_order_info_" + strconv.FormatInt(cond.Id, 10))
	if err != nil {
		return err, nil
	}
	cardOrderInfoInfo := &CardOrderInfoInfo{}
	errJson := json.Unmarshal([]byte(data), cardOrderInfoInfo)
	if errJson != nil {
		return constant.ErrUnMarshal, nil
	}
	return nil, cardOrderInfoInfo
}

func (ms *cardOrderInfoService) cacheSetCardOrderInfo(cardOrderInfoInfo *CardOrderInfoInfo) *response.FWError {
	cacheData, err := json.Marshal(cardOrderInfoInfo)
	if err != nil {
		return constant.ErrMarshal
	}
	return cachehelper.KeySet("card_order_info_"+strconv.FormatInt(int64(cardOrderInfoInfo.Id), 10), string(cacheData))
}

func (ms *cardOrderInfoService) cacheDelCardOrderInfo(cardOrderInfoInfo *CardOrderInfoInfo) *response.FWError {
	return cachehelper.KeyDel("card_order_info_" + strconv.FormatInt(int64(cardOrderInfoInfo.Id), 10))
}

// cache function end
// db function start

func (ms *cardOrderInfoService) dbGetCardOrderInfoByCond(cond *CardOrderInfoInfo) (*response.FWError, *CardOrderInfoInfo) {
	cardOrderInfoInfo := &CardOrderInfoInfo{}
	err := dbhelper.GetDataByCond(cond, cardOrderInfoInfo)
	if err == constant.ErrDBNoRecord {
		return err, cardOrderInfoInfo
	}
	return nil, cardOrderInfoInfo
}

func (ms *cardOrderInfoService) dbCreateCardOrderInfo(cardOrderInfoInfo *CardOrderInfoInfo) (*response.FWError, *CardOrderInfoInfo) {
	err, data := dbhelper.CreateData(cardOrderInfoInfo)
	cardOrderInfoInfo = data.(*CardOrderInfoInfo)
	if err != nil {
		return constant.ErrDb, nil
	}
	return nil, cardOrderInfoInfo
}
func (ms *cardOrderInfoService) dbUpdateCardOrderInfo(cardOrderInfoInfo *CardOrderInfoInfo) (*response.FWError, int64) {
	err, row := dbhelper.UpdateData(cardOrderInfoInfo)
	return err, row
}

// db function end

// service function start
func (ms *cardOrderInfoService) GetCardOrderInfoByCond(cond *CardOrderInfoInfo) (*response.FWError, *CardOrderInfoInfo) {
	// get cache
	cardOrderInfoInfo := &CardOrderInfoInfo{}
	err, cardOrderInfoInfo := ms.cacheGetCardOrderInfo(cond)
	if err == nil {
		return nil, cardOrderInfoInfo
	}
	if err == constant.ErrCacheNotExist {
		// get db
		err, cardOrderInfoInfo = ms.dbGetCardOrderInfoByCond(cond)
		if err == constant.ErrCacheNotExist {
			return err, nil
		}
		if err != nil {
			return err, nil
		} else {
			return nil, cardOrderInfoInfo
		}
	} else {
		return err, nil
	}
}

func (ms *cardOrderInfoService) CreateCardOrderInfo(cardOrderInfoInfo *CardOrderInfoInfo) (*response.FWError, *CardOrderInfoInfo) {
	cardOrderInfoInfo.Ctime = uint32(time.Now().Unix())
	cardOrderInfoInfo.Mtime = uint32(time.Now().Unix())
	err, cardOrderInfoInfo := ms.dbCreateCardOrderInfo(cardOrderInfoInfo)
	if err != nil {
		openlog.Error(err.String() + " db create cardOrderInfo failed.")
		return err, nil
	}
	err = ms.cacheSetCardOrderInfo(cardOrderInfoInfo)
	if err != nil {
		openlog.Error(err.String() + " cache set cardOrderInfo failed.")
	}
	return nil, cardOrderInfoInfo
}

func (ms *cardOrderInfoService) UpdateCardOrderInfo(cardOrderInfoInfo *CardOrderInfoInfo) (*response.FWError, int64) {
	cardOrderInfoInfo.Mtime = uint32(time.Now().Unix())
	err, row := ms.dbUpdateCardOrderInfo(cardOrderInfoInfo)
	if err != nil {
		return err, row
	}

	err = cachehelper.KeyDel("card_order_info_" + strconv.FormatInt(int64(cardOrderInfoInfo.Id), 10))
	if err != nil {
		openlog.Error("card_order_info_" + strconv.FormatInt(int64(cardOrderInfoInfo.Id), 10) + " cache del failed.")
	}
	return nil, row
}

// service function end
