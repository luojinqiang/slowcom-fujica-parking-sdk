package request

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/app/plate_device/entity"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/basic"
)

type plateDeviceRequest struct {
	basic.BaseRequest
}

var PlateDeviceRequest = new(plateDeviceRequest)

func (s *plateDeviceRequest) Get(Id int) (entity *entity.PlateDeviceEntity, err error) {
	res, err := httpclient.Get(s.BuildUrl(fmt.Sprintf("PlateDevice/Get/%d", Id)))
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
func (s *plateDeviceRequest) List(PageSize int, CurrentPage int) (page *entity.PlateDevicePageEntity, err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`PlateDevice/GetByCustom`), &basic.PageData{
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
	page = &entity.PlateDevicePageEntity{
		TotalCount: fjcPageResponse.PageAttri.TotalCount,
		List:       nil,
	}
	bytes, _ := json.Marshal(fjcPageResponse.Records)
	err = json.Unmarshal(bytes, &page.List)
	return
}
