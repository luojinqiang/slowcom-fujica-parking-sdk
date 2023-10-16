package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"slowcom-fujica-parking-sdk/app/common"
)

// ParkingEntity 停车场
type ParkingEntity struct {
	ID          int    `json:"ID"`          //车场编号
	Name        string `json:"Name"`        //名称
	LotCount    int    `json:"LotCount"`    //总车位数
	LotFree     int    `json:"LotFree"`     //剩余车位
	ParentNo    string `json:"ParentNo"`    //上级车场
	Remark      string `json:"Remark"`      //备注
	Gid         string `json:"Gid"`         //标识号
	Rid         string `json:"Rid"`         //集团标识号
	ParkingCode string `json:"ParkingCode"` //线上停车场编号
}

// ParkingCardApplyParam 月卡开通参数
type ParkingCardApplyParam struct {
	ParkId           string                `json:"parkId" dc:"车场id,必须"`
	UserName         string                `json:"userName"  dc:"用户名,必须"`
	CardTypeId       int64                 `json:"cardTypeId" dc:"月卡套餐id,必须"`
	StartTime        string                `json:"startTime"  dc:"有效期起始时间 格式:yyyy-MM-dd,必须"`
	EndTime          string                `json:"endTime" dc:"有效期结束时间 格式:yyyy-MM-dd,必须"`
	NewEndTime       string                `json:"newEndTime"  dc:"实际有效期截至时间 格式:yyyy-MM-dd,必须"`
	AmountReceivable int                   `json:"amountReceivable" dc:"应收金额 单位/分,必须"`
	CardCopies       int                   `json:"cardCopies"  dc:"购买份数 1-9999,必须"`
	CardPlateRels    []*CardPlateRelsModel `json:"cardPlateRels" dc:"车牌,必须"`
	HouseInfo        string                `json:"houseInfo"  dc:"住宅信息"`
	IsGroup          int                   `json:"isGroup" dc:"是否组车 0-否 1-是"`
	NumCar           int                   `json:"numCar"  dc:"绑定车辆数 最大3"`
	ParkList         []string              `json:"parkList" dc:"车场id列表"`
	ParkingLot       int                   `json:"parkingLot"  dc:"车位数"`
	PhoneNumber      string                `json:"phoneNumber"  dc:"联系电话"`
	Remark           string                `json:"remark"  dc:"备注"`
	Sign             string                `json:"sign" dc:"签名"`
}

type CardPlateRelsModel struct {
	Brand         string `json:"brand" dc:"品牌,如宝马"`
	LicenseNumber string `json:"licenseNumber"  dc:"车牌"`
	LicenseType   int    `json:"licenseType" dc:"车牌类型"`
	PlateColor    string `json:"plateColor"  dc:"车辆颜色，如红"`
	PlateType     string `json:"plateType" dc:"车辆类型,如SUV"`
}

// ParkingCardRenewParam 月卡续费参数
type ParkingCardRenewParam struct {
	ParkId           string `json:"parkId" dc:"车场id,必须"`
	CardId           int64  `json:"cardId"  dc:"月卡id,必须"`
	AmountReceivable int    `json:"amountReceivable" dc:"应付金额 单位/分,必须"`
	CardCopies       int    `json:"cardCopies"  dc:"续费份数,必须"`
	StartTime        string `json:"startTime"  dc:"续费开始时间 格式:yyyy-MM-dd,必须"`
	EndTime          string `json:"endTime" dc:"续费结束时间 格式:yyyy-MM-dd,必须"`
	OldExpireTime    string `json:"oldExpireTime"  dc:"原始到期时间 格式:yyyy-MM-dd,必须"`
	Sign             string `json:"sign" dc:"签名"`
}

// ParkingCardSearchParam 月卡列表查询参数
type ParkingCardSearchParam struct {
	ParkId        string `json:"parkId" dc:"车场id,必须"`
	IsCommon      string `json:"isCommon"  dc:"是否为集团通卡 0:否 1:是,必须"`
	CardTypeName  string `json:"cardTypeName" dc:"月卡套餐名称"`
	HouseInfo     string `json:"houseInfo"  dc:"房屋信息"`
	LicenseNumber string `json:"licenseNumber" dc:"月卡绑定的车牌"`
	PhoneNumber   string `json:"phoneNumber"  dc:"手机号码"`
	Status        int    `json:"status" dc:"月卡状态 1-正常 2-冻结 3-宽限期 4-过期 6-未生效"`
	UserName      string `json:"userName"  dc:"用户名称"`
	Sign          string `json:"sign" dc:"签名"`
	common.PageParam
}

