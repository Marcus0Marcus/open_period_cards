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

type cardOrderDeliveryLogService struct {
}

func NewCardOrderDeliveryLogService() *cardOrderDeliveryLogService {
	return &cardOrderDeliveryLogService{}
}

// cache function start
func (ms *cardOrderDeliveryLogService) cacheGetCardOrderDeliveryLog(cond *CardOrderDeliveryLogInfo) (*response.FWError, *CardOrderDeliveryLogInfo) {
	if cond.Id == 0 {
		return constant.ErrCacheNotExist, nil
	}
	err, data := cachehelper.KeyGet("card_order_delivery_log_" + strconv.FormatInt(cond.Id, 10))
	if err != nil {
		return err, nil
	}
	cardOrderDeliveryLogInfo := &CardOrderDeliveryLogInfo{}
	errJson := json.Unmarshal([]byte(data), cardOrderDeliveryLogInfo)
	if errJson != nil {
		return constant.ErrUnMarshal, nil
	}
	return nil, cardOrderDeliveryLogInfo
}

func (ms *cardOrderDeliveryLogService) cacheSetCardOrderDeliveryLog(cardOrderDeliveryLogInfo *CardOrderDeliveryLogInfo) *response.FWError {
	cacheData, err := json.Marshal(cardOrderDeliveryLogInfo)
	if err != nil {
		return constant.ErrMarshal
	}
	return cachehelper.KeySet("card_order_delivery_log_"+strconv.FormatInt(int64(cardOrderDeliveryLogInfo.Id), 10), string(cacheData))
}

func (ms *cardOrderDeliveryLogService) cacheDelCardOrderDeliveryLog(cardOrderDeliveryLogInfo *CardOrderDeliveryLogInfo) *response.FWError {
	return cachehelper.KeyDel("card_order_delivery_log_" + strconv.FormatInt(int64(cardOrderDeliveryLogInfo.Id), 10))
}

// cache function end
// db function start

func (ms *cardOrderDeliveryLogService) dbGetCardOrderDeliveryLogByCond(cond *CardOrderDeliveryLogInfo) (*response.FWError, *CardOrderDeliveryLogInfo) {
	cardOrderDeliveryLogInfo := &CardOrderDeliveryLogInfo{}
	err := dbhelper.GetDataByCond(cond, cardOrderDeliveryLogInfo)
	if err == constant.ErrDBNoRecord {
		return err, cardOrderDeliveryLogInfo
	}
	return nil, cardOrderDeliveryLogInfo
}

func (ms *cardOrderDeliveryLogService) dbCreateCardOrderDeliveryLog(cardOrderDeliveryLogInfo *CardOrderDeliveryLogInfo) (*response.FWError, *CardOrderDeliveryLogInfo) {
	err, data := dbhelper.CreateData(cardOrderDeliveryLogInfo)
	cardOrderDeliveryLogInfo = data.(*CardOrderDeliveryLogInfo)
	if err != nil {
		return constant.ErrDb, nil
	}
	return nil, cardOrderDeliveryLogInfo
}
func (ms *cardOrderDeliveryLogService) dbUpdateCardOrderDeliveryLog(cardOrderDeliveryLogInfo *CardOrderDeliveryLogInfo) (*response.FWError, int64) {
	err, row := dbhelper.UpdateData(cardOrderDeliveryLogInfo)
	return err, row
}

// db function end

// service function start
func (ms *cardOrderDeliveryLogService) GetCardOrderDeliveryLogByCond(cond *CardOrderDeliveryLogInfo) (*response.FWError, *CardOrderDeliveryLogInfo) {
	// get cache
	cardOrderDeliveryLogInfo := &CardOrderDeliveryLogInfo{}
	err, cardOrderDeliveryLogInfo := ms.cacheGetCardOrderDeliveryLog(cond)
	if err == nil {
		return nil, cardOrderDeliveryLogInfo
	}
	if err == constant.ErrCacheNotExist {
		// get db
		err, cardOrderDeliveryLogInfo = ms.dbGetCardOrderDeliveryLogByCond(cond)
		if err == constant.ErrCacheNotExist {
			return err, nil
		}
		if err != nil {
			return err, nil
		} else {
			return nil, cardOrderDeliveryLogInfo
		}
	} else {
		return err, nil
	}
}

func (ms *cardOrderDeliveryLogService) CreateCardOrderDeliveryLog(cardOrderDeliveryLogInfo *CardOrderDeliveryLogInfo) (*response.FWError, *CardOrderDeliveryLogInfo) {
	cardOrderDeliveryLogInfo.Ctime = uint32(time.Now().Unix())
	cardOrderDeliveryLogInfo.Mtime = uint32(time.Now().Unix())
	err, cardOrderDeliveryLogInfo := ms.dbCreateCardOrderDeliveryLog(cardOrderDeliveryLogInfo)
	if err != nil {
		openlog.Error(err.String() + " db create cardOrderDeliveryLog failed.")
		return err, nil
	}
	err = ms.cacheSetCardOrderDeliveryLog(cardOrderDeliveryLogInfo)
	if err != nil {
		openlog.Error(err.String() + " cache set cardOrderDeliveryLog failed.")
	}
	return nil, cardOrderDeliveryLogInfo
}

func (ms *cardOrderDeliveryLogService) UpdateCardOrderDeliveryLog(cardOrderDeliveryLogInfo *CardOrderDeliveryLogInfo) (*response.FWError, int64) {
	cardOrderDeliveryLogInfo.Mtime = uint32(time.Now().Unix())
	err, row := ms.dbUpdateCardOrderDeliveryLog(cardOrderDeliveryLogInfo)
	if err != nil {
		return err, row
	}

	err = cachehelper.KeyDel("card_order_delivery_log_" + strconv.FormatInt(int64(cardOrderDeliveryLogInfo.Id), 10))
	if err != nil {
		openlog.Error("card_order_delivery_log_" + strconv.FormatInt(int64(cardOrderDeliveryLogInfo.Id), 10) + " cache del failed.")
	}
	return nil, row
}

// service function end
