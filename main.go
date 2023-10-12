package main

import (
	"fmt"
	"slowcom-fujica-parking-sdk/app/common"
	"slowcom-fujica-parking-sdk/app/parking/request"
	"slowcom-fujica-parking-sdk/config"
)

func main() {
	fmt.Println("停车sdk启动")
	fsClient := &config.FsHttpClient{
		BaseUrl: common.BaseUrl,
		Token:   common.Token,
	}
	req := request.ParkingRequest{FsClient: fsClient}
	req.GetParkingById(common.ParkingId)
}
