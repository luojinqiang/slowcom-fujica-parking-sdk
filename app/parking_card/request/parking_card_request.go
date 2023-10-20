package request

import (
	"encoding/json"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/app/common"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/app/parking_card/entity"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/config"
)

type ParkingCardRequest struct {
	FsClient *config.FsHttpClient
}

// GetParkingCards 月卡-查询月卡列表
func (s *ParkingCardRequest) GetParkingCards(param *entity.ParkingCardSearchParam) (data *entity.ParkingCardResult, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = param.ParkId
	mp["isCommon"] = param.IsCommon
	mp["status"] = param.Status
	// 不必须但要加签的参数，不为空需要加签，为空也加签会签名不正确
	if param.CardTypeName != "" {
		mp["cardTypeName"] = param.CardTypeName
	}
	if param.HouseInfo != "" {
		mp["houseInfo"] = param.HouseInfo
	}
	if param.LicenseNumber != "" {
		mp["licenseNumber"] = param.LicenseNumber
	}
	if param.PhoneNumber != "" {
		mp["phoneNumber"] = param.PhoneNumber
	}
	if param.UserName != "" {
		mp["userName"] = param.UserName
	}
	param.Sign = common.GetSign(mp)
	response, err := s.FsClient.PostJson("card/queryCardList", param)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &data)
	}
	return
}

// GetParkingCardRules 月卡-查询月卡套餐
func (s *ParkingCardRequest) GetParkingCardRules(parkId string) (list []*entity.ParkingCardRuleModel, err error) {
	mp := map[string]interface{}{"parkId": parkId}
	mp["sign"] = common.GetSign(mp)
	response, err := s.FsClient.PostJson("card/cardType/query", mp)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &list)
	}
	return
}

// GetCarOwnerInfo 月卡-车主信息查询 & 根据车牌手机号查询月卡
// cardTypeName 套餐名称, plateNum 车牌
func (s *ParkingCardRequest) GetCarOwnerInfo(parkId string, plateNum string, cardTypeName string, phoneNumber string, userName string) (data *entity.CarOwnerInfoModel, err error) {
	mp := map[string]interface{}{"parkId": parkId}
	if userName != "" {
		mp["userName"] = userName
	}
	if plateNum != "" {
		mp["plateNum"] = plateNum
	}
	if cardTypeName != "" {
		mp["cardTypeName"] = cardTypeName
	}
	if phoneNumber != "" {
		mp["phoneNumber"] = phoneNumber
	}
	mp["sign"] = common.GetSign(mp)
	response, err := s.FsClient.PostJson("card/userInfo/query", mp)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &data)
	}
	return
}

// ParkingCardApply 月卡-月卡服务开通
func (s *ParkingCardRequest) ParkingCardApply(param *entity.ParkingCardApplyParam) (data *entity.ParkingCardModel, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = param.ParkId
	mp["userName"] = param.UserName
	mp["cardTypeId"] = param.CardTypeId
	mp["amountReceivable"] = param.AmountReceivable
	mp["cardCopies"] = param.CardCopies
	mp["phoneNumber"] = param.PhoneNumber
	mp["isGroup"] = param.IsGroup
	mp["numCar"] = param.NumCar
	mp["parkingLot"] = param.ParkingLot
	// 不必须但要加签的参数，不为空需要加签，为空也加签会签名不正确
	if param.HouseInfo != "" {
		mp["houseInfo"] = param.HouseInfo
	}
	if param.Remark != "" {
		mp["remark"] = param.Remark
	}
	param.Sign = common.GetSign(mp)
	response, err := s.FsClient.PostJson("card/saveCard", param)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &data)
	}
	return
}

// ParkingCardRenew 月卡-月卡续费
func (s *ParkingCardRequest) ParkingCardRenew(param *entity.ParkingCardRenewParam) (response *config.FsResponse, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = param.ParkId
	mp["cardId"] = param.CardId
	mp["amountReceivable"] = param.AmountReceivable
	mp["cardCopies"] = param.CardCopies
	param.Sign = common.GetSign(mp)
	response, err = s.FsClient.PostJson("card/modifyDate", param)
	return
}

// ParkingCardCancel 月卡-月卡注销
// cardId 月卡id-必须, cause 注销原因-必须
func (s *ParkingCardRequest) ParkingCardCancel(parkId string, cardId int64, cause string) (response *config.FsResponse, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = parkId
	mp["cardId"] = cardId
	mp["cause"] = cause
	mp["sign"] = common.GetSign(mp)
	response, err = s.FsClient.PostJson("card/cancelOneCard", mp)
	return
}
