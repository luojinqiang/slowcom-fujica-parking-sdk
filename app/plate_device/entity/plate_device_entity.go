package entity

// PlateDeviceEntity 识别器设备
type PlateDeviceEntity struct {
	ID                int    `json:"ID"`
	Name              string `json:"Name"`
	Ip                string `json:"Ip"`
	Port              int    `json:"Port"`
	UserId            string `json:"UserId"`
	Password          string `json:"Password"`
	DeviceCategory    string `json:"DeviceCategory"`
	LaneId            string `json:"LaneId"`
	Remark            string `json:"Remark"`
	Gid               string `json:"Gid"`
	HasExtendedfields int    `json:"hasExtendedfields"`
	Rid               string `json:"Rid"`
	ConfigVersion     string `json:"ConfigVersion"`
	MacAddress        string `json:"MacAddress"`
	Reserv1           int    `json:"Reserv1"`
	Reserv2           string `json:"Reserv2"`
	Reserv3           string `json:"Reserv3"`
}

// PlateDevicePageEntity 分页结果
type PlateDevicePageEntity struct {
	TotalCount int                  `json:"totalCount"`
	List       []*PlateDeviceEntity `json:"list"`
}
