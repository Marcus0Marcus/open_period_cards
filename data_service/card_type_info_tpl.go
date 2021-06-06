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

type cardTypeInfoTplService struct {
}

func NewCardTypeInfoTplService() *cardTypeInfoTplService {
	return &cardTypeInfoTplService{}
}

// cache function start
func (ms *cardTypeInfoTplService) cacheGetCardTypeInfoTpl(cond *CardTypeInfoTplInfo) (*response.FWError, *CardTypeInfoTplInfo) {
	if cond.Id == 0 {
		return constant.ErrCacheNotExist, nil
	}
	err, data := cachehelper.KeyGet("card_type_info_tpl_" + strconv.FormatUint(cond.Id,10))
	if err != nil {
		return err, nil
	}
	cardTypeInfoTplInfo := &CardTypeInfoTplInfo{}
	errJson := json.Unmarshal([]byte(data), cardTypeInfoTplInfo)
	if errJson != nil {
		return constant.ErrUnMarshal, nil
	}
	return nil, cardTypeInfoTplInfo
}

func (ms *cardTypeInfoTplService) cacheSetCardTypeInfoTpl(cardTypeInfoTplInfo *CardTypeInfoTplInfo) *response.FWError {
	cacheData, err := json.Marshal(cardTypeInfoTplInfo)
	if err != nil {
		return constant.ErrMarshal
	}
	return cachehelper.KeySet("card_type_info_tpl_" + strconv.FormatUint(cardTypeInfoTplInfo.Id,10), string(cacheData))
}

func (ms *cardTypeInfoTplService) cacheDelCardTypeInfoTpl(cardTypeInfoTplInfo *CardTypeInfoTplInfo) *response.FWError {
	return cachehelper.KeyDel("card_type_info_tpl_" + strconv.FormatUint(cardTypeInfoTplInfo.Id,10))
}

// cache function end
// db function start

func (ms *cardTypeInfoTplService) dbGetCardTypeInfoTplByCond(cond *CardTypeInfoTplInfo) (*response.FWError, *CardTypeInfoTplInfo) {
	cardTypeInfoTplInfo := &CardTypeInfoTplInfo{}
	err := dbhelper.GetDataByCond(cond, cardTypeInfoTplInfo)
	if err == constant.ErrDBNoRecord {
		return err, cardTypeInfoTplInfo
	}
	return nil, cardTypeInfoTplInfo
}

func (ms *cardTypeInfoTplService) dbCreateCardTypeInfoTpl(cardTypeInfoTplInfo *CardTypeInfoTplInfo) (*response.FWError, *CardTypeInfoTplInfo) {
	err, data := dbhelper.CreateData(cardTypeInfoTplInfo)
	cardTypeInfoTplInfo = data.(*CardTypeInfoTplInfo)
	if err != nil {
		return constant.ErrDb, nil
	}
	return nil, cardTypeInfoTplInfo
}
func (ms *cardTypeInfoTplService) dbUpdateCardTypeInfoTpl(cardTypeInfoTplInfo *CardTypeInfoTplInfo) (*response.FWError, int64) {
	return dbhelper.UpdateData(cardTypeInfoTplInfo)
}
func (ms *cardTypeInfoTplService) dbGetPagedCardTypeInfoTplByCond(cond *CardTypeInfoTplInfo, data *[]*CardTypeInfoTplInfo, pageNo int64, pageSize int64) (*response.FWError, int64) {
	return dbhelper.GetPagedDataByCond(cond, data, pageNo, pageSize)
}


// db function end

// service function start
func (ms *cardTypeInfoTplService) GetCardTypeInfoTplByCond(cond *CardTypeInfoTplInfo) (*response.FWError, *CardTypeInfoTplInfo) {
	// get cache
	cardTypeInfoTplInfo := &CardTypeInfoTplInfo{}
	err, cardTypeInfoTplInfo := ms.cacheGetCardTypeInfoTpl(cond)
	if err == nil {
		return nil, cardTypeInfoTplInfo
	}
	if err == constant.ErrCacheNotExist {
		// get db
		err, cardTypeInfoTplInfo = ms.dbGetCardTypeInfoTplByCond(cond)
		if err == constant.ErrCacheNotExist {
			return err, nil
		}
		if err != nil {
			return err, nil
		} else {
			return nil, cardTypeInfoTplInfo
		}
	} else {
		return err, nil
	}
}

func (ms *cardTypeInfoTplService) CreateCardTypeInfoTpl(cardTypeInfoTplInfo *CardTypeInfoTplInfo) (*response.FWError, *CardTypeInfoTplInfo) {
	cardTypeInfoTplInfo.Ctime = uint32(time.Now().Unix())
	cardTypeInfoTplInfo.Mtime = uint32(time.Now().Unix())
	err, cardTypeInfoTplInfo := ms.dbCreateCardTypeInfoTpl(cardTypeInfoTplInfo)
	if err != nil {
		openlog.Error(err.String() + " db create cardTypeInfoTpl failed.")
		return err, nil
	}
	err = ms.cacheSetCardTypeInfoTpl(cardTypeInfoTplInfo)
	if err != nil {
		openlog.Error(err.String() + " cache set cardTypeInfoTpl failed.")
	}
	return nil, cardTypeInfoTplInfo
}

func (ms *cardTypeInfoTplService) UpdateCardTypeInfoTpl(cardTypeInfoTplInfo *CardTypeInfoTplInfo) (*response.FWError, int64) {
	cardTypeInfoTplInfo.Mtime = uint32(time.Now().Unix())
	err, row := ms.dbUpdateCardTypeInfoTpl(cardTypeInfoTplInfo)
	if err != nil {
		return err, row
	}

	err = cachehelper.KeyDel("card_type_info_tpl_" + strconv.FormatUint(cardTypeInfoTplInfo.Id,10))
	if err != nil {
		openlog.Error("card_type_info_tpl_" + strconv.FormatUint(cardTypeInfoTplInfo.Id,10) + " cache del failed.")
	}
	return nil, row
}
func (ms *cardTypeInfoTplService) GetPagedCardTypeInfoTplByCond(cond *CardTypeInfoTplInfo, data *[]*CardTypeInfoTplInfo, pageNo int64, pageSize int64) (*response.FWError, int64) {
	return ms.dbGetPagedCardTypeInfoTplByCond(cond, data, pageNo, pageSize)
}

// service function end
