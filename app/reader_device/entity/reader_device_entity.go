package entity

// ReaderDeviceEntity 读卡器设备
type ReaderDeviceEntity struct {
	ID                int    `json:"ID"`
	Name              string `json:"Name"`
	DevNo             int    `json:"DevNo"`
	CommType          string `json:"CommType"`
	EndPoint          string `json:"EndPoint"`
	Parames           string `json:"Parames"`
	Model             string `json:"Model"`
	UserId            string `json:"UserId"`
	Password          string `json:"Password"`
	DeviceCategory    string `json:"DeviceCategory"`
	LaneId            string `json:"LaneId"`
	Remark            string `json:"Remark"`
	Gid               string `json:"Gid"`
	HasExtendedfields string `json:"hasExtendedfields"`
	Rid               string `json:"Rid"`
}

// ReaderDevicePageEntity 分页结果
type ReaderDevicePageEntity struct {
	TotalCount int                   `json:"totalCount"`
	List       []*ReaderDeviceEntity `json:"list"`
}
