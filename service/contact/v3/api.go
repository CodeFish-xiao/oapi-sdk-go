// Package contact code generated by oapi sdk gen
package larkcontact

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *ContactService {
	c := &ContactService{config: config}
	c.CustomAttr = &customAttr{service: c}
	c.CustomAttrEvent = &customAttrEvent{service: c}
	c.Department = &department{service: c}
	c.EmployeeTypeEnum = &employeeTypeEnum{service: c}
	c.Group = &group{service: c}
	c.GroupMember = &groupMember{service: c}
	c.Scope = &scope{service: c}
	c.Unit = &unit{service: c}
	c.User = &user{service: c}
	return c
}

// 业务域服务定义
type ContactService struct {
	config           *larkcore.Config
	CustomAttr       *customAttr
	CustomAttrEvent  *customAttrEvent
	Department       *department
	EmployeeTypeEnum *employeeTypeEnum
	Group            *group
	GroupMember      *groupMember
	Scope            *scope
	Unit             *unit
	User             *user
}

// 资源服务定义
type customAttr struct {
	service *ContactService
}
type customAttrEvent struct {
	service *ContactService
}
type department struct {
	service *ContactService
}
type employeeTypeEnum struct {
	service *ContactService
}
type group struct {
	service *ContactService
}
type groupMember struct {
	service *ContactService
}
type scope struct {
	service *ContactService
}
type unit struct {
	service *ContactService
}
type user struct {
	service *ContactService
}

