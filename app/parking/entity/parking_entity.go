package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"slowcom-fujica-parking-sdk/app/common"
)

// CarBillRecordsParam 车辆计费查询参数
type CarBillRecordsParam struct {
	ParkId        string   `json:"parkId" dc:"车场id,必须"`
	LicenseNumber string   `json:"licenseNumber" dc:"车牌号,必须"`
	LicenseType   int8     `json:"licenseType"  dc:"车牌类型"`
	CouponIds     []string `json:"couponIds" dc:"优惠券ids"`
	IsGenPayOrder string   `json:"isGenPayOrder"  dc:"是否生成预订单 0-否，1-是"`
	ParkName      string   `json:"parkName"  dc:"停车场名称"`
	Sign          string   `json:"sign" dc:"签名"`
	common.PageParam
}

// CarPayAddParam 上传支付结果参数
type CarPayAddParam struct {
	ParkId      string `json:"parkId" dc:"车场id,必须"`
	OrderId     string `json:"orderId" v:"required#支付订单编号不能为空" dc:"支付订单编号,必须"`
	PayStatus   string `json:"payStatus"  dc:"支付状态:0支付中,1支付完成/成功,2支付关闭/失败,3退款成功,8建行无感支付,9第三方支付(线上),10第三方(线下),必须"`
	PayWay      int8   `json:"payWay" dc:"支付方式 第三方平台来源枚举 1：上海公共信息平台,2:苏州建发云锦湾,必须"`
	SuccessTime string `json:"successTime"  dc:"支付时间yyyy-MM-dd HH:mm:ss"`
	Sign        string `json:"sign" dc:"签名"`
}

// CarPayRecordsParam 缴费记录查询参数
type CarPayRecordsParam struct {
	ParkId        string `json:"parkId" dc:"车场id,必须"`
	CarType       int8   `json:"carType" dc:"车辆类型,必须"`
	LicenseNumber string `json:"licenseNumber" dc:"车牌号,必须"`
	LicenseType   int8   `json:"licenseType"  dc:"车牌类型,必须"`
	ParkName      string `json:"parkName"  dc:"停车场名称"`
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
	ReserveTime    int         `json:"reserveTime" dc:"预约停车时长(分钟)"`
	EarliestInTime *gtime.Time `json:"earliestInTime"  dc:"最早入场时间(格式:yyyy-MM-dd HH:mm:ss)"`
	LatestInTime   *gtime.Time `json:"latestInTime"  dc:"最晚入场时间(格式:yyyy-MM-dd HH:mm:ss)"`
	CreateTime     *gtime.Time `json:"createTime"  dc:"创建时间(格式:yyyy-MM-dd HH:mm:ss)"`
	LicenseType    int8        `json:"licenseType"  dc:"车牌类型"`
	UserId         string      `json:"userId"  dc:"用户id"`
	UserName       string      `json:"userName"  dc:"用户名"`
	Phone          string      `json:"phone" dc:"用户手机号"`
	Sign           string      `json:"sign" dc:"签名"`
}

// ParkingInoutParam 进出场信息查询参数
type ParkingInoutParam struct {
	ParkId        string `json:"parkId" dc:"车场id,必须"`
	ChargeType    int8   `json:"chargeType" dc:"计费类型 1=长租车 2=临停车"`
	ModelName     string `json:"modelName"  dc:"计费名称"`
	ExceptionType int8   `json:"exceptionType"  dc:"异常类型 1-手动入场 2-岗亭收费 3-免费开闸 4-异常离场 5-手动出场"`
	LicenseNumber string `json:"licenseNumber" dc:"车牌号"`
	Sign          string `json:"sign" dc:"签名"`
	Current       int    `json:"current" dc:"分页-当前页码（从1开始），必须(出场接口不必须)"`
	Size          int    `json:"size" dc:"分页-每页记录数（默认为10），必须(出场接口不必须)"`
}

// GrantCouponParam 发放优惠券参数
type GrantCouponParam struct {
	ParkId        string `json:"parkId" dc:"车场id,必须"`
	Couponname    string `json:"couponname"  dc:"优惠券名称,必须"`
	Coupontype    int8   `json:"coupontype" dc:"1=全免,2=当天全免,3=金额减免,4=小时减免,5=折扣减免（待上线）,必须"`
	Couponrule    int    `json:"couponrule"  dc:"优惠规则N，如果是小时券，代表优惠N小时，如果是金额券，代表优惠N元,必须"`
	LicenseNumber string `json:"licenseNumber" dc:"车牌号,必须"`
	Num           int    `json:"num" dc:"分配张数（默认1）"`
	Parkname      string `json:"parkname" dc:"车场名称"`
	Expday        int    `json:"expday" dc:"优惠券有效期（天）0-不限；其他为具体有效天数"`
	Source        int8   `json:"source" dc:"来源1-普通电子券（默认）2-纸质券3-活动券4-时段券"`
	Userule       int8   `json:"userule" dc:"核销判断标识 1-出场即核销（默认） 0-使用核销（如果用户没有选定此券，下次入场继续可用）"`
	Sign          string `json:"sign" dc:"签名"`
}
