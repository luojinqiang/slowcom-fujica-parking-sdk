package request

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"slowcom-fujica-parking-sdk/app/parking/entity"
	"slowcom-fujica-parking-sdk/basic"
)

type parkingRequest struct {
	basic.BaseRequest
}

var ParkingRequest = new(parkingRequest)

func (s *parkingRequest) Get(Id int) (entity *entity.ParkingEntity, err error) {
	res, err := httpclient.Get(s.BuildUrl(fmt.Sprintf("Park/Get/%d", Id)))
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
func (s *parkingRequest) List(PageSize int, CurrentPage int) (page *entity.ParkingPageEntity, err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`Park/GetByCustom`), &basic.PageData{
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
	page = &entity.ParkingPageEntity{
		TotalCount: fjcPageResponse.PageAttri.TotalCount,
		List:       nil,
	}
	bytes, _ := json.Marshal(fjcPageResponse.Records)
	err = json.Unmarshal(bytes, &page.List)
	return
}
