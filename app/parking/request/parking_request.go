package request

import (
	"fmt"
	"slowcom-fujica-parking-sdk/app/common"
	"slowcom-fujica-parking-sdk/config"
)

type ParkingRequest struct {
	FsClient *config.FsHttpClient
}

// GetParkingById 根据id获取车场信息
func (s *ParkingRequest) GetParkingById(parkingId string) (response *config.FsResponse, err error) {
	sign := common.GetSign()

	//param := map[string]string{"appId": common.AppId, "appSecret": common.AppSecret, "mchId": common.MchId, "sign": sign}
	//response, err = s.FsClient.TokenPost("/fujica/api/v1/public/token", param)
	//if err != nil {
	//	return
	//}

	mp := make(map[string]string)
	mp["sign"] = sign
	mp["parkId"] = parkingId
	response, err = s.FsClient.PostJson("/fujica/api/v1/park/queryParkInfo", mp)
	if err != nil {
		fmt.Println(err)
	}
	return
}
