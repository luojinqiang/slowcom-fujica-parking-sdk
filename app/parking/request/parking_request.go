package request

import (
	"encoding/json"
	"fmt"
	"slowcom-fujica-parking-sdk/app/common"
	"slowcom-fujica-parking-sdk/app/parking/entity"
	"slowcom-fujica-parking-sdk/config"
)

type ParkingRequest struct {
	FsClient *config.FsHttpClient
}

// GetParkingById 车场-根据id获取车场信息
func (s *ParkingRequest) GetParkingById(parkId string) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId}
	mp["sign"] = common.GetSign(mp)
	response, err = s.FsClient.PostJson("park/queryParkInfo", mp)
	fmt.Printf("%+v", response)
	return
}

// GetParkingCards 月卡-查询月卡列表
func (s *ParkingRequest) GetParkingCards(param *entity.ParkingCardSearchParam) (data *entity.ParkingCardResult, err error) {
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
	fmt.Printf("%+v", response)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &data)
	}
	return
}

// GetParkingCardRules 月卡-查询月卡套餐
func (s *ParkingRequest) GetParkingCardRules(parkId string) (list []*entity.ParkingCardRuleModel, err error) {
	mp := map[string]interface{}{"parkId": parkId}
	mp["sign"] = common.GetSign(mp)
	response, err := s.FsClient.PostJson("card/cardType/query", mp)
	fmt.Printf("result: %+v", response)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &list)
	}
	return
}

// GetCarOwnerInfo 月卡-车主信息查询 & 根据车牌手机号查询月卡
// cardTypeName 套餐名称, plateNum 车牌
func (s *ParkingRequest) GetCarOwnerInfo(parkId string, plateNum string, cardTypeName string, phoneNumber string, userName string) (data *entity.CarOwnerInfoModel, err error) {
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
	fmt.Printf("%+v", response)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &data)
	}
	return
}

// ParkingCardApply 月卡-月卡服务开通
func (s *ParkingRequest) ParkingCardApply(param *entity.ParkingCardApplyParam) (data *entity.ParkingCardModel, err error) {
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
	fmt.Printf("%+v", response)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &data)
	}
	return
}

// ParkingCardRenew 月卡-月卡续费
func (s *ParkingRequest) ParkingCardRenew(param *entity.ParkingCardRenewParam) (response *config.FsResponse, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = param.ParkId
	mp["cardId"] = param.CardId
	mp["amountReceivable"] = param.AmountReceivable
	mp["cardCopies"] = param.CardCopies
	param.Sign = common.GetSign(mp)
	response, err = s.FsClient.PostJson("card/modifyDate", param)
	fmt.Printf("%+v", response)
	return
}

// ParkingCardCancel 月卡-月卡注销
// cardId 月卡id-必须, cause 注销原因-必须
func (s *ParkingRequest) ParkingCardCancel(parkId string, cardId int64, cause string) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId, "sign": common.InterfaceSign}
	mp["cardId"] = cardId
	mp["cause"] = cause
	response, err = s.FsClient.PostJson("card/cancelOneCard", mp)
	fmt.Printf("%+v", response)
	return
}

// CarBillRecords 计费-车辆计费（只能查询单个车辆停车费）
func (s *ParkingRequest) CarBillRecords(param *entity.CarBillRecordsParam) (response *config.FsResponse, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = param.ParkId
	mp["licenseNumber"] = param.LicenseNumber
	mp["licenseType"] = param.LicenseType
	param.Sign = common.GetSign(mp)
	response, err = s.FsClient.PostJson("payrecord/charge", param)
	fmt.Printf("%+v", response)
	return
}

// CarPayAdd 支付-上传支付结果
func (s *ParkingRequest) CarPayAdd(param *entity.CarPayAddParam) (response *config.FsResponse, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = param.ParkId
	mp["orderId"] = param.OrderId
	mp["payStatus"] = param.PayStatus
	mp["payWay"] = param.PayWay
	param.Sign = common.GetSign(mp)
	response, err = s.FsClient.PostJson("payment/up/payResult", param)
	fmt.Printf("%+v", response)
	return
}

