package request

import (
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
	mp := map[string]string{"parkId": parkId, "sign": common.InterfaceSign}
	response, err = s.FsClient.PostJson("fujica/api/v1/park/queryParkInfo", mp)
	fmt.Printf("%+v", response)
	return
}

// GetParkingCards 月卡-查询月卡列表
func (s *ParkingRequest) GetParkingCards(param *entity.ParkingCardSearchParam) (response *config.FsResponse, err error) {
	param.Sign = common.InterfaceSign
	response, err = s.FsClient.PostJson("fujica/api/v1/card/queryCardList", param)
	fmt.Printf("%+v", response)
	return
}

// GetParkingCardRules 月卡-查询月卡套餐
func (s *ParkingRequest) GetParkingCardRules(parkId string) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId}
	mp["sign"] = common.GetSign(mp)
	response, err = s.FsClient.PostJson("fujica/api/v1/card/cardType/query", mp)
	fmt.Printf("result: %+v", response)
	return
}

// GetCarOwnerInfo 月卡-车主信息查询 & 根据车牌手机号查询月卡
// cardTypeName 套餐名称, plateNum 车牌
func (s *ParkingRequest) GetCarOwnerInfo(parkId string, plateNum string, cardTypeName string, phoneNumber string, userName string) (response *config.FsResponse, err error) {
	mp := map[string]string{"parkId": parkId, "sign": common.InterfaceSign}
	mp["plateNum"] = plateNum
	mp["cardTypeName"] = cardTypeName
	mp["phoneNumber"] = phoneNumber
	mp["userName"] = userName
	response, err = s.FsClient.PostJson("fujica/api/v1/card/userInfo/query", mp)
	fmt.Printf("%+v", response)
	return
}

// ParkingCardApply 月卡-月卡服务开通
func (s *ParkingRequest) ParkingCardApply(param *entity.ParkingCardApplyParam) (response *config.FsResponse, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = param.ParkId
	mp["userName"] = param.UserName
	mp["cardTypeId"] = param.CardTypeId
	mp["amountReceivable"] = param.AmountReceivable
	mp["cardCopies"] = param.CardCopies
	mp["phoneNumber"] = param.PhoneNumber
	//mp["houseInfo"] = param.HouseInfo
	//mp["isGroup"] = param.IsGroup
	//mp["numCar"] = param.NumCar
	//mp["parkingLot"] = param.ParkingLot
	//mp["remark"] = param.Remark
	param.Sign = common.GetSign(mp)
	response, err = s.FsClient.PostJson("fujica/api/v1/card/saveCard", param)
	fmt.Printf("%+v", response)
	return
}

// ParkingCardRenew 月卡-月卡续费
func (s *ParkingRequest) ParkingCardRenew(param *entity.ParkingCardRenewParam) (response *config.FsResponse, err error) {
	param.Sign = common.InterfaceSign
	response, err = s.FsClient.PostJson("fujica/api/v1/card/modifyDate", param)
	fmt.Printf("%+v", response)
	return
}

// ParkingCardCancel 月卡-月卡注销
// cardId 月卡id-必须, cause 注销原因-必须
func (s *ParkingRequest) ParkingCardCancel(parkId string, cardId int64, cause string) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId, "sign": common.InterfaceSign}
	mp["cardId"] = cardId
	mp["cause"] = cause
	response, err = s.FsClient.PostJson("fujica/api/v1/card/cancelOneCard", mp)
	fmt.Printf("%+v", response)
	return
}

// CarBillRecords 计费-车辆计费（查询车辆停车费）
func (s *ParkingRequest) CarBillRecords(param *entity.CarBillRecordsParam) (response *config.FsResponse, err error) {
	param.Sign = common.InterfaceSign
	response, err = s.FsClient.PostJson("fujica/api/v1/payrecord/charge", param)
	fmt.Printf("%+v", response)
	return
}

// CarPayRecords 支付-查询已缴费记录
func (s *ParkingRequest) CarPayRecords(param *entity.CarPayRecordsParam) (response *config.FsResponse, err error) {
	param.Sign = common.InterfaceSign
	response, err = s.FsClient.PostJson("fujica/api/v1/charge/record/query", param)
	fmt.Printf("%+v", response)
	return
}

// ParkingSpaceReserve 车位-车场级预约
func (s *ParkingRequest) ParkingSpaceReserve(param *entity.ParkingSpaceReserveParam) (response *config.FsResponse, err error) {
	param.Sign = common.InterfaceSign
	response, err = s.FsClient.PostJson("fujica/api/v1/reserves/saveReserve", param)
	fmt.Printf("%+v", response)
	return
}

// ParkingSpaceCancelReserve 车位-取消预订
// id 预定订单ID-必须
func (s *ParkingRequest) ParkingSpaceCancelReserve(parkId string, id string) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId, "sign": common.InterfaceSign, "id": id}
	response, err = s.FsClient.PostJson("fujica/api/v1/reserves/cancelReserve", mp)
	fmt.Printf("%+v", response)
	return
}

// ParkingInRecords 进出场-入场信息查询
func (s *ParkingRequest) ParkingInRecords(param *entity.ParkingInoutParam) (response *config.FsResponse, err error) {
	param.Sign = common.InterfaceSign
	response, err = s.FsClient.PostJson("fujica/api/v1/inout/queryPageParkin", param)
	fmt.Printf("%+v", response)
	return
}

// ParkingOutRecords 进出场-出场信息查询
func (s *ParkingRequest) ParkingOutRecords(param *entity.ParkingInoutParam) (response *config.FsResponse, err error) {
	param.Sign = common.InterfaceSign
	response, err = s.FsClient.PostJson("fujica/api/v1/inout/queryPageParkout", param)
	fmt.Printf("%+v", response)
	return
}

// GetLaneByparkId 进出场-根据车场id获取车道信息
func (s *ParkingRequest) GetLaneByparkId(parkId string) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId, "sign": common.InterfaceSign}
	response, err = s.FsClient.PostJson("fujica/api/v1/park/seltlandebyparkid", mp)
	fmt.Printf("%+v", response)
	return
}

// GrantCouponToCar 优惠券-按车牌号发放优惠券
func (s *ParkingRequest) GrantCouponToCar(param *entity.GrantCouponParam) (response *config.FsResponse, err error) {
	param.Sign = common.InterfaceSign
	response, err = s.FsClient.PostJson("fujica/api/v1/coupon/addCouponNotUsed", param)
	fmt.Printf("%+v", response)
	return
}

// GetMerchantCoupons 优惠券-查询商户优惠券列表
func (s *ParkingRequest) GetMerchantCoupons(parkId string, shopId int64) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId, "sign": common.InterfaceSign, "shopId": shopId}
	response, err = s.FsClient.PostJson("fujica/api/v1/coupon/distribution/getCarCouponList", mp)
	fmt.Printf("%+v", response)
	return
}
