package request

import (
	"encoding/json"
	"fmt"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/app/common"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/app/parking/entity"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/config"
)

type ParkingRequest struct {
	FsClient *config.FsHttpClient
}

// GetParkingById 车场-根据id获取车场信息
func (s *ParkingRequest) GetParkingById(parkId string) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId}
	mp["sign"] = common.GetSign(mp)
	response, err = s.FsClient.PostJson("park/queryParkInfo", mp)
	return
}

// CarBillRecords 计费-车辆计费（只能查询单个车辆停车费）
func (s *ParkingRequest) CarBillRecords(param *entity.CarBillRecordsParam) (data *entity.CarBillRecordsModel, err error) {
	mp := make(map[string]interface{})
	mp["parkId"] = param.ParkId
	mp["licenseNumber"] = param.LicenseNumber
	mp["licenseType"] = param.LicenseType
	param.Sign = common.GetSign(mp)
	response, err := s.FsClient.PostJson("payrecord/charge", param)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &data)
	}
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
	return
}

// ParkingSpaceReserve 车位-车场级预约
func (s *ParkingRequest) ParkingSpaceReserve(param *entity.ParkingSpaceReserveParam) (data *entity.ParkingSpaceReserveResult, err error) {
	mp := map[string]interface{}{"parkId": param.ParkId}
	param.Sign = common.GetSign(mp)
	response, err := s.FsClient.PostJson("reserves/saveReserve", param)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &data)
	}
	return
}

// ParkingSpaceCancelReserve 车位-取消预订
// id 预定订单ID-必须
func (s *ParkingRequest) ParkingSpaceCancelReserve(parkId string, id string) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId, "id": id}
	mp["sign"] = common.GetSign(mp)
	response, err = s.FsClient.PostJson("reserves/cancelReserve", mp)
	return
}

// ParkingInRecords 进场记录查询
func (s *ParkingRequest) ParkingInRecords(param *entity.ParkingInParam) (data *entity.ParkingInResult, err error) {
	mp := common.StructToMap(param)
	signMap := make(map[string]interface{})
	signMap["parkId"] = param.ParkId
	signMap["chargeType"] = param.ChargeType
	signMap["exceptionType"] = param.ExceptionType
	if param.ModelName != "" {
		signMap["modelName"] = param.ModelName
	}
	if param.LicenseNumber != "" {
		signMap["licenseNumber"] = param.LicenseNumber
	}
	if param.MaxStayTime > 0 {
		signMap["maxStayTime"] = param.MaxStayTime
	} else {
		delete(mp, "maxStayTime")
	}
	if param.MinStayTime > 0 {
		signMap["minStayTime"] = param.MinStayTime
	} else {
		delete(mp, "minStayTime")
	}
	mp["sign"] = common.GetSign(signMap)
	url := "inout/queryPageParkin"
	response, err := s.FsClient.PostJson(url, mp)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &data)
	}
	return
}

// ParkingOutRecords 出场记录查询
func (s *ParkingRequest) ParkingOutRecords(param *entity.ParkingOutParam) (data *entity.ParkingOutResult, err error) {
	mp := common.StructToMap(param)
	signMap := make(map[string]interface{})
	signMap["parkId"] = param.ParkId
	signMap["chargeType"] = param.ChargeType
	if param.ExceptionType > 0 {
		signMap["exceptionType"] = param.ExceptionType
	} else {
		delete(mp, "exceptionType")
	}
	if param.ModelName != "" {
		signMap["modelName"] = param.ModelName
	}
	if param.LicenseNumber != "" {
		signMap["licenseNumber"] = param.LicenseNumber
	}
	if param.MaxStayTime > 0 {
		signMap["maxStayTime"] = param.MaxStayTime
	} else {
		delete(mp, "maxStayTime")
	}
	if param.MinStayTime > 0 {
		signMap["minStayTime"] = param.MinStayTime
	} else {
		delete(mp, "minStayTime")
	}
	if param.BeginOutTime == `` {
		delete(mp, "beginOutTime")
	}
	if param.BeginEnterTime == `` {
		delete(mp, "beginEnterTime")
	}
	if param.EndOutTime == `` {
		delete(mp, "endOutTime")
	}
	if param.EndEnterTime == `` {
		delete(mp, "endEnterTime")
	}
	mp["sign"] = common.GetSign(signMap)
	url := "inout/queryPageParkout"
	response, err := s.FsClient.PostJson(url, mp)
	fmt.Println("response = ", response)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &data)
	}
	return
}

// GetLaneByParkId 进出场-根据车场id获取车道信息
func (s *ParkingRequest) GetLaneByParkId(parkId string) (list []*entity.ParkingLaneModel, err error) {
	mp := map[string]interface{}{"parkId": parkId}
	mp["sign"] = common.GetSign(mp)
	response, err := s.FsClient.PostJson("park/seltlandebyparkid", mp)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &list)
	}
	return
}

// GrantCouponToCar 优惠券-按车牌号发放优惠券
func (s *ParkingRequest) GrantCouponToCar(param *entity.GrantCouponParam) (response *config.FsResponse, err error) {
	mp := make(map[string]interface{})
	mp["parkid"] = param.Parkid
	param.Sign = common.GetSign(mp)
	response, err = s.FsClient.PostJson("coupon/addCouponNotUsed", param)
	return
}

// GetMerchantCoupons 优惠券-查询商户优惠券列表
func (s *ParkingRequest) GetMerchantCoupons(parkId string, shopId int64) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkId, "shopId": shopId}
	mp["sign"] = common.GetSign(mp)
	response, err = s.FsClient.PostJson("coupon/distribution/getCarCouponList", mp)
	return
}

// UnlicensedInOut 进出场-无牌车扫码出、入场
func (s *ParkingRequest) UnlicensedInOut(param *entity.UnlicensedInOutParam) (data *entity.UnlicensedInOutModel, err error) {
	mp := make(map[string]interface{})
	mp["parkid"] = param.Parkid
	mp["openId"] = param.OpenId
	mp["entranceid"] = param.Entranceid
	param.Sign = common.GetSign(mp)
	url := "parkin/noLicenseParkin"
	if param.InOut == 2 {
		url = "parkout/scanOut"
	}
	response, err := s.FsClient.PostJson(url, param)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &data)
	}
	return
}

// GetCarNoByPressure 进出场-压地感查询，根据车道码查询返回车牌
// laneno 车道编码-必须
func (s *ParkingRequest) GetCarNoByPressure(parkId string, laneno string) (data *entity.CarNoPressureModel, err error) {
	mp := map[string]interface{}{"parkId": parkId}
	mp["laneno"] = laneno
	mp["sign"] = common.GetSign(mp)
	response, err := s.FsClient.PostJson("payment/getParkOutInfo", mp)
	if err == nil && response != nil {
		j, _ := json.Marshal(&response.Data)
		json.Unmarshal(j, &data)
	}
	return
}
