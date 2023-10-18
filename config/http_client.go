package config

import (
	"fmt"
	"github.com/ddliu/go-httpclient"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/app/common"
	url2 "net/url"
	"sync"
)

const (
	USERAGENT      = "slowcom_agent"
	TIMEOUT        = 30
	ConnectTimeout = 10
)

// FsHttpClient 富士httpClient
type FsHttpClient struct {
	rwLock sync.RWMutex
}

// 生成一个http请求客户端
func buildHttpClient() *httpclient.HttpClient {
	return httpclient.NewHttpClient().Defaults(httpclient.Map{
		"opt_useragent":      USERAGENT,
		"opt_timeout":        TIMEOUT,
		"Accept-Encoding":    "gzip,deflate,sdch",
		"opt_connecttimeout": ConnectTimeout,
		"OPT_DEBUG":          true,
	})
}

// Get get请求
func (s *FsHttpClient) Get(url string) (response *FsResponse, err error) {
	res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+common.Token).Get(common.BuildUrl(url), url2.Values{})
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}

// PostJson postJson请求
func (s *FsHttpClient) PostJson(url string, data interface{}) (response *FsResponse, err error) {
	res, err := buildHttpClient().WithHeader("Token", common.Token).PostJson(common.BuildUrl(url), data)
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	fmt.Printf("%+v", response)
	return
}

// GetToken 获取token请求
func (s *FsHttpClient) GetToken() (response *FsResponse, err error) {
	mp := map[string]interface{}{"appId": common.AppId, "appSecret": common.AppSecret, "mchId": common.MchId}
	mp["sign"] = common.GetSign(mp)
	res, err := buildHttpClient().PostJson(common.BuildUrl("fujica/api/v1/public/token"), mp)
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}
