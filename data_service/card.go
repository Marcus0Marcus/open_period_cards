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

type cardService struct {
}

func NewCardService() *cardService {
	return &cardService{}
}

// cache function start
func (ms *cardService) cacheGetCard(cond *CardInfo) (*response.FWError, *CardInfo) {
	if cond.Id == 0 {
		return constant.ErrCacheNotExist, nil
	}
	err, data := cachehelper.KeyGet("card_" + strconv.FormatUint(cond.Id,10))
	if err != nil {
		return err, nil
	}
	cardInfo := &CardInfo{}
	errJson := json.Unmarshal([]byte(data), cardInfo)
	if errJson != nil {
		return constant.ErrUnMarshal, nil
	}
	return nil, cardInfo
}

func (ms *cardService) cacheSetCard(cardInfo *CardInfo) *response.FWError {
	cacheData, err := json.Marshal(cardInfo)
	if err != nil {
		return constant.ErrMarshal
	}
	return cachehelper.KeySet("card_" + strconv.FormatUint(cardInfo.Id,10), string(cacheData))
}

func (ms *cardService) cacheDelCard(cardInfo *CardInfo) *response.FWError {
	return cachehelper.KeyDel("card_" + strconv.FormatUint(cardInfo.Id,10))
}

// cache function end
// db function start

func (ms *cardService) dbGetCardByCond(cond *CardInfo) (*response.FWError, *CardInfo) {
	cardInfo := &CardInfo{}
	err := dbhelper.GetDataByCond(cond, cardInfo)
	if err == constant.ErrDBNoRecord {
		return err, cardInfo
	}
	return nil, cardInfo
}

func (ms *cardService) dbCreateCard(cardInfo *CardInfo) (*response.FWError, *CardInfo) {
	err, data := dbhelper.CreateData(cardInfo)
	cardInfo = data.(*CardInfo)
	if err != nil {
		return constant.ErrDb, nil
	}
	return nil, cardInfo
}
func (ms *cardService) dbUpdateCard(cardInfo *CardInfo) (*response.FWError, int64) {
	return dbhelper.UpdateData(cardInfo)
}
func (ms *cardService) dbGetPagedCardByCond(cond *CardInfo, data *[]*CardInfo, pageNo int64, pageSize int64) (*response.FWError, int64) {
	return dbhelper.GetPagedDataByCond(cond, data, pageNo, pageSize)
}


// db function end

// service function start
func (ms *cardService) GetCardByCond(cond *CardInfo) (*response.FWError, *CardInfo) {
	// get cache
	cardInfo := &CardInfo{}
	err, cardInfo := ms.cacheGetCard(cond)
	if err == nil {
		return nil, cardInfo
	}
	if err == constant.ErrCacheNotExist {
		// get db
		err, cardInfo = ms.dbGetCardByCond(cond)
		if err == constant.ErrCacheNotExist {
			return err, nil
		}
		if err != nil {
			return err, nil
		} else {
			return nil, cardInfo
		}
	} else {
		return err, nil
	}
}

func (ms *cardService) CreateCard(cardInfo *CardInfo) (*response.FWError, *CardInfo) {
	cardInfo.Ctime = uint32(time.Now().Unix())
	cardInfo.Mtime = uint32(time.Now().Unix())
	err, cardInfo := ms.dbCreateCard(cardInfo)
	if err != nil {
		openlog.Error(err.String() + " db create card failed.")
		return err, nil
	}
	err = ms.cacheSetCard(cardInfo)
	if err != nil {
		openlog.Error(err.String() + " cache set card failed.")
	}
	return nil, cardInfo
}

func (ms *cardService) UpdateCard(cardInfo *CardInfo) (*response.FWError, int64) {
	cardInfo.Mtime = uint32(time.Now().Unix())
	err, row := ms.dbUpdateCard(cardInfo)
	if err != nil {
		return err, row
	}

	err = cachehelper.KeyDel("card_" + strconv.FormatUint(cardInfo.Id,10))
	if err != nil {
		openlog.Error("card_" + strconv.FormatUint(cardInfo.Id,10) + " cache del failed.")
	}
	return nil, row
}
func (ms *cardService) GetPagedCardByCond(cond *CardInfo, data *[]*CardInfo, pageNo int64, pageSize int64) (*response.FWError, int64) {
	return ms.dbGetPagedCardByCond(cond, data, pageNo, pageSize)
}

// service function end