// CarBillRecordsParam 车辆计费查询参数
type CarBillRecordsParam struct {
	ParkId        string   `json:"parkId" dc:"车场id,必须"`
	LicenseNumber string   `json:"licenseNumber" dc:"车牌号,必须"`
	LicenseType   int      `json:"licenseType"  dc:"车牌类型,必须"`
	CouponIds     []string `json:"couponIds" dc:"优惠券ids"`
	IsGenPayOrder string   `json:"isGenPayOrder"  dc:"是否生成预订单 0-否，1-是"`
	ParkName      string   `json:"parkName"  dc:"停车场名称"`
	Sign          string   `json:"sign" dc:"签名"`
	common.PageParam
}

// CarPayRecordsParam 缴费记录查询参数
type CarPayRecordsParam struct {
	ParkId        string `json:"parkId" dc:"车场id,必须"`
	ParkName      string `json:"parkName"  dc:"停车场名称,必须"`
	CarType       string `json:"carType" dc:"车辆类型,必须"`
	LicenseNumber string `json:"licenseNumber" dc:"车牌号,必须"`
	LicenseType   int    `json:"licenseType"  dc:"车牌类型,必须"`
	EnterTime     string `json:"enterTime"  dc:"进场时间"`
	OutTime       string `json:"outTime" dc:"出场时间"`
	Paystarttime  string `json:"paystarttime"  dc:"支付起始时间"`
	Payendtime    string `json:"payendtime" dc:"支付结束时间"`
	Sign          string `json:"sign" dc:"签名"`
	common.PageParam
}

// ParkingSpaceReserveParam 车位预定参数
type ParkingSpaceReserveParam struct {
	ParkId         string      `json:"parkId" dc:"车场id,必须"`
	LicenseNumber  string      `json:"licenseNumber" dc:"车牌号,必须"`
	ReserveInTime  *gtime.Time `json:"reserveInTime"  dc:"预约入场时间(格式:yyyy-MM-dd HH:mm:ss),必须"`
	ReserveOutTime *gtime.Time `json:"reserveOutTime"  dc:"预约出场时间(格式:yyyy-MM-dd HH:mm:ss),必须"`
	ReserveTime    int         `json:"reserveTime" dc:"预约停车时长(分钟),必须"`
	EarliestInTime *gtime.Time `json:"earliestInTime"  dc:"最早入场时间(格式:yyyy-MM-dd HH:mm:ss)"`
	LatestInTime   *gtime.Time `json:"latestInTime"  dc:"最晚入场时间(格式:yyyy-MM-dd HH:mm:ss)"`
	CreateTime     *gtime.Time `json:"createTime"  dc:"创建时间(格式:yyyy-MM-dd HH:mm:ss)"`
	LicenseType    int         `json:"licenseType"  dc:"车牌类型"`
	UserId         string      `json:"userId"  dc:"用户id"`
	UserName       string      `json:"userName"  dc:"用户名"`
	Phone          string      `json:"phone" dc:"用户手机号"`
	Sign           string      `json:"sign" dc:"签名"`
}

// ParkingInoutParam 进出场信息查询参数
type ParkingInoutParam struct {
	ParkId        string `json:"parkId" dc:"车场id,必须"`
	ChargeType    int    `json:"chargeType" dc:"计费类型 1=长租车 2=临停车"`
	ModelName     string `json:"modelName"  dc:"计费名称"`
	ExceptionType int    `json:"exceptionType"  dc:"异常类型 1-手动入场 2-岗亭收费 3-免费开闸 4-异常离场 5-手动出场"`
	LicenseNumber string `json:"licenseNumber" dc:"车牌号"`
	Sign          string `json:"sign" dc:"签名"`
	common.PageParam
}

// GrantCouponParam 发放优惠券参数
type GrantCouponParam struct {
	ParkId        string `json:"parkId" dc:"车场id,必须"`
	Couponname    string `json:"couponname"  dc:"优惠券名称,必须"`
	Coupontype    int    `json:"coupontype" dc:"1=全免,2=当天全免,3=金额减免,4=小时减免,5=折扣减免（待上线）,必须"`
	Couponrule    int    `json:"couponrule"  dc:"优惠规则N，如果是小时券，代表优惠N小时，如果是金额券，代表优惠N元,必须"`
	LicenseNumber string `json:"licenseNumber" dc:"车牌号,必须"`
	Num           int    `json:"num" dc:"分配张数（默认1）"`
	Parkname      string `json:"parkname" dc:"车场名称"`
	Expday        int    `json:"expday" dc:"优惠券有效期（天）0-不限；其他为具体有效天数"`
	Source        int    `json:"source" dc:"来源1-普通电子券（默认）2-纸质券3-活动券4-时段券"`
	Userule       int    `json:"userule" dc:"核销判断标识 1-出场即核销（默认） 0-使用核销（如果用户没有选定此券，下次入场继续可用）"`
	Sign          string `json:"sign" dc:"签名"`
}
