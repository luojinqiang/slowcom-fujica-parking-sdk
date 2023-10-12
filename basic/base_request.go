package basic

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"slowcom-fujica-parking-sdk/app/common"
	"slowcom-fujica-parking-sdk/gloabal"
)

type PageData struct {
	PageSize    int    `json:"PageSize"`
	CurrentPage int    `json:"CurrentPage"`
	OrderBy     string `json:"OrderBy"`    //排序字段
	OrderType   bool   `json:"OrderType"`  //排序类型
	Where       string `json:"where"`      //查询条件
	Append      string `json:"Append"`     //附加条件
	TotalCount  int    `json:"TotalCount"` //总结果数
}

type BaseRequest struct {
}

// BuildUrl 构建URL
func (s *BaseRequest) BuildUrl(url string) string {
	return fmt.Sprintf("%s%s", common.BaseUrl, url)
}

// CheckFjcResState 请求是否异常
func (s *BaseRequest) CheckFjcResState(res *httpclient.Response) error {
	if res.StatusCode != 200 {
		return gloabal.ErrIs请求状态异常
	}
	bytes, err := res.ReadAll()
	if err != nil {
		return gloabal.ErrIs数据解析异常
	}
	var fjcResState gloabal.FjcResState
	err = json.Unmarshal(bytes, &fjcResState)
	if err != nil {
		return gloabal.ErrIs数据解析异常
	}
	if fjcResState.IsSucess && fjcResState.Code == 0 {
		return nil
	} else {
		return gloabal.BusinessErr(fjcResState.Describe)
	}
}

// CheckFjcResponse 请求是否异常
func (s *BaseRequest) CheckFjcResponse(res *httpclient.Response) (fjcResponse *gloabal.FjcResponse, err error) {
	if res.StatusCode != 200 {
		return nil, gloabal.ErrIs请求状态异常
	}
	bytes, err := res.ReadAll()
	if err != nil {
		return nil, gloabal.ErrIs数据解析异常
	}
	err = json.Unmarshal(bytes, &fjcResponse)
	if err != nil {
		return nil, gloabal.ErrIs数据解析异常
	}
	if fjcResponse.State.IsSucess && fjcResponse.State.Code == 0 {
		return fjcResponse, nil
	} else {
		return nil, gloabal.BusinessErr(fjcResponse.State.Describe)
	}
}

// CheckFjcPageResponse 分页请求是否异常
func (s *BaseRequest) CheckFjcPageResponse(res *httpclient.Response) (fjcPageResponse *gloabal.FjcPageResponse, err error) {
	if res.StatusCode != 200 {
		return nil, gloabal.ErrIs请求状态异常
	}
	bytes, err := res.ReadAll()
	if err != nil {
		return nil, gloabal.ErrIs数据解析异常
	}
	err = json.Unmarshal(bytes, &fjcPageResponse)
	if err != nil {
		return nil, gloabal.ErrIs数据解析异常
	}
	if fjcPageResponse.State.IsSucess && fjcPageResponse.State.Code == 0 {
		return fjcPageResponse, nil
	} else {
		return nil, gloabal.BusinessErr(fjcPageResponse.State.Describe)
	}
}
