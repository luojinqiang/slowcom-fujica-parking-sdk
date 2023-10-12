package config

import (
	"fmt"
	"github.com/ddliu/go-httpclient"
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
	BaseUrl string
	rwLock  sync.RWMutex
	Token   string
}

//func init() {
//	httpclient.Defaults(httpclient.Map{
//		"opt_useragent":   USERAGENT,
//		"opt_timeout":     TIMEOUT,
//		"Accept-Encoding": "gzip,deflate,sdch",
//	})
//}

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
	res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+s.Token).Get(fmt.Sprintf("%s%s", s.BaseUrl, url), url2.Values{})
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}

// PostJson postJson请求
func (s *FsHttpClient) PostJson(url, data interface{}) (response *FsResponse, err error) {
	//res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+s.Token).PostJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	res, err := buildHttpClient().WithHeader("Token", s.Token).PostJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}

// TokenPost 获取token请求
func (s *FsHttpClient) TokenPost(url, data interface{}) (response *FsResponse, err error) {
	res, err := buildHttpClient().PostJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}
