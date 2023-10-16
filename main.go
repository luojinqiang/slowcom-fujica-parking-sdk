package main

import (
	"fmt"
	"slowcom-fujica-parking-sdk/app/common"
	"slowcom-fujica-parking-sdk/app/parking/entity"
	"slowcom-fujica-parking-sdk/app/parking/request"
	"slowcom-fujica-parking-sdk/config"
)

func main() {
	fmt.Println("停车sdk启动")
	fsClient := &config.FsHttpClient{}
	req := request.ParkingRequest{FsClient: fsClient}
	//req.GetParkingCardRules(common.ParkId)
	req.ParkingCardApply(&entity.ParkingCardApplyParam{
		ParkId:           common.ParkId,
		UserName:         "李四",
		CardTypeId:       1101,
		AmountReceivable: 10000,
		CardCopies:       1,
		StartTime:        "2023-10-16",
		EndTime:          "2023-11-16",
		NewEndTime:       "2023-11-16",
		PhoneNumber:      "15096782734",
		Sign:             "abcd",
		CardPlateRels: []*entity.CardPlateRelsModel{
			{
				Brand:         "宝马",
				LicenseNumber: "赣AFP273",
				LicenseType:   1,
				PlateColor:    "红",
				PlateType:     "SUV",
			},
		},
	})
}
