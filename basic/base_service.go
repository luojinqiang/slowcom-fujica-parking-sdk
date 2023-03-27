package basic

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"slowcom-fujica-parking-sdk/config"
	gloabal "slowcom-fujica-parking-sdk/gloabal"
)

type BaseService struct {
}

// BuildUrl 构建URL
func (s *BaseService) BuildUrl(url string) string {
	return fmt.Sprintf("%s%s", config.BaseUrl, url)
}

// IsErr 请求是否异常
func (s *BaseService) IsErr(res *httpclient.Response, err error) error {
	if err != nil {
		return gloabal.ErrIs系统异常
	}
	if res.StatusCode != 200 {
		return gloabal.ErrIs请求状态异常
	}
	bytes, err := res.ReadAll()
	if err != nil {
		return gloabal.ErrIs数据解析异常
	}
	var resData gloabal.OperationRes
	err = json.Unmarshal(bytes, &resData)
	if err != nil {
		return gloabal.ErrIs数据解析异常
	}
	if resData.IsSucess && resData.Code == 0 {
		return nil
	} else {
		return gloabal.BusinessErr(resData.Describe)
	}
}
