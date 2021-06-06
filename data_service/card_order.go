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

type cardOrderService struct {
}

func NewCardOrderService() *cardOrderService {
	return &cardOrderService{}
}

// cache function start
func (ms *cardOrderService) cacheGetCardOrder(cond *CardOrderInfo) (*response.FWError, *CardOrderInfo) {
	if cond.Id == 0 {
		return constant.ErrCacheNotExist, nil
	}
	err, data := cachehelper.KeyGet("card_order_" + strconv.FormatUint(cond.Id,10))
	if err != nil {
		return err, nil
	}
	cardOrderInfo := &CardOrderInfo{}
	errJson := json.Unmarshal([]byte(data), cardOrderInfo)
	if errJson != nil {
		return constant.ErrUnMarshal, nil
	}
	return nil, cardOrderInfo
}

func (ms *cardOrderService) cacheSetCardOrder(cardOrderInfo *CardOrderInfo) *response.FWError {
	cacheData, err := json.Marshal(cardOrderInfo)
	if err != nil {
		return constant.ErrMarshal
	}
	return cachehelper.KeySet("card_order_" + strconv.FormatUint(cardOrderInfo.Id,10), string(cacheData))
}

func (ms *cardOrderService) cacheDelCardOrder(cardOrderInfo *CardOrderInfo) *response.FWError {
	return cachehelper.KeyDel("card_order_" + strconv.FormatUint(cardOrderInfo.Id,10))
}

// cache function end
// db function start

func (ms *cardOrderService) dbGetCardOrderByCond(cond *CardOrderInfo) (*response.FWError, *CardOrderInfo) {
	cardOrderInfo := &CardOrderInfo{}
	err := dbhelper.GetDataByCond(cond, cardOrderInfo)
	if err == constant.ErrDBNoRecord {
		return err, cardOrderInfo
	}
	return nil, cardOrderInfo
}

func (ms *cardOrderService) dbCreateCardOrder(cardOrderInfo *CardOrderInfo) (*response.FWError, *CardOrderInfo) {
	err, data := dbhelper.CreateData(cardOrderInfo)
	cardOrderInfo = data.(*CardOrderInfo)
	if err != nil {
		return constant.ErrDb, nil
	}
	return nil, cardOrderInfo
}
func (ms *cardOrderService) dbUpdateCardOrder(cardOrderInfo *CardOrderInfo) (*response.FWError, int64) {
	return dbhelper.UpdateData(cardOrderInfo)
}
func (ms *cardOrderService) dbGetPagedCardOrderByCond(cond *CardOrderInfo, data *[]*CardOrderInfo, pageNo int64, pageSize int64) (*response.FWError, int64) {
	return dbhelper.GetPagedDataByCond(cond, data, pageNo, pageSize)
}


// db function end

// service function start
func (ms *cardOrderService) GetCardOrderByCond(cond *CardOrderInfo) (*response.FWError, *CardOrderInfo) {
	// get cache
	cardOrderInfo := &CardOrderInfo{}
	err, cardOrderInfo := ms.cacheGetCardOrder(cond)
	if err == nil {
		return nil, cardOrderInfo
	}
	if err == constant.ErrCacheNotExist {
		// get db
		err, cardOrderInfo = ms.dbGetCardOrderByCond(cond)
		if err == constant.ErrCacheNotExist {
			return err, nil
		}
		if err != nil {
			return err, nil
		} else {
			return nil, cardOrderInfo
		}
	} else {
		return err, nil
	}
}

func (ms *cardOrderService) CreateCardOrder(cardOrderInfo *CardOrderInfo) (*response.FWError, *CardOrderInfo) {
	cardOrderInfo.Ctime = uint32(time.Now().Unix())
	cardOrderInfo.Mtime = uint32(time.Now().Unix())
	err, cardOrderInfo := ms.dbCreateCardOrder(cardOrderInfo)
	if err != nil {
		openlog.Error(err.String() + " db create cardOrder failed.")
		return err, nil
	}
	err = ms.cacheSetCardOrder(cardOrderInfo)
	if err != nil {
		openlog.Error(err.String() + " cache set cardOrder failed.")
	}
	return nil, cardOrderInfo
}

func (ms *cardOrderService) UpdateCardOrder(cardOrderInfo *CardOrderInfo) (*response.FWError, int64) {
	cardOrderInfo.Mtime = uint32(time.Now().Unix())
	err, row := ms.dbUpdateCardOrder(cardOrderInfo)
	if err != nil {
		return err, row
	}

	err = cachehelper.KeyDel("card_order_" + strconv.FormatUint(cardOrderInfo.Id,10))
	if err != nil {
		openlog.Error("card_order_" + strconv.FormatUint(cardOrderInfo.Id,10) + " cache del failed.")
	}
	return nil, row
}
func (ms *cardOrderService) GetPagedCardOrderByCond(cond *CardOrderInfo, data *[]*CardOrderInfo, pageNo int64, pageSize int64) (*response.FWError, int64) {
	return ms.dbGetPagedCardOrderByCond(cond, data, pageNo, pageSize)
}

// service function end
