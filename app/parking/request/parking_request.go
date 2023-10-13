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
	//sign, timestamp := common.GetSign()
	//param := map[string]string{"appId": common.AppId, "appSecret": common.AppSecret, "mchId": common.MchId, "sign": sign, "timestamp": timestamp}
	//response, err = s.FsClient.TokenPost("fujica/api/v1/public/token", param)
	//if err != nil {
	//	fmt.Println("tokenErr", err)
	//}
	//fmt.Println("token", response.Data)

	sign := "J+KfVVHJZCtvWhxKswvpeuHk2YHS27Z7DiTLRGL4qf2MzWpZ2GoPv4I7D2CsKji8WvxTZxk0+8dfD4oE23YeRPxHnfce/I9D"
	//url := "fujica/api/v1/card/queryCardList"
	url := "fujica/api/v1/park/queryParkInfo"
	mp := make(map[string]string)
	mp["sign"] = sign
	mp["parkId"] = parkingId
	s.FsClient.Token = fmt.Sprintf("%v", common.Token)
	response, err = s.FsClient.PostJson(url, mp)
	fmt.Printf("%+v", response)
	return
}
