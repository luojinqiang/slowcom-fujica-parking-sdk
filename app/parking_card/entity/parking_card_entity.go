package entity

import "github.com/luojinqiang/slowcom-fujica-parking-sdk/app/common"

// ParkingCardRuleModel 月卡套餐model
type ParkingCardRuleModel struct {
	CardTypeId       int64  `json:"cardTypeId" dc:"月卡套餐id"`
	CardCategory     int8   `json:"cardCategory" dc:"套餐类别 1-月卡 2-季卡 3-半年卡 4-年卡"`
	CardTypeName     string `json:"cardTypeName" dc:"月卡套餐名称"`
	ChronographType  int8   `json:"chronographType"  dc:"计时类型"`
	Fixed            int8   `json:"fixed" dc:"固定车位(0-否/1-是)"`
	FreeType         int8   `json:"freeType"  dc:"优免类型 1-全时段 2-分时段"`
	GracePeriod      int    `json:"gracePeriod" dc:"续费宽限时长 单位/天"`
	HolidayEndTime   string `json:"holidayEndTime"  dc:"节假日免费时间 结束时间--23:59"`
	HolidayStartTime string `json:"holidayStartTime" dc:"节假日免费时间 开始时间--00:00"`
	LicenseType      int8   `json:"licenseType"  dc:"车牌类型(1蓝 2绿 3黄 4白 5黑 6其他 7黄绿)"`
	OverdueCharge    int8   `json:"overdueCharge"  dc:"过期计费类型 1-按临时车计费 2-延期有效期内按月卡计费"`
	ParkId           string `json:"parkId"  dc:"停车场ID"`
	Price            int    `json:"price" dc:"套餐金额 单位/分"`
	Remark           string `json:"remark"  dc:"备注"`
	WeekendEndTime   string `json:"weekendEndTime"  dc:"周末免费时间 结束时间"`
	WeekendStartTime string `json:"weekendStartTime"  dc:"周末免费时间 开始时间"`
	WorkdayEndTime   string `json:"workdayEndTime"  dc:"工作日免费时间 结束时间"`
	WorkdayStartTime string `json:"workdayStartTime"  dc:"工作日免费时间 开始时间"`
}

type CardPlateRelsModel struct {
	Brand         string `json:"brand" dc:"品牌,如宝马"`
	LicenseNumber string `json:"licenseNumber"  dc:"车牌"`
	LicenseType   int8   `json:"licenseType" dc:"车牌类型"`
	PlateColor    string `json:"plateColor"  dc:"车辆颜色，如红"`
	PlateType     string `json:"plateType" dc:"车辆类型,如SUV"`
}

// ParkingCardModel 月卡model
type ParkingCardModel struct {
	CardId        int64                 `json:"cardId" dc:"月卡id"`
	CardTypeId    int64                 `json:"cardTypeId" dc:"月卡套餐id"`
	ParkId        string                `json:"parkId" dc:"车场id"`
	CardCopies    int                   `json:"cardCopies"  dc:"购买份数 1-9999"`
	CardPlateRels []*CardPlateRelsModel `json:"cardPlateRels" dc:"车牌信息"`
	StartTime     string                `json:"startTime"  dc:"有效期起始时间 格式:yyyy-MM-dd"`
	EndTime       string                `json:"endTime" dc:"有效期结束时间 格式:yyyy-MM-dd"`
	NewEndTime    string                `json:"newEndTime"  dc:"实际有效期截至时间 格式:yyyy-MM-dd"`
	ParkList      []string              `json:"parkList" dc:"车场id列表"`
	UserName      string                `json:"userName"  dc:"用户名"`
	PhoneNumber   string                `json:"phoneNumber"  dc:"联系电话"`
	Remark        string                `json:"remark"  dc:"备注"`
	Status        int8                  `json:"status" dc:"月卡状态(1-正常 2-冻结 3-宽限期 4-过期 5-注销 6-未生效)"`
}

// ParkingCardResult 月卡列表结果
type ParkingCardResult struct {
	Records []*ParkingCardModel `json:"records" dc:"月卡列表"`
	common.PageResult
}

// ParkingCardApplyParam 月卡开通参数
type ParkingCardApplyParam struct {
	ParkId           string                `json:"parkId" dc:"车场id,必须"`
	UserName         string                `json:"userName"  dc:"用户名,必须"`
	PhoneNumber      string                `json:"phoneNumber"  dc:"联系电话,必须"`
	CardTypeId       int64                 `json:"cardTypeId" dc:"月卡套餐id,必须"`
	StartTime        string                `json:"startTime"  dc:"有效期起始时间 格式:yyyy-MM-dd,必须"`
	EndTime          string                `json:"endTime" dc:"有效期结束时间 格式:yyyy-MM-dd,必须"`
	NewEndTime       string                `json:"newEndTime"  dc:"实际有效期截至时间 格式:yyyy-MM-dd,必须"`
	AmountReceivable int                   `json:"amountReceivable" dc:"应收金额 单位/分,必须"`
	CardCopies       int                   `json:"cardCopies"  dc:"购买份数 1-9999,必须"`
	CardPlateRels    []*CardPlateRelsModel `json:"cardPlateRels" dc:"车牌,必须"`
	HouseInfo        string                `json:"houseInfo"  dc:"住宅信息"`
	IsGroup          int8                  `json:"isGroup" dc:"是否组车 0-否 1-是"`
	NumCar           int                   `json:"numCar"  dc:"绑定车辆数 最大3"`
	ParkList         []string              `json:"parkList" dc:"车场id列表"`
	ParkingLot       int                   `json:"parkingLot"  dc:"车位数"`
	Remark           string                `json:"remark"  dc:"备注"`
	Sign             string                `json:"sign" dc:"签名"`
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
	Status        int8   `json:"status" dc:"月卡状态 1-正常 2-冻结 3-宽限期 4-过期 6-未生效，必须"`
	UserName      string `json:"userName"  dc:"用户名称"`
	Sign          string `json:"sign" dc:"签名"`
	Current       int    `json:"current" dc:"分页-当前页码（从1开始），必须"`
	Size          int    `json:"size" dc:"分页-每页记录数（默认为10），必须"`
}

// CarOwnerInfoModel 车主及月卡信息model
type CarOwnerInfoModel struct {
	CardId      string                `json:"cardId" dc:"月卡id"`
	StartTime   string                `json:"startTime"  dc:"月卡开始日期 格式:yyyy-MM-dd"`
	EndTime     string                `json:"endTime" dc:"月卡到期日期 格式:yyyy-MM-dd"`
	PlateInfo   []*CardPlateRelsModel `json:"plateInfo" dc:"车牌信息"`
	PhoneNumber string                `json:"phoneNumber"  dc:"联系电话"`
	UserName    string                `json:"userName"  dc:"用户名"`
	Status      int8                  `json:"status" dc:"月卡状态(1-正常 2-冻结 3-宽限期 4-过期 5-注销 6-未生效)"`
	HouseInfo   string                `json:"houseInfo"  dc:"住宅信息"`
	ParkName    string                `json:"parkName"  dc:"车场名称"`
}
