package config

import (
	"encoding/json"
	"github.com/ddliu/go-httpclient"
	"strconv"
)

// FsResponse 返回
type FsResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// checkResponse 校验请求
func checkResponse(res *httpclient.Response) (FsResponse *FsResponse, err error) {
	if res.StatusCode != 200 {
		return nil, businessErr(500, "请求状态异常")
	}
	bytes, err := res.ReadAll()
	if err != nil {
		return nil, businessErr(500, "数据解析异常")
	}
	err = json.Unmarshal(bytes, &FsResponse)
	if err != nil {
		return nil, businessErr(500, "数据解析异常")
	}
	if FsResponse.Code == 1000000 || FsResponse.Code == 0 {
		return
	} else {
		for _, data := range ErrArr {
			if FsResponse.Code == data.Code {
				FsResponse.Message = data.Message
				err = businessErr(data.Code, data.Message)
				return nil, err
			}
		}
		if err == nil {
			return nil, businessErr(FsResponse.Code, FsResponse.Message)
		}
	}
	return
}

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return strconv.Itoa(e.Code) + "-" + e.Message
}

var ErrArr = []*Error{
	{
		Code:    1000001,
		Message: "操作失败",
	},
	{
		Code:    100000001,
		Message: "超时未响应",
	},
	{
		Code:    100000014,
		Message: "未录入授权TOKEN",
	},
	{
		Code:    100000017,
		Message: "签名不正确",
	},
}

// businessErr 业务异常
func businessErr(code int, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
