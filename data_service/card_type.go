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

type cardTypeService struct {
}

func NewCardTypeService() *cardTypeService {
	return &cardTypeService{}
}

// cache function start
func (ms *cardTypeService) cacheGetCardType(cond *CardTypeInfo) (*response.FWError, *CardTypeInfo) {
	if cond.Id == 0 {
		return constant.ErrCacheNotExist, nil
	}
	err, data := cachehelper.KeyGet("card_type_" + strconv.FormatUint(cond.Id,10))
	if err != nil {
		return err, nil
	}
	cardTypeInfo := &CardTypeInfo{}
	errJson := json.Unmarshal([]byte(data), cardTypeInfo)
	if errJson != nil {
		return constant.ErrUnMarshal, nil
	}
	return nil, cardTypeInfo
}

func (ms *cardTypeService) cacheSetCardType(cardTypeInfo *CardTypeInfo) *response.FWError {
	cacheData, err := json.Marshal(cardTypeInfo)
	if err != nil {
		return constant.ErrMarshal
	}
	return cachehelper.KeySet("card_type_" + strconv.FormatUint(cardTypeInfo.Id,10), string(cacheData))
}

func (ms *cardTypeService) cacheDelCardType(cardTypeInfo *CardTypeInfo) *response.FWError {
	return cachehelper.KeyDel("card_type_" + strconv.FormatUint(cardTypeInfo.Id,10))
}

// cache function end
// db function start

func (ms *cardTypeService) dbGetCardTypeByCond(cond *CardTypeInfo) (*response.FWError, *CardTypeInfo) {
	cardTypeInfo := &CardTypeInfo{}
	err := dbhelper.GetDataByCond(cond, cardTypeInfo)
	if err == constant.ErrDBNoRecord {
		return err, cardTypeInfo
	}
	return nil, cardTypeInfo
}

func (ms *cardTypeService) dbCreateCardType(cardTypeInfo *CardTypeInfo) (*response.FWError, *CardTypeInfo) {
	err, data := dbhelper.CreateData(cardTypeInfo)
	cardTypeInfo = data.(*CardTypeInfo)
	if err != nil {
		return constant.ErrDb, nil
	}
	return nil, cardTypeInfo
}
func (ms *cardTypeService) dbUpdateCardType(cardTypeInfo *CardTypeInfo) (*response.FWError, int64) {
	return dbhelper.UpdateData(cardTypeInfo)
}
func (ms *cardTypeService) dbGetPagedCardTypeByCond(cond *CardTypeInfo, data *[]*CardTypeInfo, pageNo int64, pageSize int64) (*response.FWError, int64) {
	return dbhelper.GetPagedDataByCond(cond, data, pageNo, pageSize)
}


// db function end

// service function start
func (ms *cardTypeService) GetCardTypeByCond(cond *CardTypeInfo) (*response.FWError, *CardTypeInfo) {
	// get cache
	cardTypeInfo := &CardTypeInfo{}
	err, cardTypeInfo := ms.cacheGetCardType(cond)
	if err == nil {
		return nil, cardTypeInfo
	}
	if err == constant.ErrCacheNotExist {
		// get db
		err, cardTypeInfo = ms.dbGetCardTypeByCond(cond)
		if err == constant.ErrCacheNotExist {
			return err, nil
		}
		if err != nil {
			return err, nil
		} else {
			return nil, cardTypeInfo
		}
	} else {
		return err, nil
	}
}

func (ms *cardTypeService) CreateCardType(cardTypeInfo *CardTypeInfo) (*response.FWError, *CardTypeInfo) {
	cardTypeInfo.Ctime = uint32(time.Now().Unix())
	cardTypeInfo.Mtime = uint32(time.Now().Unix())
	err, cardTypeInfo := ms.dbCreateCardType(cardTypeInfo)
	if err != nil {
		openlog.Error(err.String() + " db create cardType failed.")
		return err, nil
	}
	err = ms.cacheSetCardType(cardTypeInfo)
	if err != nil {
		openlog.Error(err.String() + " cache set cardType failed.")
	}
	return nil, cardTypeInfo
}

func (ms *cardTypeService) UpdateCardType(cardTypeInfo *CardTypeInfo) (*response.FWError, int64) {
	cardTypeInfo.Mtime = uint32(time.Now().Unix())
	err, row := ms.dbUpdateCardType(cardTypeInfo)
	if err != nil {
		return err, row
	}

	err = cachehelper.KeyDel("card_type_" + strconv.FormatUint(cardTypeInfo.Id,10))
	if err != nil {
		openlog.Error("card_type_" + strconv.FormatUint(cardTypeInfo.Id,10) + " cache del failed.")
	}
	return nil, row
}
func (ms *cardTypeService) GetPagedCardTypeByCond(cond *CardTypeInfo, data *[]*CardTypeInfo, pageNo int64, pageSize int64) (*response.FWError, int64) {
	return ms.dbGetPagedCardTypeByCond(cond, data, pageNo, pageSize)
}

// service function end
