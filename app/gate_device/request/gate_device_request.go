package request

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"slowcom-fujica-parking-sdk/app/gate_device/entity"
	"slowcom-fujica-parking-sdk/basic"
)

type gateDeviceRequest struct {
	basic.BaseRequest
}

var GateDeviceRequest = new(gateDeviceRequest)

func (s *gateDeviceRequest) Get(Id int) (entity *entity.GateDeviceEntity, err error) {
	res, err := httpclient.Get(s.BuildUrl(fmt.Sprintf("GateControllerDevice/Get/%d", Id)))
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
func (s *gateDeviceRequest) List(PageSize int, CurrentPage int) (page *entity.GateDevicePageEntity, err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`GateControllerDevice/GetByCustom`), &basic.PageData{
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
	page = &entity.GateDevicePageEntity{
		TotalCount: fjcPageResponse.PageAttri.TotalCount,
		List:       nil,
	}
	bytes, _ := json.Marshal(fjcPageResponse.Records)
	err = json.Unmarshal(bytes, &page.List)
	return
}
