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
