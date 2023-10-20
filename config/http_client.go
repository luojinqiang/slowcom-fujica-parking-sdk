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
	Token   string
	BaseUrl string
	RwLock  sync.RWMutex
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
	res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+s.Token).Get(common.BuildUrl(s.BaseUrl, url), url2.Values{})
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}

// PostJson postJson请求
func (s *FsHttpClient) PostJson(url string, data interface{}) (response *FsResponse, err error) {
	res, err := buildHttpClient().WithHeader("Token", s.Token).PostJson(common.BuildUrl(s.BaseUrl, url), data)
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	fmt.Printf("%+v", response)
	return
}

// GetToken 获取token请求
func (s *FsHttpClient) GetToken(appId string, appSecret string, mchId string) (response *FsResponse, err error) {
	mp := map[string]interface{}{"appId": appId, "appSecret": appSecret, "mchId": mchId}
	mp["sign"] = common.GetSign(mp)
	res, err := buildHttpClient().PostJson(common.BuildUrl(s.BaseUrl, "public/token"), mp)
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}
