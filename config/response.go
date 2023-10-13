package config

import (
	"encoding/json"
	"github.com/ddliu/go-httpclient"
)

// FsResponse 返回
type FsResponse struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// checkResponse 解析请求结果
func checkResponse(res *httpclient.Response) (fsResponse *FsResponse, err error) {
	if res.StatusCode != 200 {
		return nil, buildErr(res.StatusCode, res.Status)
	}
	bytes, err := res.ReadAll()
	if err != nil {
		return nil, buildErr(500, "数据解析异常")
	}
	err = json.Unmarshal(bytes, &fsResponse)
	if err != nil {
		return nil, buildErr(500, "数据解析异常")
	}
	if fsResponse != nil && fsResponse.Code != 1000000 {
		return fsResponse, buildErr(fsResponse.Code, fsResponse.Msg)
	}
	return
}
