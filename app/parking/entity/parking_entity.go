package entity

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

// ParkingPageEntity 分页结果
type ParkingPageEntity struct {
	TotalCount int              `json:"totalCount"`
	List       []*ParkingEntity `json:"list"`
}

type ParkingCardParam struct {
	Sign             string   `json:"sign" dc:"签名,必须"`
	ParkId           string   `json:"parkId" dc:"车场id,必须"`
	UserName         string   `json:"userName"  dc:"用户名,必须"`
	CardTypeId       int64    `json:"cardTypeId" dc:"月卡套餐id,必须"`
	StartTime        string   `json:"startTime"  dc:"有效期起始时间 格式:yyyy-MM-dd,必须"`
	EndTime          string   `json:"endTime" dc:"有效期结束时间 格式:yyyy-MM-dd,必须"`
	NewEndTime       string   `json:"newEndTime"  dc:"实际有效期截至时间 格式:yyyy-MM-dd	,必须"`
	AmountReceivable int      `json:"amountReceivable" dc:"应收金额 单位/分,必须"`
	CardCopies       int      `json:"cardCopies"  dc:"购买份数 1-9999,必须"`
	CardPlateRels    []string `json:"cardPlateRels" dc:"车牌,必须"`
	HouseInfo        string   `json:"houseInfo"  dc:"住宅信息"`
	IsGroup          int      `json:"isGroup" dc:"是否组车 0-否 1-是"`
	NumCar           int      `json:"numCar"  dc:"绑定车辆数 最大3"`
	ParkList         []string `json:"parkList" dc:"车场id列表"`
	ParkingLot       int      `json:"parkingLot"  dc:"车位数"`
	PhoneNumber      string   `json:"phoneNumber"  dc:"联系电话"`
	Remark           string   `json:"remark"  dc:"备注"`
}
