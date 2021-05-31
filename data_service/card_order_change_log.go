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

type cardOrderChangeLogService struct {
}

func NewCardOrderChangeLogService() *cardOrderChangeLogService {
	return &cardOrderChangeLogService{}
}

// cache function start
func (ms *cardOrderChangeLogService) cacheGetCardOrderChangeLog(cond *CardOrderChangeLogInfo) (*response.FWError, *CardOrderChangeLogInfo) {
	if cond.Id == 0 {
		return constant.ErrCacheNotExist, nil
	}
	err, data := cachehelper.KeyGet("card_order_change_log_" + strconv.FormatInt(cond.Id, 10))
	if err != nil {
		return err, nil
	}
	cardOrderChangeLogInfo := &CardOrderChangeLogInfo{}
	errJson := json.Unmarshal([]byte(data), cardOrderChangeLogInfo)
	if errJson != nil {
		return constant.ErrUnMarshal, nil
	}
	return nil, cardOrderChangeLogInfo
}

func (ms *cardOrderChangeLogService) cacheSetCardOrderChangeLog(cardOrderChangeLogInfo *CardOrderChangeLogInfo) *response.FWError {
	cacheData, err := json.Marshal(cardOrderChangeLogInfo)
	if err != nil {
		return constant.ErrMarshal
	}
	return cachehelper.KeySet("card_order_change_log_"+strconv.FormatInt(int64(cardOrderChangeLogInfo.Id), 10), string(cacheData))
}

func (ms *cardOrderChangeLogService) cacheDelCardOrderChangeLog(cardOrderChangeLogInfo *CardOrderChangeLogInfo) *response.FWError {
	return cachehelper.KeyDel("card_order_change_log_" + strconv.FormatInt(int64(cardOrderChangeLogInfo.Id), 10))
}

// cache function end
// db function start

func (ms *cardOrderChangeLogService) dbGetCardOrderChangeLogByCond(cond *CardOrderChangeLogInfo) (*response.FWError, *CardOrderChangeLogInfo) {
	cardOrderChangeLogInfo := &CardOrderChangeLogInfo{}
	err := dbhelper.GetDataByCond(cond, cardOrderChangeLogInfo)
	if err == constant.ErrDBNoRecord {
		return err, cardOrderChangeLogInfo
	}
	return nil, cardOrderChangeLogInfo
}

func (ms *cardOrderChangeLogService) dbCreateCardOrderChangeLog(cardOrderChangeLogInfo *CardOrderChangeLogInfo) (*response.FWError, *CardOrderChangeLogInfo) {
	err, data := dbhelper.CreateData(cardOrderChangeLogInfo)
	cardOrderChangeLogInfo = data.(*CardOrderChangeLogInfo)
	if err != nil {
		return constant.ErrDb, nil
	}
	return nil, cardOrderChangeLogInfo
}
func (ms *cardOrderChangeLogService) dbUpdateCardOrderChangeLog(cardOrderChangeLogInfo *CardOrderChangeLogInfo) (*response.FWError, int64) {
	err, row := dbhelper.UpdateData(cardOrderChangeLogInfo)
	return err, row
}

// db function end

// service function start
func (ms *cardOrderChangeLogService) GetCardOrderChangeLogByCond(cond *CardOrderChangeLogInfo) (*response.FWError, *CardOrderChangeLogInfo) {
	// get cache
	cardOrderChangeLogInfo := &CardOrderChangeLogInfo{}
	err, cardOrderChangeLogInfo := ms.cacheGetCardOrderChangeLog(cond)
	if err == nil {
		return nil, cardOrderChangeLogInfo
	}
	if err == constant.ErrCacheNotExist {
		// get db
		err, cardOrderChangeLogInfo = ms.dbGetCardOrderChangeLogByCond(cond)
		if err == constant.ErrCacheNotExist {
			return err, nil
		}
		if err != nil {
			return err, nil
		} else {
			return nil, cardOrderChangeLogInfo
		}
	} else {
		return err, nil
	}
}

func (ms *cardOrderChangeLogService) CreateCardOrderChangeLog(cardOrderChangeLogInfo *CardOrderChangeLogInfo) (*response.FWError, *CardOrderChangeLogInfo) {
	cardOrderChangeLogInfo.Ctime = uint32(time.Now().Unix())
	cardOrderChangeLogInfo.Mtime = uint32(time.Now().Unix())
	err, cardOrderChangeLogInfo := ms.dbCreateCardOrderChangeLog(cardOrderChangeLogInfo)
	if err != nil {
		openlog.Error(err.String() + " db create cardOrderChangeLog failed.")
		return err, nil
	}
	err = ms.cacheSetCardOrderChangeLog(cardOrderChangeLogInfo)
	if err != nil {
		openlog.Error(err.String() + " cache set cardOrderChangeLog failed.")
	}
	return nil, cardOrderChangeLogInfo
}

func (ms *cardOrderChangeLogService) UpdateCardOrderChangeLog(cardOrderChangeLogInfo *CardOrderChangeLogInfo) (*response.FWError, int64) {
	cardOrderChangeLogInfo.Mtime = uint32(time.Now().Unix())
	err, row := ms.dbUpdateCardOrderChangeLog(cardOrderChangeLogInfo)
	if err != nil {
		return err, row
	}

	err = cachehelper.KeyDel("card_order_change_log_" + strconv.FormatInt(int64(cardOrderChangeLogInfo.Id), 10))
	if err != nil {
		openlog.Error("card_order_change_log_" + strconv.FormatInt(int64(cardOrderChangeLogInfo.Id), 10) + " cache del failed.")
	}
	return nil, row
}

// service function end
