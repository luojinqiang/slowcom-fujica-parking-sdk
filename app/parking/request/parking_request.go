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

// GetParkingById 根据id获取车场信息
func (s *ParkingRequest) GetParkingById(parkingId string) (response *config.FsResponse, err error) {
	// response, err = s.FsClient.GetToken()
	mp := map[string]string{"parkId": parkingId, "sign": common.InterfaceSign}
	s.FsClient.Token = fmt.Sprintf("%v", common.Token)
	response, err = s.FsClient.PostJson("fujica/api/v1/park/queryParkInfo", mp)
	fmt.Printf("%+v", response)
	return
}

// GetParkingCard 查询停车卡列表
// isCommon 是否为集团通卡 0:否 1:是, current 分页-当前页码, size 分页-每页记录数
func (s *ParkingRequest) GetParkingCard(parkingId string, isCommon string, current int, size int) (response *config.FsResponse, err error) {
	mp := map[string]interface{}{"parkId": parkingId, "sign": common.InterfaceSign, "isCommon": isCommon}
	mp["current"] = current
	mp["size"] = size
	s.FsClient.Token = fmt.Sprintf("%v", common.Token)
	response, err = s.FsClient.PostJson("fujica/api/v1/card/queryCardList", mp)
	fmt.Printf("%+v", response)
	return
}

// GetParkingCardRule 查询停车卡套餐列表
func (s *ParkingRequest) GetParkingCardRule(parkingId string) (response *config.FsResponse, err error) {
	mp := map[string]string{"parkId": parkingId, "sign": common.InterfaceSign}
	s.FsClient.Token = fmt.Sprintf("%v", common.Token)
	response, err = s.FsClient.PostJson("fujica/api/v1/card/cardType/query", mp)
	fmt.Printf("%+v", response)
	return
}

// GetCarOwnerInfo 车主信息查询 & 根据车牌手机号查询月卡
// cardTypeName 套餐名称, plateNum 车牌
func (s *ParkingRequest) GetCarOwnerInfo(parkingId string, plateNum string, cardTypeName string, phoneNumber string, userName string) (response *config.FsResponse, err error) {
	mp := map[string]string{"parkId": parkingId, "sign": common.InterfaceSign}
	mp["plateNum"] = plateNum
	mp["cardTypeName"] = cardTypeName
	mp["phoneNumber"] = phoneNumber
	mp["userName"] = userName
	s.FsClient.Token = fmt.Sprintf("%v", common.Token)
	response, err = s.FsClient.PostJson("fujica/api/v1/card/userInfo/query", mp)
	fmt.Printf("%+v", response)
	return
}

// ApplyParkingCard 月卡服务开通
// cardTypeName 套餐名称, plateNum 车牌
func (s *ParkingRequest) ApplyParkingCard(param *entity.ParkingCardParam) (response *config.FsResponse, err error) {
	param.Sign = common.InterfaceSign
	s.FsClient.Token = fmt.Sprintf("%v", common.Token)
	response, err = s.FsClient.PostJson("fujica/api/v1/card/saveCard", param)
	fmt.Printf("%+v", response)
	return
}