// CarPayRecords 支付-查询已缴费记录（只能查询单个车辆缴费记录）
func (s *ParkingRequest) CarPayRecords(param *entity.CarPayRecordsParam) (response *config.FsResponse, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = param.ParkId
	mp["carType"] = param.CarType
	mp["licenseNumber"] = param.LicenseNumber
	mp["licenseType"] = param.LicenseType
	param.Sign = common.GetSign(mp)
	response, err = s.FsClient.PostJson("charge/record/query", param)
	fmt.Printf("%+v", response)
	return
}

// ParkingSpaceReserve 车位-车场级预约
func (s *ParkingRequest) ParkingSpaceReserve(param *entity.ParkingSpaceReserveParam) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": param.ParkId}
	param.Sign = common.GetSign(mp)
	response, err = s.FsClient.PostJson("reserves/saveReserve", param)
	fmt.Printf("%+v", response)
	return
}

// ParkingSpaceCancelReserve 车位-取消预订
// id 预定订单ID-必须
func (s *ParkingRequest) ParkingSpaceCancelReserve(parkId string, id string) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId, "id": id}
	mp["sign"] = common.GetSign(mp)
	response, err = s.FsClient.PostJson("reserves/cancelReserve", mp)
	fmt.Printf("%+v", response)
	return
}

// ParkingInRecords 进出场-入场信息查询
func (s *ParkingRequest) ParkingInRecords(param *entity.ParkingInoutParam) (response *config.FsResponse, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = param.ParkId
	mp["chargeType"] = param.ChargeType
	mp["exceptionType"] = param.ExceptionType
	if param.ModelName != "" {
		mp["modelName"] = param.ModelName
	}
	if param.LicenseNumber != "" {
		mp["licenseNumber"] = param.LicenseNumber
	}
	param.Sign = common.GetSign(mp)
	response, err = s.FsClient.PostJson("inout/queryPageParkin", param)
	fmt.Printf("%+v", response)
	return
}

// ParkingOutRecords 进出场-出场信息查询
func (s *ParkingRequest) ParkingOutRecords(param *entity.ParkingInoutParam) (response *config.FsResponse, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = param.ParkId
	mp["chargeType"] = param.ChargeType
	mp["exceptionType"] = param.ExceptionType
	if param.ModelName != "" {
		mp["modelName"] = param.ModelName
	}
	if param.LicenseNumber != "" {
		mp["licenseNumber"] = param.LicenseNumber
	}
	param.Sign = common.GetSign(mp)
	response, err = s.FsClient.PostJson("inout/queryPageParkout", param)
	fmt.Printf("%+v", response)
	return
}

// GetLaneByParkId 进出场-根据车场id获取车道信息
func (s *ParkingRequest) GetLaneByParkId(parkId string) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId}
	mp["sign"] = common.GetSign(mp)
	response, err = s.FsClient.PostJson("park/seltlandebyparkid", mp)
	fmt.Printf("%+v", response)
	return
}

// GrantCouponToCar 优惠券-按车牌号发放优惠券
func (s *ParkingRequest) GrantCouponToCar(param *entity.GrantCouponParam) (response *config.FsResponse, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = param.ParkId
	//mp["couponname"] = param.Couponname
	//mp["coupontype"] = param.Coupontype
	//mp["couponrule"] = param.Couponrule
	//mp["licensenumber"] = param.LicenseNumber
	//mp["expday"] = param.Expday
	//mp["num"] = param.Num
	//mp["source"] = param.Source
	//mp["userule"] = param.Userule
	param.Sign = common.GetSign(mp)
	response, err = s.FsClient.PostJson("coupon/addCouponNotUsed", param)
	fmt.Printf("%+v", response)
	return
}

// GetMerchantCoupons 优惠券-查询商户优惠券列表
func (s *ParkingRequest) GetMerchantCoupons(parkId string, shopId int64) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId, "shopId": shopId}
	mp["sign"] = common.GetSign(mp)
	response, err = s.FsClient.PostJson("coupon/distribution/getCarCouponList", mp)
	fmt.Printf("%+v", response)
	return
}
