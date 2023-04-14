package request

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"slowcom-fujica-parking-sdk/app/reader_device/entity"
	"slowcom-fujica-parking-sdk/basic"
)

type readerDeviceRequest struct {
	basic.BaseRequest
}

var ReaderDeviceRequest = new(readerDeviceRequest)

func (s *readerDeviceRequest) Get(Id int) (entity *entity.ReaderDeviceEntity, err error) {
	res, err := httpclient.Get(s.BuildUrl(fmt.Sprintf("ReaderDevice/Get/%d", Id)))
	if err != nil {
		return
	}
	fjcResponse, err := s.CheckFjcResponse(res)
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(fjcResponse.Model)
	err = json.Unmarshal(bytes, &entity)
	return
}

// List 查询列表
func (s *readerDeviceRequest) List(PageSize int, CurrentPage int) (page *entity.ReaderDevicePageEntity, err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`ReaderDevice/GetByCustom`), &basic.PageData{
		PageSize:    PageSize,
		CurrentPage: CurrentPage,
	})
	if err != nil {
		return
	}
	fjcPageResponse, err := s.CheckFjcPageResponse(res)
	if err != nil {
		return
	}
	page = &entity.ReaderDevicePageEntity{
		TotalCount: fjcPageResponse.PageAttri.TotalCount,
		List:       nil,
	}
	bytes, _ := json.Marshal(fjcPageResponse.Records)
	err = json.Unmarshal(bytes, &page.List)
	return
}
