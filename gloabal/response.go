package gloabal

// OperationRes 增加，删除，修改操作统一返回
type OperationRes struct {
	IsSucess       bool   `json:"IsSucess"`       //是否成功
	IsError        bool   `json:"IsError"`        // 是否失败
	Code           int    `json:"Code"`           //业务码（0=成功 1=执行成功但没有数据 2=失败
	RecordAffected int    `json:"RecordAffected"` // 受影响行数
	Describe       string `json:"Describe"`       // 描述信息
}
