package service

import (
	"github.com/ddliu/go-httpclient"
	"slowcom-fujica-parking-sdk/app/organization/entity"
	"slowcom-fujica-parking-sdk/basic"
)

type organizationService struct {
	basic.BaseService
}

var OrganizationRequest = new(organizationService)

// Add 添加组织机构
func (s *organizationService) Add(entity *entity.OrganizationEntity) (err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/Organization/Add`), entity)
	return s.IsErr(res, err)
}
