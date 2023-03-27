package entity

// OrganizationEntity 组织机构
type OrganizationEntity struct {
	ID       int    `json:"ID"`       //新增可不传或者传0
	Name     string `json:"Name"`     //机构名称
	ParentId int    `json:"ParentId"` //上级组织，无则不填
	Remark   string `json:"Remark"`   //备注
	Rid      string `json:"Rid"`      //标识号，自定义
	Gid      string `json:"Gid"`      // 集团标识号，由富士分配填写
}
