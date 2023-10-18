package main

import (
	"fmt"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/app/common"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/app/parking/request"
	cardEntity "github.com/luojinqiang/slowcom-fujica-parking-sdk/app/parking_card/entity"
	cardRequest "github.com/luojinqiang/slowcom-fujica-parking-sdk/app/parking_card/request"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/config"
)

func main() {
	fmt.Println("停车sdk启动")
	fsClient := &config.FsHttpClient{}
	//cardReq := cardRequest.ParkingCardRequest{FsClient: fsClient}
	req := request.ParkingRequest{FsClient: fsClient}
	req.GetParkingById(common.ParkId)
	// 1、获取月卡套餐
	//cardReq.GetParkingCardRules(common.ParkId)
	// 2、办卡
	//cardApplyTest(cardReq)
	// 3、续费
	//cardRenewTest(req)
	// 4、获取月卡列表
	//req.GetParkingCards(&entity.ParkingCardSearchParam{
	//	ParkId:   common.ParkId,
	//	IsCommon: "0",
	//	Status:   6,
	//	Current:  1,
	//	Size:     10,
	//})
	// 5、查个人信息及月卡
	// req.GetCarOwnerInfo(common.ParkId, "", "", "15076453821", "")
	// 6、查询车辆停车费
	//req.CarBillRecords(&entity.CarBillRecordsParam{
	//	ParkId:        common.ParkId,
	//	LicenseNumber: "粤B98212",
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
	//	ParkId:      common.ParkId,
	//	OrderId:     "AB115353",
	//	PayStatus:   "2",
	//	PayWay:      1,
	//	SuccessTime: "2023-10-17 16:07:02",
	//})
	// 9、车位预约
	//req.ParkingSpaceReserve(&entity.ParkingSpaceReserveParam{
	//	ParkId:         common.ParkId,
	//	LicenseNumber:  "粤B12323",
	//	ReserveInTime:  gtime.Now(),
	//	ReserveOutTime: gtime.Now().Add(time.Duration(5) * time.Hour),
	//})
	// 10、车位预约取消
	//req.ParkingSpaceCancelReserve(common.ParkId, "210")
	// 11、入场信息查询
	//req.ParkingInRecords(&entity.ParkingInoutParam{
	//	ParkId: common.ParkId,
	//	//ChargeType:    1,
	//	//ModelName:     "22",
	//	//ExceptionType: 1,
	//	//LicenseNumber: "粤B10232",
	//	Current: 1,
	//	Size:    10,
	//})
	// 12、出场信息查询
	//req.ParkingOutRecords(&entity.ParkingInoutParam{
	//	ParkId: common.ParkId,
	//	//ChargeType:    1,
	//	//ModelName:     "22",
	//	//ExceptionType: 1,
	//	//LicenseNumber: "粤B10232",
	//	Current:       1,
	//	Size:          10,
	//})
	// 13、根据车场id获取车道信息
	//req.GetLaneByParkId(common.ParkId)
	// 14、发放优惠券
	//req.GrantCouponToCar(&entity.GrantCouponParam{
	//	Parkid:        common.ParkId,
	//	Couponname:    "充电减免2小时",
	//	Coupontype:    4,
	//	Couponrule:    2,
	//	LicenseNumber: "粤B98212",
	//	Expday:        7,
	//	Num:           2,
	//	Source:        1,
	//	Userule:       0,
	//})
	// 15、获取商户优惠券
	//req.GetMerchantCoupons(common.ParkId, 1)
}

func cardRenewTest(req cardRequest.ParkingCardRequest) {
	req.ParkingCardRenew(&cardEntity.ParkingCardRenewParam{
		ParkId:           common.ParkId,
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
		ParkId:           common.ParkId,
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
