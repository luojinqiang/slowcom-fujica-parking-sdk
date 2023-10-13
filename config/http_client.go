package config

import (
	"github.com/ddliu/go-httpclient"
	url2 "net/url"
	"slowcom-fujica-parking-sdk/app/common"
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
	Token  string
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
	res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+s.Token).Get(common.BuildUrl(url), url2.Values{})
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}

// PostJson postJson请求
func (s *FsHttpClient) PostJson(url string, data interface{}) (response *FsResponse, err error) {
	res, err := buildHttpClient().WithHeader("Token", s.Token).PostJson(common.BuildUrl(url), data)
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}

// GetToken 获取token请求
func (s *FsHttpClient) GetToken() (response *FsResponse, err error) {
	sign, timestamp := common.GetSign()
	param := map[string]string{"appId": common.AppId, "appSecret": common.AppSecret, "mchId": common.MchId, "sign": sign, "timestamp": timestamp}
	res, err := buildHttpClient().PostJson(common.BuildUrl("fujica/api/v1/public/token"), param)
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}
