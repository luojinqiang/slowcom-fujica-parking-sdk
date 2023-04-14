package entity

// GateDeviceEntity 开闸设备
type GateDeviceEntity struct {
	ID                int    `json:"ID"`
	Name              string `json:"Name"`
	CommType          string `json:"CommType"`
	EndPoint          string `json:"EndPoint"`
	DeviceCategory    string `json:"DeviceCategory"`
	LaneId            string `json:"LaneId"`
	Remark            string `json:"Remark"`
	Gid               string `json:"Gid"`
	HasExtendedfields string `json:"hasExtendedfields"` //是否扩展属性
	Rid               string `json:"Rid"`
}

// GateDevicePageEntity 分页结果
type GateDevicePageEntity struct {
	TotalCount int                 `json:"totalCount"`
	List       []*GateDeviceEntity `json:"list"`
}
