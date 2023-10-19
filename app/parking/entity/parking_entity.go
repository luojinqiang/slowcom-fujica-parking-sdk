package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/app/common"
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
	InOut         int8   `json:"inOut" dc:"进出场 1=进场 2=出场，必须"`
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
	Parkid        string `json:"parkid" dc:"车场id,必须"`
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

// ParkingInoutResult 车辆进出场记录结果
type ParkingInoutResult struct {
	Records []*ParkingInoutModel `json:"records" dc:"进出场记录列表"`
	common.PageResult
}

// ParkingInoutModel 车辆进出场记录model
type ParkingInoutModel struct {
	ChargeType        int8   `json:"chargeType" dc:"计费类型 1-长租车 2-临停车"`
	CouponAmount      int    `json:"couponAmount" dc:"优惠金额"`
	EnterTime         string `json:"enterTime" dc:"进场时间"`
	EntranceId        string `json:"entranceId" dc:"停车场入口id"`
	EntranceName      string `json:"entranceName" dc:"停车场入口名"`
	Havechildren      string `json:"havechildren" dc:"是否有子车场的入场记录 0-没有 1-有"`
	Hour              int    `json:"hour" dc:"已停时长 小时数"`
	Id                string `json:"id" dc:"进出场id"`
	InUrl             string `json:"inUrl" dc:"车牌入场url–大图"`
	InUrlSmall        string `json:"inUrlSmall" dc:"车牌入场url–小图"`
	IsHavechildren    bool   `json:"isHavechildren" dc:"是否有子车场的入场记录 false-没有 true-有"`
	LicenseNumber1    string `json:"licenseNumber1" dc:"车牌号1"`
	LicenseNumber2    string `json:"licenseNumber2" dc:"车牌号2"`
	LicenseType       int8   `json:"licenseType" dc:"车牌类型 1-蓝牌 2-绿牌 3-黄牌 4-白牌 5-黑牌 6-其他"`
	Minute            int    `json:"minute" dc:"已停时长 分钟数"`
	OrderStatus       int8   `json:"orderStatus" dc:"订单状态 0-正常 1-异常"`
	Parentparkinid    string `json:"parentparkinid" dc:"父车场的入场记录id"`
	ParkId            string `json:"parkId" dc:"停车场id"`
	ParkName          string `json:"parkName" dc:"停车场名称"`
	Paytype           int8   `json:"paytype" dc:"支付方式 1：微信 2：支付宝 3：银联 4：月卡 5：现金"`
	Phone             string `json:"phone" dc:"手机号"`
	RealTotalAmount   int    `json:"realTotalAmount" dc:"实收总金额"`
	ShouldTotalAmount int    `json:"shouldTotalAmount" dc:"应收总金额"`
	Staytime          int    `json:"staytime" dc:"已停时长 总分钟数"`
	// 出场信息字段
	EnterId       string `json:"enterId" dc:"入场通道"`
	ExitTime      string `json:"exitTime" dc:"出场时间"`
	ImgUrl        string `json:"imgUrl" dc:"出场照片url"`
	ImgUrl2       string `json:"imgUrl2" dc:"出场照片url"`
	LicenseNumber string `json:"licenseNumber" dc:"车牌号1"`
	ParkInId      string `json:"parkInId" dc:"入场id"`
	ParkInImg     string `json:"parkInImg" dc:"入场照片"`
	PayType       int    `json:"payType" dc:"支付方式 1：微信 2：支付宝 3：银联 4：月卡 5：现金"`
}

// ParkingLaneModel 车道信息model
type ParkingLaneModel struct {
	Boxid        string `json:"boxid" dc:"岗亭id"`
	Commonname   string `json:"commonname" dc:"共道车道名称"`
	Createtime   string `json:"createtime" dc:"创建时间"`
	Deviceserial string `json:"deviceserial" dc:"萤石云 设备序列号,存在英文字母的设备序列号，字母需为大写"`
	Hangup       int8   `json:"hangup" dc:"是否挂起0:挂起 1:未挂起"`
	Id           int    `json:"id" dc:"id"`
	Iscommon     int    `json:"iscommon" dc:"是否共道"`
	Lanecode     string `json:"lanecode" dc:"交警上传编码"`
	Lanecodeid   string `json:"lanecodeid" dc:"交警进出口编码"`
	Laneid       string `json:"laneid" dc:"车道id"`
	Lanename     string `json:"lanename" dc:"车道名称"`
	Laneno       int    `json:"laneno" dc:"通道编码(中转使用)"`
	Lanerelation string `json:"lanerelation" dc:"车道关联关系"`
	Lanetype     string `json:"lanetype" dc:"车道类型"`
	OpType       int8   `json:"opType" dc:"车道类型"`
	Parkid       string `json:"parkid" dc:"停车场id"`
	Remark       string `json:"remark" dc:"挂起原因"`
	Status       string `json:"status" dc:"萤石云设备状态1-开启 2-关闭"`
	Updatetime   string `json:"updatetime" dc:"修改时间"`
	Validatecode string `json:"validatecode" dc:"萤石云 设备验证码，设备机身上的六位大写字母"`
}