// 资源服务方法定义
func (c *customAttr) List(ctx context.Context, req *ListCustomAttrReq, options ...larkcore.RequestOptionFunc) (*ListCustomAttrResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, c.service.config, http.MethodGet,
		"/open-apis/contact/v3/custom_attrs", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListCustomAttrResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *customAttr) ListByIterator(ctx context.Context, req *ListCustomAttrReq, options ...larkcore.RequestOptionFunc) (*ListCustomAttrIterator, error) {
	return &ListCustomAttrIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (d *department) Children(ctx context.Context, req *ChildrenDepartmentReq, options ...larkcore.RequestOptionFunc) (*ChildrenDepartmentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/contact/v3/departments/:department_id/children", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ChildrenDepartmentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *department) ChildrenByIterator(ctx context.Context, req *ChildrenDepartmentReq, options ...larkcore.RequestOptionFunc) (*ChildrenDepartmentIterator, error) {
	return &ChildrenDepartmentIterator{
		ctx:      ctx,
		req:      req,
		listFunc: d.Children,
		options:  options,
		limit:    req.Limit}, nil
}
func (d *department) Create(ctx context.Context, req *CreateDepartmentReq, options ...larkcore.RequestOptionFunc) (*CreateDepartmentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodPost,
		"/open-apis/contact/v3/departments", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDepartmentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *department) Delete(ctx context.Context, req *DeleteDepartmentReq, options ...larkcore.RequestOptionFunc) (*DeleteDepartmentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodDelete,
		"/open-apis/contact/v3/departments/:department_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteDepartmentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *department) Get(ctx context.Context, req *GetDepartmentReq, options ...larkcore.RequestOptionFunc) (*GetDepartmentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/contact/v3/departments/:department_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDepartmentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *department) List(ctx context.Context, req *ListDepartmentReq, options ...larkcore.RequestOptionFunc) (*ListDepartmentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/contact/v3/departments", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListDepartmentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *department) ListByIterator(ctx context.Context, req *ListDepartmentReq, options ...larkcore.RequestOptionFunc) (*ListDepartmentIterator, error) {
	return &ListDepartmentIterator{
		ctx:      ctx,
		req:      req,
		listFunc: d.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (d *department) Parent(ctx context.Context, req *ParentDepartmentReq, options ...larkcore.RequestOptionFunc) (*ParentDepartmentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/contact/v3/departments/parent", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ParentDepartmentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *department) ParentByIterator(ctx context.Context, req *ParentDepartmentReq, options ...larkcore.RequestOptionFunc) (*ParentDepartmentIterator, error) {
	return &ParentDepartmentIterator{
		ctx:      ctx,
		req:      req,
		listFunc: d.Parent,
		options:  options,
		limit:    req.Limit}, nil
}
func (d *department) Patch(ctx context.Context, req *PatchDepartmentReq, options ...larkcore.RequestOptionFunc) (*PatchDepartmentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodPatch,
		"/open-apis/contact/v3/departments/:department_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchDepartmentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *department) Search(ctx context.Context, req *SearchDepartmentReq, options ...larkcore.RequestOptionFunc) (*SearchDepartmentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodPost,
		"/open-apis/contact/v3/departments/search", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchDepartmentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *department) SearchByIterator(ctx context.Context, req *SearchDepartmentReq, options ...larkcore.RequestOptionFunc) (*SearchDepartmentIterator, error) {
	return &SearchDepartmentIterator{
		ctx:      ctx,
		req:      req,
		listFunc: d.Search,
		options:  options,
		limit:    req.Limit}, nil
}
func (d *department) UnbindDepartmentChat(ctx context.Context, req *UnbindDepartmentChatDepartmentReq, options ...larkcore.RequestOptionFunc) (*UnbindDepartmentChatDepartmentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodPost,
		"/open-apis/contact/v3/departments/unbind_department_chat", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UnbindDepartmentChatDepartmentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *department) Update(ctx context.Context, req *UpdateDepartmentReq, options ...larkcore.RequestOptionFunc) (*UpdateDepartmentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodPut,
		"/open-apis/contact/v3/departments/:department_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateDepartmentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *employeeTypeEnum) Create(ctx context.Context, req *CreateEmployeeTypeEnumReq, options ...larkcore.RequestOptionFunc) (*CreateEmployeeTypeEnumResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodPost,
		"/open-apis/contact/v3/employee_type_enums", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateEmployeeTypeEnumResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *employeeTypeEnum) Delete(ctx context.Context, req *DeleteEmployeeTypeEnumReq, options ...larkcore.RequestOptionFunc) (*DeleteEmployeeTypeEnumResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodDelete,
		"/open-apis/contact/v3/employee_type_enums/:enum_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteEmployeeTypeEnumResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *employeeTypeEnum) List(ctx context.Context, req *ListEmployeeTypeEnumReq, options ...larkcore.RequestOptionFunc) (*ListEmployeeTypeEnumResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodGet,
		"/open-apis/contact/v3/employee_type_enums", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListEmployeeTypeEnumResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *employeeTypeEnum) ListByIterator(ctx context.Context, req *ListEmployeeTypeEnumReq, options ...larkcore.RequestOptionFunc) (*ListEmployeeTypeEnumIterator, error) {
	return &ListEmployeeTypeEnumIterator{
		ctx:      ctx,
		req:      req,
		listFunc: e.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (e *employeeTypeEnum) Update(ctx context.Context, req *UpdateEmployeeTypeEnumReq, options ...larkcore.RequestOptionFunc) (*UpdateEmployeeTypeEnumResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodPut,
		"/open-apis/contact/v3/employee_type_enums/:enum_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateEmployeeTypeEnumResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) Create(ctx context.Context, req *CreateGroupReq, options ...larkcore.RequestOptionFunc) (*CreateGroupResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, g.service.config, http.MethodPost,
		"/open-apis/contact/v3/group", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateGroupResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) Delete(ctx context.Context, req *DeleteGroupReq, options ...larkcore.RequestOptionFunc) (*DeleteGroupResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, g.service.config, http.MethodDelete,
		"/open-apis/contact/v3/group/:group_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteGroupResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) Get(ctx context.Context, req *GetGroupReq, options ...larkcore.RequestOptionFunc) (*GetGroupResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, g.service.config, http.MethodGet,
		"/open-apis/contact/v3/group/:group_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetGroupResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) MemberBelong(ctx context.Context, req *MemberBelongGroupReq, options ...larkcore.RequestOptionFunc) (*MemberBelongGroupResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, g.service.config, http.MethodGet,
		"/open-apis/contact/v3/group/member_belong", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &MemberBelongGroupResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) Patch(ctx context.Context, req *PatchGroupReq, options ...larkcore.RequestOptionFunc) (*PatchGroupResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, g.service.config, http.MethodPatch,
		"/open-apis/contact/v3/group/:group_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchGroupResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) Simplelist(ctx context.Context, req *SimplelistGroupReq, options ...larkcore.RequestOptionFunc) (*SimplelistGroupResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, g.service.config, http.MethodGet,
		"/open-apis/contact/v3/group/simplelist", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SimplelistGroupResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) SimplelistByIterator(ctx context.Context, req *SimplelistGroupReq, options ...larkcore.RequestOptionFunc) (*SimplelistGroupIterator, error) {
	return &SimplelistGroupIterator{
		ctx:      ctx,
		req:      req,
		listFunc: g.Simplelist,
		options:  options,
		limit:    req.Limit}, nil
}
func (g *groupMember) Add(ctx context.Context, req *AddGroupMemberReq, options ...larkcore.RequestOptionFunc) (*AddGroupMemberResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, g.service.config, http.MethodPost,
		"/open-apis/contact/v3/group/:group_id/member/add", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &AddGroupMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *groupMember) BatchAdd(ctx context.Context, req *BatchAddGroupMemberReq, options ...larkcore.RequestOptionFunc) (*BatchAddGroupMemberResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, g.service.config, http.MethodPost,
		"/open-apis/contact/v3/group/:group_id/member/batch_add", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchAddGroupMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *groupMember) BatchRemove(ctx context.Context, req *BatchRemoveGroupMemberReq, options ...larkcore.RequestOptionFunc) (*BatchRemoveGroupMemberResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, g.service.config, http.MethodPost,
		"/open-apis/contact/v3/group/:group_id/member/batch_remove", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchRemoveGroupMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *groupMember) Remove(ctx context.Context, req *RemoveGroupMemberReq, options ...larkcore.RequestOptionFunc) (*RemoveGroupMemberResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, g.service.config, http.MethodPost,
		"/open-apis/contact/v3/group/:group_id/member/remove", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RemoveGroupMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *groupMember) Simplelist(ctx context.Context, req *SimplelistGroupMemberReq, options ...larkcore.RequestOptionFunc) (*SimplelistGroupMemberResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, g.service.config, http.MethodGet,
		"/open-apis/contact/v3/group/:group_id/member/simplelist", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SimplelistGroupMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *scope) List(ctx context.Context, req *ListScopeReq, options ...larkcore.RequestOptionFunc) (*ListScopeResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/contact/v3/scopes", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListScopeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *unit) BindDepartment(ctx context.Context, req *BindDepartmentUnitReq, options ...larkcore.RequestOptionFunc) (*BindDepartmentUnitResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/contact/v3/unit/bind_department", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BindDepartmentUnitResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *unit) Create(ctx context.Context, req *CreateUnitReq, options ...larkcore.RequestOptionFunc) (*CreateUnitResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/contact/v3/unit", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateUnitResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *unit) Delete(ctx context.Context, req *DeleteUnitReq, options ...larkcore.RequestOptionFunc) (*DeleteUnitResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodDelete,
		"/open-apis/contact/v3/unit/:unit_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteUnitResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *unit) Get(ctx context.Context, req *GetUnitReq, options ...larkcore.RequestOptionFunc) (*GetUnitResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodGet,
		"/open-apis/contact/v3/unit/:unit_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetUnitResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *unit) List(ctx context.Context, req *ListUnitReq, options ...larkcore.RequestOptionFunc) (*ListUnitResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodGet,
		"/open-apis/contact/v3/unit", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListUnitResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *unit) ListDepartment(ctx context.Context, req *ListDepartmentUnitReq, options ...larkcore.RequestOptionFunc) (*ListDepartmentUnitResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodGet,
		"/open-apis/contact/v3/unit/list_department", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListDepartmentUnitResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *unit) Patch(ctx context.Context, req *PatchUnitReq, options ...larkcore.RequestOptionFunc) (*PatchUnitResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodPatch,
		"/open-apis/contact/v3/unit/:unit_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchUnitResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *unit) UnbindDepartment(ctx context.Context, req *UnbindDepartmentUnitReq, options ...larkcore.RequestOptionFunc) (*UnbindDepartmentUnitResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/contact/v3/unit/unbind_department", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UnbindDepartmentUnitResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *user) BatchGetId(ctx context.Context, req *BatchGetIdUserReq, options ...larkcore.RequestOptionFunc) (*BatchGetIdUserResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/contact/v3/users/batch_get_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchGetIdUserResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *user) Create(ctx context.Context, req *CreateUserReq, options ...larkcore.RequestOptionFunc) (*CreateUserResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/contact/v3/users", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateUserResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *user) Delete(ctx context.Context, req *DeleteUserReq, options ...larkcore.RequestOptionFunc) (*DeleteUserResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodDelete,
		"/open-apis/contact/v3/users/:user_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteUserResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *user) FindByDepartment(ctx context.Context, req *FindByDepartmentUserReq, options ...larkcore.RequestOptionFunc) (*FindByDepartmentUserResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodGet,
		"/open-apis/contact/v3/users/find_by_department", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &FindByDepartmentUserResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *user) FindByDepartmentByIterator(ctx context.Context, req *FindByDepartmentUserReq, options ...larkcore.RequestOptionFunc) (*FindByDepartmentUserIterator, error) {
	return &FindByDepartmentUserIterator{
		ctx:      ctx,
		req:      req,
		listFunc: u.FindByDepartment,
		options:  options,
		limit:    req.Limit}, nil
}
func (u *user) Get(ctx context.Context, req *GetUserReq, options ...larkcore.RequestOptionFunc) (*GetUserResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodGet,
		"/open-apis/contact/v3/users/:user_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetUserResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *user) List(ctx context.Context, req *ListUserReq, options ...larkcore.RequestOptionFunc) (*ListUserResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodGet,
		"/open-apis/contact/v3/users", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListUserResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *user) ListByIterator(ctx context.Context, req *ListUserReq, options ...larkcore.RequestOptionFunc) (*ListUserIterator, error) {
	return &ListUserIterator{
		ctx:      ctx,
		req:      req,
		listFunc: u.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (u *user) Patch(ctx context.Context, req *PatchUserReq, options ...larkcore.RequestOptionFunc) (*PatchUserResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodPatch,
		"/open-apis/contact/v3/users/:user_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchUserResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *user) Update(ctx context.Context, req *UpdateUserReq, options ...larkcore.RequestOptionFunc) (*UpdateUserResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, u.service.config, http.MethodPut,
		"/open-apis/contact/v3/users/:user_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateUserResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
