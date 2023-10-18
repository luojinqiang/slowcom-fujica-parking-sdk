package request

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/app/organization/entity"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/basic"
	"github.com/luojinqiang/slowcom-fujica-parking-sdk/gloabal"
)

type organizationRequest struct {
	basic.BaseRequest
}

var OrganizationRequest = new(organizationRequest)

// Add 添加组织机构
func (s *organizationRequest) Add(entity *entity.OrganizationEntity) (err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/Organization/Add`), entity)
	if err != nil {
		return gloabal.ErrIs系统异常
	}
	return s.CheckFjcResState(res)
}

// Edit 修改组织机构
func (s *organizationRequest) Edit(entity *entity.OrganizationEntity) (err error) {
	res, err := httpclient.PutJson(s.BuildUrl(`Organization/Modify`), entity)
	if err != nil {
		return gloabal.ErrIs系统异常
	}
	return s.CheckFjcResState(res)
}

// Delete 删除机构组织
func (s *organizationRequest) Delete(Rid string) (err error) {
	res, err := httpclient.PutJson(s.BuildUrl(`Organization/DeleteFunc`), "(Rid ='"+Rid+"')")
	if err != nil {
		return gloabal.ErrIs系统异常
	}
	return s.CheckFjcResState(res)
}

// Get 获取机构组织信息
func (s *organizationRequest) Get(Id int) (entity *entity.OrganizationEntity, err error) {
	res, err := httpclient.Get(s.BuildUrl(fmt.Sprintf("Organization/Get/%d", Id)))
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
func (s *organizationRequest) List() (list []*entity.OrganizationEntity, err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`Organization/GetByFunc`), "(ID > 0)")
	if err != nil {
		return
	}
	fjcResponse, err := s.CheckFjcResponse(res)
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(fjcResponse.Models)
	err = json.Unmarshal(bytes, &list)
	return
}
