package gloabal

import "github.com/luojinqiang/slowcom-fujica-parking-sdk/basic"

// FjcResState 返回结果状态
type FjcResState struct {
	IsSucess       bool   `json:"IsSucess"`       //是否成功
	IsError        bool   `json:"IsError"`        // 是否失败
	Code           int    `json:"Code"`           //业务码（0=成功 1=执行成功但没有数据 2=失败
	RecordAffected int    `json:"RecordAffected"` // 受影响行数
	Describe       string `json:"Describe"`       // 描述信息
}

// FjcResponse 返回结果
type FjcResponse struct {
	State  *FjcResState             `json:"State"`
	Model  map[string]interface{}   `json:"Model"`
	Models []map[string]interface{} `json:"Models"`
}

// FjcPageResponse 分页返回结果
type FjcPageResponse struct {
	PageAttri *basic.PageData          `json:"PageAttri"`
	State     *FjcResState             `json:"State"`
	Records   []map[string]interface{} `json:"Models"`
}
