package main

import (
	"fmt"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/app/parking/request"
	cardEntity "github.com/luojinqiang/slowcom-fujica-parking-sdk/app/parking_card/entity"
	cardRequest "github.com/luojinqiang/slowcom-fujica-parking-sdk/app/parking_card/request"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/config"
)

func main() {
	fmt.Println("停车sdk启动")
	parkId := ""
	fsClient := &config.FsHttpClient{
		BaseUrl: "",
		Token:   "",
	}
	req := request.ParkingRequest{FsClient: fsClient}
	req.GetParkingById(parkId)
	//cardReq := cardRequest.ParkingCardRequest{FsClient: fsClient}
	// 1、获取月卡套餐
	//cardReq.GetParkingCardRules(parkId)
	// 2、办卡
	//cardApplyTest(cardReq)
	// 3、续费
	//cardRenewTest(req)
	// 注销月卡
	//cardReq.ParkingCardCancel(common.ParkId, 213355, "测试月卡注销")
	// 4、获取月卡列表
	//cardReq.GetParkingCards(&cardEntity.ParkingCardSearchParam{
	//	ParkId:   common.ParkId,
	//	IsCommon: "0",
	//	Status:   1,
	//	Current:  1,
	//	Size:     10,
	//})
	// 5、查个人信息及月卡
	// req.GetCarOwnerInfo(common.ParkId, "", "", "15076453821", "")
	// 6、查询车辆停车费
	//req.CarBillRecords(&entity.CarBillRecordsParam{
	//	ParkId:        parkId,
	//	LicenseNumber: "粤B98212",
	//	CouponIds:     []string{"1573280"},
	//})
	// 7、查询缴费记录
	//req.CarPayRecords(&entity.CarPayRecordsParam{
	//	ParkId:        common.ParkId,
	//	LicenseNumber: "粤B98212",
	//	CarType:       1,
	//	LicenseType:   1,
	//})
	// 8、上传支付结果
	//req.CarPayAdd(&entity.CarPayAddParam{
	//	ParkId:      parkId,
	//	OrderId:     "AB115353",
	//	PayStatus:   "2",
	//	PayWay:      1,
	//	SuccessTime: "2023-10-17 16:07:02",
	//})
	// 9、车位预约
	//req.ParkingSpaceReserve(&entity.ParkingSpaceReserveParam{
	//	ParkId:         common.ParkId,
	//	LicenseNumber:  "粤B12343",
	//	ReserveInTime:  gtime.Now(),
	//	ReserveOutTime: gtime.Now().Add(time.Duration(5) * time.Hour),
	//})
	// 10、车位预约取消
	//req.ParkingSpaceCancelReserve(common.ParkId, "210")
	// 11、出入场信息查询
	//req.ParkingInOutRecords(&entity.ParkingInoutParam{
	//	InOut:  2,
	//	ParkId: common.ParkId,
	//	//ChargeType:    1,
	//	//ModelName:     "22",
	//	//ExceptionType: 1,
	//	//LicenseNumber: "粤B10232",
	//	Current: 1,
	//	Size:    10,
	//})
	// 12、根据车场id获取车道信息
	//req.GetLaneByParkId(common.ParkId)
	// 13、发放优惠券
	//req.GrantCouponToCar(&entity.GrantCouponParam{
	//	Parkid:        parkId,
	//	Couponname:    "充电减免2小时",
	//	Coupontype:    4,
	//	Couponrule:    2,
	//	LicenseNumber: "粤B93412",
	//	Expday:        7,
	//	Num:           2,
	//	Source:        1,
	//	Userule:       0,
	//})
	// 14、获取商户优惠券
	//req.GetMerchantCoupons(common.ParkId, 1)
}

func cardRenewTest(req cardRequest.ParkingCardRequest) {
	req.ParkingCardRenew(&cardEntity.ParkingCardRenewParam{
		ParkId:           "",
		CardId:           212537,
		AmountReceivable: 10000,
		CardCopies:       1,
		StartTime:        "2023-11-19",
		EndTime:          "2023-12-19",
		OldExpireTime:    "2023-11-18",
	})
}

func cardApplyTest(req cardRequest.ParkingCardRequest) {
	req.ParkingCardApply(&cardEntity.ParkingCardApplyParam{
		ParkId:           "",
		UserName:         "王凯新",
		CardTypeId:       1101,
		AmountReceivable: 10000,
		CardCopies:       1,
		StartTime:        "2023-10-18",
		EndTime:          "2023-11-18",
		NewEndTime:       "2023-11-18",
		PhoneNumber:      "15076453821",
		CardPlateRels: []*cardEntity.CardPlateRelsModel{
			{
				Brand:         "宝马",
				LicenseNumber: "粤B98212",
				LicenseType:   1,
				PlateColor:    "红",
				PlateType:     "SUV",
			},
		},
		Remark:     "测试",
		HouseInfo:  "1栋",
		IsGroup:    0,
		NumCar:     1,
		ParkingLot: 1,
	})
}
